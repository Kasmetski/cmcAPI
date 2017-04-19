package cmcAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	baseURL = "https://api.coinmarketcap.com/v1"
	url     string
)

//GetMarketData - Get information about the global market data of the cryptocurrencies
func GetMarketData() (GlobalMarketData, error) {
	url = fmt.Sprintf(baseURL + "/global/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := DoRequest(req)
	if err != nil {
		log.Fatal(err)
	}
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := DoRequest(req)
	if err != nil {
		return nil, err
	}

	var data Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//GetAllCoinInfo - Get iinformation about all coins listed in Coin Market Cap
func GetAllCoinInfo() (Coin, error) {
	url = fmt.Sprintf("%s/ticker/", baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)

	}

	resp, err := DoRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	var data Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)

	}
	return data, err
}
