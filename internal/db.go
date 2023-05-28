package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Init() {
	var err error

	err = godotenv.Load("../.db.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	// Print out the environment variables for debugging purposes
	// fmt.Println("POSTGRES_USER:", username)
	// fmt.Println("POSTGRES_PASSWORD:", password)
	// fmt.Println("POSTGRES_DB:", dbname)
	// fmt.Println("POSTGRES_HOST:", host)
	// fmt.Println("POSTGRES_PORT:", port)

	status := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err = sql.Open("postgres", status)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected!")
}
