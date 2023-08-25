package models

import (
	"database/sql"
	"io"
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserService Suite")
}

// CopyDB copies a SQLite database file from srcPath to dstPath.
func CopyDB(srcPath, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

var db *sql.DB

var _ = BeforeSuite(func() {
	Expect(CopyDB("../test.db", "../current_test.db")).To(Succeed())
	newdb, err := sql.Open("sqlite3", "../current_test.db")
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}
	db = newdb
})

var _ = AfterSuite(func() {
	Expect(db.Close()).To(Succeed())
	Expect(os.Remove("../current_test.db")).To(Succeed())
})

var _ = Describe("CreateUser", func() {

	When("Happy path", func() {
		It("returns the new user", func() {
			svc := UserService{DB: db}
			uForm := UserCreationForm{
				UserName:   "user4",
				FirstName:  "test",
				LastName:   "user",
				Email:      "a@b.com",
				Department: "",
			}

			NewUser, err := svc.CreateUser(uForm)
			Expect(err).ToNot(HaveOccurred())
			Expect(NewUser.UserName).To(Equal(uForm.UserName))
		})
	})

	When("Duplicate userName submitted", func() {
		It("returns an error", func() {
			svc := UserService{DB: db}
			uForm := UserCreationForm{
				UserName:   "user1",
				FirstName:  "test",
				LastName:   "user",
				Email:      "a@b.com",
				Department: "",
			}

			_, err := svc.CreateUser(uForm)
			Expect(err.Error()).To(ContainSubstring("user1"))
		})
	})
})

var _ = Describe("GetUsers", func() {
	When("Happy path", func() {
		It("returns a list of active users", func() {
			svc := UserService{DB: db}
			users, err := svc.GetUsers(UserFilters{})
			Expect(err).ToNot(HaveOccurred())
			Expect(len(users) > 0).To(BeTrue())
		})
	})
})

var _ = Describe("DeleteUser", func() {
	When("Happy path", func() {
		It("returns nil", func() {
			svc := UserService{DB: db}
			err := svc.DeleteUser(4)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
