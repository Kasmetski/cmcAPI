package cmcAPI

//Tickers -
type Tickers struct {
	Name map[string]Ticker
}

//Ticker -
type Ticker struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int     `json:"rank"`
	PriceUsd         float64 `json:"price_usd"`
	PriceBtc         float64 `json:"price_btc"`
	Two4HVolumeUsd   float64 `json:"24h_volume_usd"`
	MarketCapUsd     float64 `json:"market_cap_usd"`
	AvailableSupply  float64 `json:"available_supply"`
	TotalSupply      float64 `json:"total_supply"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	LastUpdated      string  `json:"last_updated"`
}

//GlobalMarketData -
type GlobalMarketData struct {
	TotalMarketCapUsd float64 `json:"total_market_cap_usd"`
	Total24HVolumeUsd float64 `json:"total_24h_volume_usd"`
	BitcoinDominance  float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies  int     `json:"active_currencies"`
	ActiveAssets      int     `json:"active_assets"`
	ActiveMarkets     int     `json:"active_markets"`
}
