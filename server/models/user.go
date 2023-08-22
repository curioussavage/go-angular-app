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
	UserID     int        `json:"userID" db:"user_id"`
	UserName   string     `json:"userName" db:"username"`
	FirstName  string     `json:"firstName" db:"first_name"`
	LastName   string     `json:"lastName" db:"last_name"`
	Email      string     `json:"email" db:"email"`
	Department string     `json:"department" db:"department"`
	UserStatus UserStatus `json:"userStatus" db:"user_status"`
}

func GetUsers() []User {
	return make([]User, 0)
}

func CreateUser(userName string) {}

func DeleteUser(userId int) {}

func updateUser(user User) {}
