package main

import (
	"fmt"

	"github.com/kasmetski/cmcAPI"
)

func main() {
	getMarket, _ := cmcAPI.GetMarketData()
	fmt.Println(getMarket.ActiveMarkets)
}
