package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"net/http"
	"strconv"
	"top-coins-task/src/producer/top_assets/model"
)

type Assets struct {
	Data []Asset
}

type Asset struct {
	CoinInfo CoinInfo
}
type CoinInfo struct {
	Name     string
	FullName string
}

func SaveTopAssets(db *pg.DB) {
	client := &http.Client{}
	page := 0
	rank := 1
	for page <= 1 {

		topAssetsResult := getTopAssets(client, strconv.Itoa(page))
		var topAssets Assets
		json.Unmarshal(topAssetsResult, &topAssets)
		fmt.Println(topAssets)
		for _, coinPrice := range topAssets.Data {
			newCP := &model.TopAsset{
				Rank:   rank,
				Symbol: coinPrice.CoinInfo.Name,
			}
			newCP.Save(db)
			fmt.Println(coinPrice.CoinInfo.Name, rank)
			rank++
		}
		page++
	}
}
