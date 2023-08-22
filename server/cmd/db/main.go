package main

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/curioussavage/integra/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}
	defer db.Close()

	sql := `
	  CREATE TABLE users (
				user_id SERIAL PRIMARY KEY,
				username VARCHAR(50) NOT NULL,
				first_name VARCHAR(255) NOT NULL,
				last_name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				user_status VARCHAR(1) NOT NULL,
				department VARCHAR(255) NULL
		)
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatalf("Problem creating table: %v", err)
	}

	log.Println("Created user table")

	users := []models.User{
		{
			UserName:   "user1",
			FirstName:  "jane",
			LastName:   "doe",
			Email:      "jane@b.com",
			UserStatus: models.Active,
		},
		{
			UserName:   "user1",
			FirstName:  "john",
			LastName:   "doe",
			Email:      "foo@b.com",
			UserStatus: models.Active,
		},
		{
			UserName:   "user1",
			FirstName:  "junior",
			LastName:   "doe",
			Email:      "junior@b.com",
			UserStatus: models.Active,
		},
	}

	insertBuilder := squirrel.Insert("users")
	for _, user := range users {
		ins := insertBuilder.Columns("username", "first_name", "last_name", "email", "user_status").
			Values(user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus)
		_, err := ins.RunWith(db).Exec()
		if err != nil {
			log.Printf("Could not insert user %s: %v\n", user.UserName, err)
			continue
		}
	}

	log.Println("added default users")

}
