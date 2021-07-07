package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
user@unix(/path/to/socket)/dbname?charset=utf8
  user:password@tcp(localhost:5555)/dbname?charset=utf8
  user:password@/dbname
  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
*/
func main() {
	var dataSourceName string = "admin:Moha#1996@(127.0.0.1:3306)/test?parseTime=true"

	db, err := sql.Open("mysql", dataSourceName)
	checkErr(err)
	defer db.Close()

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
		// stmt, err := db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
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
