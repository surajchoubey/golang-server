package database

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	// connect to database
	DATABASE_URL := "postgres://localhost:5432/practice?sslmode=disable"
	db, err := sql.Open("postgres", DATABASE_URL) // os.Getenv("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Database connection failed ❌")
	}
	fmt.Println("Database connection successful ✅")
	return db
}
