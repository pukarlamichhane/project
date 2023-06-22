package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Getconnection() *sql.DB {
	db, err := sql.Open("mysql", "root:123456789@tcp(localhost:3306)/project")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
