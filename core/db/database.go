package db

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    connStr := os.Getenv("DB_CONNECTION")
    if connStr == "" {
        connStr = "koch:password@tcp(db:3306)/MoneyWeight"
    }
    DB, err = sql.Open("mysql", connStr)
    if err != nil {
        panic(fmt.Sprintf("Error connecting to the database: %v", err))
    }
    err = DB.Ping()
    if err != nil {
        panic(fmt.Sprintf("Error pinging the database: %v", err))
    }

    fmt.Println("Database connection successful!")
}
