package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/curioussavage/integra/models"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MockUserSvc struct {
	GetUsersError error
	Users         []models.User
}

func (m *MockUserSvc) GetUsers(query models.UserFilters) ([]models.User, error) {
	if m.GetUsersError != nil {
		return []models.User{}, m.GetUsersError
	}
	return m.Users, nil
}

func (m *MockUserSvc) CreateUser(user models.UserCreationForm) (models.User, error) {
	return models.User{}, nil
}

func (m *MockUserSvc) DeleteUser(userID int) error {
	return nil
}

func (m *MockUserSvc) UpdateUser(userID int, user models.UserUpdateForm) (models.User, error) {
	return models.User{}, nil
}

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = Describe("GetUsersController", func() {
	var e *echo.Echo
	var controller UserController
	var userSvc *MockUserSvc

	BeforeEach(func() {
		e = echo.New()
		userSvc = &MockUserSvc{}
		controller = NewUserController(userSvc)
	})

	When("Happy path", func() {
		It("returns a list of users", func() {
			testUser := models.User{
				UserID:    1,
				UserName:  "foo",
				FirstName: "peter",
				LastName:  "parker",
			}
			userSvc.Users = []models.User{testUser}

			req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := controller.GetUsers(c)

			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusOK))
			j, _ := json.Marshal(userSvc.Users)
			Expect(rec.Body.String()).To(MatchJSON(j))
		})
	})

	When("Query arg binding fails", func() {
		It("Should return bad request", func() {
			q := make(url.Values)
			q.Set("id", "foobar")
			req := httptest.NewRequest(http.MethodGet, "/api/v1/users?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := controller.GetUsers(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
		})
	})

	When("User service returns an error", func() {
		It("Should return internal server error", func() {
			userSvc.GetUsersError = errors.New("DB error")
			req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := controller.GetUsers(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
		})
	})

})
