package models

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
)

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

type UserService struct {
	DB *sql.DB
}

func (u *UserService) GetUsers() ([]User, error) {
	rows, err := squirrel.Select("*").From("users").RunWith(u.DB).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		var department sql.NullString
		err := rows.Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &department, &user.UserStatus)
		if err != nil {
			log.Printf("error while scanning rows: %v", err)
			continue
		}
		if department.Valid {
			user.Department = department.String
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) CreateUser(user User) (User, error) {
	return User{}, nil
}

func (u *UserService) DeleteUser(userId int) error {
	return nil
}

func (u *UserService) UpdateUser(user User) (User, error) {
	return User{}, nil
}
