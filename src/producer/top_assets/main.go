package main

import (
	"time"
	"top-coins-task/src/producer/top_assets/db"
)

func main() {
	channel := make(chan int)
	interval := 5

	//setInterval
	go func() {
		for c := time.Tick(time.Duration(interval) * time.Minute); ; <-c {
			pgDb := db.Connect()
			SaveTopAssets(pgDb)
		}
	}()
	<-channel
}
