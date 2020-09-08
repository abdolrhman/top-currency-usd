package db

import (
	"github.com/go-pg/pg"
	"log"
	"os"
	"top-coins-task/src/producer/top_assets/model"
)

func Connect() *pg.DB{
	opt, err := pg.ParseURL("postgres://postgres:password@localhost:5432/wattax-top-coins?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	if db == nil {
		log.Printf("Failed to connect to db. \n")
		os.Exit(100)
	}
	log.Printf("Connection to db successful. \n")
	model.CreateTopAssetsTable(db)
	return db
}
