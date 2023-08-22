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
				user_id INTEGER PRIMARY KEY AUTOINCREMENT,
				username VARCHAR(50) NOT NULL UNIQUE,
				first_name VARCHAR(255) NOT NULL,
				last_name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				department VARCHAR(255) NULL,
				user_status VARCHAR(1) NOT NULL DEFAULT 'A'
		)
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatalf("Problem creating table: %v", err)
	}

	log.Println("Created user table")

	users := []models.User{
		{
			UserName:  "user1",
			FirstName: "jane",
			LastName:  "doe",
			Email:     "jane@b.com",
		},
		{
			UserName:  "user2",
			FirstName: "john",
			LastName:  "doe",
			Email:     "foo@b.com",
		},
		{
			UserName:  "user3",
			FirstName: "junior",
			LastName:  "doe",
			Email:     "junior@b.com",
		},
	}

	insertBuilder := squirrel.Insert("users")
	for _, user := range users {
		ins := insertBuilder.Columns("username", "first_name", "last_name", "email").
			Values(user.UserName, user.FirstName, user.LastName, user.Email)
		_, err := ins.RunWith(db).Exec()
		if err != nil {
			log.Printf("Could not insert user %s: %v\n", user.UserName, err)
			continue
		}
	}

	log.Println("added default users")

}
