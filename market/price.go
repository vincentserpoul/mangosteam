package market

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/vincentserpoul/mangosteam"
)

// PriceOverview represent a price for an object
type PriceOverview struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	Volume      int    `json:"volume,string"`
	MedianPrice string `json:"median_price"`
}

const priceURL string = "/market/priceoverview"

// GetPrice returns a float for the price of the item
func GetPrice(appID mangosteam.AppID, marketHashName string) (float64, error) {
	priceOverview, err := getPriceOverview(appID, marketHashName)
	if err != nil {
		return 0, fmt.Errorf("market GetPrice(%d, %s): %v", appID, marketHashName, err)
	}
	if !priceOverview.Success {
		return 0, fmt.Errorf("market GetPrice(%d, %s): steam replied with success:false", appID, marketHashName)
	}

	priceOverview.LowestPrice = strings.TrimPrefix(priceOverview.LowestPrice, "$")

	price, err := strconv.ParseFloat(priceOverview.LowestPrice, 64)
	if err != nil {
		return 0, fmt.Errorf("market GetPrice(%d, %s): %v", appID, marketHashName, err)
	}

	return price, nil
}

// GetPrice will return a price for an item
func getPriceOverview(appID mangosteam.AppID, marketHashName string) (*PriceOverview, error) {
	var price PriceOverview
	res, err := http.Get(getPriceURL(appID, marketHashName))
	if err != nil {
		return nil, fmt.Errorf("market getPriceOverview(%d, %s): %v", appID, marketHashName, err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&price)
	if err != nil {
		return nil, fmt.Errorf("market getPriceOverview(%d, %s): %v", appID, marketHashName, err)
	}

	return &price, nil

}

func getPriceURL(appID mangosteam.AppID, marketHashName string) string {

	v := url.Values{}
	v.Add("currency", "1")
	v.Add("appid", appID.String())
	v.Add("format", "json")
	v.Add("language", "en")
	v.Add("market_hash_name", marketHashName)

	return mangosteam.BaseSteamWebURL + priceURL + `?` + v.Encode()
}
