package cmcAPI

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	baseURL = "https://api.coinmarketcap.com/v1"
	url     string
	l       string
)

//GetMarketData - Get information about the global market data of the cryptocurrencies
func GetMarketData() (GlobalMarketData, error) {
	url = fmt.Sprintf(baseURL + "/global/")

	resp, err := MakeReq(url)

	var data GlobalMarketData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

//GetCoinInfo - Get information about single crypto currency
func GetCoinInfo(coin string) (Coin, error) {
	url = fmt.Sprintf("%s/ticker/%s", baseURL, coin)
	resp, err := MakeReq(url)

	var data Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

//GetAllCoinInfo - Get information about all coins listed in Coin Market Cap. If you want to limit the search to top 10 coins pass 10 as int, if you want all - pass 0 == No Limit
func GetAllCoinInfo(limit int) (Coin, error) {
	if limit > 0 {
		l = fmt.Sprintf("?limit=%v", limit)
	}
	url = fmt.Sprintf("%s/ticker/%s", baseURL, l)

	resp, err := MakeReq(url)

	var data Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
