package main

import (
	"time"
	"top-coins-task/src/producer/coins_prices/db"
)

func main() {
	channel := make(chan int)
	interval := 5

	// setInterval
	go func() {
		for c := time.Tick(time.Duration(interval) * time.Minute); ; <-c {
			// Establish Connection
			pgDb := db.Connect()

			// Consume The CoinsPrices API
			// And saves Data
			SaveCoinPrice(pgDb)
		}
	}()
	<-channel
}
