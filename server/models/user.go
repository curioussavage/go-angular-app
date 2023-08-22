package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/mattn/go-sqlite3"
)

type UsernameTakenError struct {
	name string
}

func (e *UsernameTakenError) Error() string {
	return fmt.Sprintf("username %s already exists", e.name)
}

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
	UserName   string     `json:"userName" db:"username" validate:"min=3,max=50"`
	FirstName  string     `json:"firstName" db:"first_name" validate:"min=1,max=255"`
	LastName   string     `json:"lastName" db:"last_name" validate:"min=1,max=255"`
	Email      string     `json:"email" db:"email" validate:"email,min=3,max=255"`
	Department string     `json:"department" db:"department" validate:"omitempty,min=1,max=255"`
	UserStatus UserStatus `json:"userStatus" db:"user_status" validate:"omitempty,oneof=I A T"`
}

// User create form
// @Description new user form data
type UserCreationForm struct {
	UserName   string `json:"userName" validate:"min=3,max=50"`
	FirstName  string `json:"firstName" validate:"min=1,max=255"`
	LastName   string `json:"lastName" validate:"min=1,max=255"`
	Email      string `json:"email" validate:"email,min=3,max=255"`
	Department string `json:"department" validate:"omitempty,min=1,max=255"`
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

func (u *UserService) CreateUser(user UserCreationForm) (User, error) {
	ins := squirrel.Insert("users").
		Columns("username", "first_name", "last_name", "email", "department").
		Values(user.UserName, user.FirstName, user.LastName, user.Email, user.Department)

	res, err := ins.RunWith(u.DB).Exec()
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return User{}, &UsernameTakenError{name: user.UserName}
			}
		}
		return User{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}

	if newUser, err := u.getUser(int(lastId)); err != nil {
		return User{}, err
	} else {
		return newUser, nil
	}
}

func (u *UserService) DeleteUser(userId int) error {
	del := squirrel.Delete("users").Where(squirrel.Eq{"user_id": userId})
	_, err := del.RunWith(u.DB).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateUser(user User) (User, error) {
	updateQuery := squirrel.Update("users").Where(squirrel.Eq{"user_id": user.UserID})

	if user.UserName != "" {
		updateQuery = updateQuery.Set("username", user.UserName)
	}
	if user.FirstName != "" {
		updateQuery = updateQuery.Set("first_name", user.FirstName)
	}
	if user.LastName != "" {
		updateQuery = updateQuery.Set("last_name", user.LastName)
	}
	if user.Email != "" {
		updateQuery = updateQuery.Set("email", user.Email)
	}
	if user.UserStatus != "" {
		updateQuery = updateQuery.Set("user_status", user.UserStatus)
	}

	if user.Department == "" {
		updateQuery = updateQuery.Set("department", nil)
	} else {
		updateQuery = updateQuery.Set("department", user.Department)
	}

	_, err := updateQuery.RunWith(u.DB).Exec()
	if err != nil {
		return User{}, fmt.Errorf("Could not update user: %v", err)
	}

	newUser, err := u.getUser(user.UserID)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (u *UserService) getUser(userId int) (User, error) {
	query := squirrel.Select("*").
		From("users").
		Where(squirrel.Eq{"user_id": userId})

	var user User
	var dep sql.NullString
	err := query.RunWith(u.DB).QueryRow().
		Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &dep, &user.UserStatus)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
