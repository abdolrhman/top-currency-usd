package main

import (
	"github.com/go-pg/pg"
	"log"
	"os"
)

func connect() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "password",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to db. \n")
		os.Exit(100)
	}
	log.Printf("Connection to db successful. \n")
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed succuessfully")
	return
}

func main() {

}
