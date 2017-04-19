package main

import (
	"fmt"
	"log"

	"github.com/kasmetski/cmcAPI"
)

func main() {
	getMarket, err := cmcAPI.GetMarketData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getMarket)
	getBTC, err := cmcAPI.GetCoinInfo("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getBTC)

	getCoins, err := cmcAPI.GetAllCoinInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getCoins[1])

}
