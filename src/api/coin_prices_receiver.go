package main

import (
	"database/sql"
	"log"
)

type CoinPrice struct {
	Symbol string
	Price  float64
}

func getCoinPrices(pgDb *sql.DB) []CoinPrice {
	// Query the DB
	rows, err := pgDb.Query(`SELECT cp.symbol, cp.price FROM coin_prices cp;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var coinPrices []CoinPrice
	var symbol string
	var price float64
	for rows.Next() {

		err := rows.Scan(&symbol, &price)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(symbol, price)
		coinPrices = append(coinPrices, CoinPrice{Symbol: symbol, Price: price})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return coinPrices
}
