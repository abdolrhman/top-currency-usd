package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

type TopAsset struct {
	ID     int    `pg:"id,pk"`
	Symbol string `pg:"symbol,unique"`
	Rank   int    `pg:"rank"`
}

func (cp *TopAsset) Save(db *pg.DB) error {
	_, insertErr := db.Model(cp).OnConflict("(symbol) DO UPDATE").
		Set("rank = EXCLUDED.rank").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting a new asset  item %v\n", insertErr)
		return insertErr
	}
	log.Printf("Coin %s inserted asset succeusffully.\n", cp.Symbol)
	return nil
}

func CreateTopAssetsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&TopAsset{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table top_assets %v\n", createErr)
	}
	log.Printf("Table top_assets created Succesfully")
	return nil
}
