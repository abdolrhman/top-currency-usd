package db

import (
	"database/sql"
	_ "github.com/lib/pq" // here
	"log"
	"os"
)

func Connect() *sql.DB {
	connStr := "user=postgres dbname=wattax-top-coins password=password host=localhost sslmode=disable"
	//driver name part of "github.com/lib/pq"
	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to connect to db. \n")
		os.Exit(100)
	}
	log.Printf("Connection to db successful. \n")
	return db
}
