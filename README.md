# cmcAPI
Coin Market Cap API package - written in Go

`go get -u -v github.com/kasmetski/cmcAPI`

Get single coin info 	
```
getBTC, err := cmcAPI.GetCoinInfo("bitcoin")
```

Get Allcoins info (all/limit)
```
getCoins, err := cmcAPI.GetAllCoinInfo(0)
```
If you want to limit results to top 10, use `cmcAPI.GetAllCoinInfo(10)`


Get Global Market Data
```
getMarket, err := cmcAPI.GetMarketData()

```

See the `example folder` for more details.

//TODO: fiat currencies convert