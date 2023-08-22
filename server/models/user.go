package models

type UserStatus string

const (
	Inactive   UserStatus = "I"
	Active     UserStatus = "A"
	Terminated UserStatus = "T"
)

// User model
// @Description User account information
type User struct {
	UserID     int        `json:"userID"`
	UserName   string     `json:"userName"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	UserStatus UserStatus `json:"userStatus"`
}

func GetUsers() []User {
	return make([]User, 0)
}

func CreateUser(userName string) {}

func DeleteUser(userId int) {}

func updateUser(user User) {}
