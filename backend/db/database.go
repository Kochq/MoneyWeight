package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "koch:password@tcp(localhost:3306)/MoneyWeight"
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}
	if err = DB.Ping(); err != nil {
		panic(fmt.Sprintf("Error pinging the database: %v", err))
	}
	fmt.Println("Database connection successful!")
}
