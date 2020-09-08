package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getTopAssets(client *http.Client, page string) []byte {
	req, err := http.NewRequest("GET", "https://min-api.cryptocompare.com/data/top/totalvolfull", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("limit", "100")
	q.Add("page", page)
	q.Add("tsym", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("api_key", "df9dc44dfecd5029e389f683bcf955815cb8aba053efd13171dc91d1ff8eac76")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to server %s\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}
