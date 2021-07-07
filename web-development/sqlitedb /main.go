package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html
func main() {

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()
	q := `CREATE TABLE IF NOT EXISTS userinfo (
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
    );`
	// Executes the SQL query in our database. Check err to ensure there was no error.
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
	// insert
	{
		stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES (?,?,?)")
		res, err := stmt.Exec("Mohamed", "CS", time.Now())
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
	}

	//update
	{
		stmt, err := db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
		checkErr(err)

		tx, err := db.Begin()
		checkErr(err)
		_, err = tx.Stmt(stmt).Exec("Moha", 2)

		if err != nil {
			fmt.Println("doing rollback")
			tx.Rollback()
		} else {
			tx.Commit()
		}
		// res, err := stmt.Exec("Moha", 1)
		// n, err := res.RowsAffected()
		// checkErr(err)
		// fmt.Println(n)
	}
	{
		type User struct {
			Id        int
			Name      string
			Dprt      string
			CreatedAt time.Time
		}
		var users []User
		// select
		rows, err := db.Query("SELECT * FROM userinfo")
		checkErr(err)
		for rows.Next() {
			var u User
			err := rows.Scan(&u.Id, &u.Name, &u.Dprt, &u.CreatedAt)
			checkErr(err)
			users = append(users, u)
		}
		fmt.Printf("%v\n", users)
	}

	// delete
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
