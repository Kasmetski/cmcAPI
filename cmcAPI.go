package cmcAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	baseURL = "https://api.coinmarketcap.com/v1/"
	url     string
)

//GetMarketData -
func GetMarketData() (*GlobalMarketData, error) {
	url = fmt.Sprintf(baseURL + "global/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := DoRequest(req)
	if err != nil {
		return nil, err
	}
	var data GlobalMarketData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
