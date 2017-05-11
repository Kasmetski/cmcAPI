package main

import (
	"fmt"
	"log"

	"github.com/kasmetski/cmcAPI"
)

func main() {
	//Get global market data
	getMarket, err := cmcAPI.GetMarketData()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getMarket)
	}
	//Get info about one asset
	getBTC, err := cmcAPI.GetCoinInfo("bitcoin")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getBTC)

	}
	//GetAllCoinInfo(0) for all coins, GetAllCoinInfo(10) for top 10 coins & etc.
	getCoins, err := cmcAPI.GetAllCoinInfo(0)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCoins["ethereum"])
	}

}
