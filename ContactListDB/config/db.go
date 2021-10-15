package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Connect() *sqlx.DB {
	// Loading environment variables
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Open connection to database
	db, err := sqlx.Connect(dialect, dbURI)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connected to Database !!!")
	}
	return db
}
