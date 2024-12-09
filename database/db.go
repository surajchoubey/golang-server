package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// connect to database
	// DATABASE_URL := "postgres://localhost:5432/practice?sslmode=disable"
	DATABASE_URL := "host=localhost dbname=practice port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{}) // os.Getenv("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Database connection failed ❌")
	}
	fmt.Println("Database connection successful ✅")
	return db
}
