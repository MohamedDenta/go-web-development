package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dataSourceName string = "admin:Moha#1996@(127.0.0.1:3306)/test?parseTime=true"

func main() {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	{ // Create a new table
		query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
		// Executes the SQL query in our database. Check err to ensure there was no error.
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
	{ // Insert a new user
		userName := "Mohamed"
		password := "secret"
		createdAt := time.Now()
		result, err := db.Exec(`INSERT INTO users(username,password,created_at) VALUES (?,?,?)`, userName, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		fmt.Println(id)

	}

	{ // Query a single user
		var (
			id        int
			userName  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id,username,password,created_at FROM users WHERE id=?"
		if err := db.QueryRow(query, 1).Scan(&id, &userName, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, userName, password, createdAt)
	}
	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		q := "SELECT id,username,password,created_at FROM users"
		rows, err := db.Query(q)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, u.username, u.password, u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", users)
	}
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
