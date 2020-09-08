package main

import "database/sql"

type TopAssetsPrices struct {
	Symbol string
	Price  float64
	Rank   int
}

func handler(pgDb *sql.DB, limit int) []TopAssetsPrices {
	var topAssetsPrices []TopAssetsPrices

	var rows, err = pgDb.Query("SELECT * FROM top_assets ORDER BY rank limit $1", limit)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var uid int
		var symbol string
		var rank int
		err = rows.Scan(&uid, &symbol, &rank)
		if err != nil {
			panic(err)
		}
		coinsPrices := getCoinPrices(pgDb)
		for _, coinPrice := range coinsPrices {
			if coinPrice.Symbol == symbol {
				topAssetsPrices = append(topAssetsPrices, TopAssetsPrices{
					Symbol: symbol,
					Rank:   rank,
					Price:  coinPrice.Price,
				})
				break
			}

		}
		//fmt.Println("uid | symbol | rank | created ")
		//fmt.Printf("%3v | %8v | %6v ", uid, symbol, rank)
	}
	return topAssetsPrices
}
