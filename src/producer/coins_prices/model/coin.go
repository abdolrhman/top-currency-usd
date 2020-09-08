package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

type CoinPrice struct {
	ID     int     `pg:"id,pk"`
	Symbol string  `pg:"symbol,unique"`
	Price  float64 `pg:"price,type:real"`
}

func (cp *CoinPrice) Save(db *pg.DB) error {
	_, insertErr := db.Model(cp).OnConflict("(symbol) DO UPDATE").
		Set("price = EXCLUDED.price").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting a new coin price item %v\n", insertErr)
		return insertErr
	}
	log.Printf("Coin %s inserted price succeusffully.\n", cp.Symbol)
	return nil
}

func CreateCoinsPricesTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&CoinPrice{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table coins prices %v\n", createErr)
	}
	log.Printf("Table coins prices created Succesfully")
	return nil
}
