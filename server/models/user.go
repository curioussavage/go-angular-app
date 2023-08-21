package models

type UserStatus string

const (
	Inactive   UserStatus = "I"
	Active     UserStatus = "A"
	Terminated UserStatus = "T"
)

type User struct {
	UserName   string
	FirstName  string
	LastName   string
	Email      string
	UserStatus UserStatus
}

func GetUsers() []User {
	return make([]User, 0)
}

func CreateUser(userName string) {}

func DeleteUser(userId int) {}

func updateUser(user User) {}
