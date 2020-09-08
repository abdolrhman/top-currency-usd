package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"net/http"
	"top-coins-task/src/producer/coins_prices/model"
)

type Assets struct {
	Data []Asset
}
// API Response Model
type Asset struct {
	Symbol string
	Quote  struct {
		USD struct {
			Price float64
		}
	}
}

func SaveCoinPrice(db *pg.DB) {
	client := &http.Client{}

	// Gets API Data
	userPricesResult := getUSDPrices(client)
	var usdPrices Assets
	json.Unmarshal(userPricesResult, &usdPrices)
	for _, coinPrice := range usdPrices.Data {
		newCP := &model.CoinPrice{
			Price:  coinPrice.Quote.USD.Price,
			Symbol: coinPrice.Symbol,
		}
		newCP.Save(db)
		fmt.Println(coinPrice.Symbol, coinPrice.Quote.USD.Price)
	}
}

