package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type User struct {
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Role      string
	Pic       string
}

type Users struct {
	Username string
	Pic      string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("MySQLUserName")
	dbPass := os.Getenv("MySQLPassword")
	dbName := os.Getenv("DBName")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateUser() {
}

func ReadUser(userName string) User {
	db := dbConn()

	user := User{}

	selectUser := `SELECT username, email, first_name, last_name, role, pic FROM users WHERE username=?`
	row := db.QueryRow(selectUser, userName)
	err := row.Scan(&user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Role, &user.Pic)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	return user
}

func UpdateUser() {

}

func DeleteUser() {

}

func ShowUsers() []Users {
	db := dbConn()

	var users []Users

	selectUsers := `SELECT username, pic FROM users`
	rows, err := db.Query(selectUsers)
	defer rows.Close()

	for rows.Next() {
		user := Users{}
		err = rows.Scan(&user.Username, &user.Pic)
		if err != nil {
			panic(err)
		}
		users = append(users, Users{Username: user.Username, Pic: user.Pic})
	}

	return users
}
