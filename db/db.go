package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connectionString := "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}
