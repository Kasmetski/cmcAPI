package cmcAPI

//Coin struct
type Coin []struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

//GlobalMarketData struct
type GlobalMarketData struct {
	TotalMarketCapUsd float64 `json:"total_market_cap_usd"`
	Total24HVolumeUsd float64 `json:"total_24h_volume_usd"`
	BitcoinDominance  float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies  int     `json:"active_currencies"`
	ActiveAssets      int     `json:"active_assets"`
	ActiveMarkets     int     `json:"active_markets"`
}
