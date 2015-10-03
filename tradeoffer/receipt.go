package tradeoffer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/inventory"
)

// ReceiptItem represents an item in the receipt page
type ReceiptItem struct {
	ItemID         inventory.ItemID     `json:"id,string"`
	OwnerSteamID   mangosteam.SteamID   `json:"owner,string"`
	ClassID        inventory.ClassID    `json:"classid,string"`
	InstanceID     inventory.InstanceID `json:"instanceid,string"`
	MarketHashName string               `json:"market_hash_name"`
}

const (
	tradeOfferReceiptURL string = "/trade/%d/receipt/"
)

// GetItemsFromReceipt allows the retrieval of new itemids for the items
func GetItemsFromReceipt(
	client *http.Client,
	baseSteamWebURL string,
	tradeID uint64,
) ([]ReceiptItem, error) {

	var emptyItems []ReceiptItem

	resp, err := client.Get(baseSteamWebURL + fmt.Sprintf(tradeOfferReceiptURL, tradeID))

	if err != nil {
		return emptyItems, fmt.Errorf("tradeoffer GetItemsFromReceipt(%d): %v",
			tradeID, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return emptyItems, fmt.Errorf("tradeoffer GetItemsFromReceipt(%d): http error %d steam",
			tradeID, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyItems, fmt.Errorf("tradeoffer GetItemsFromReceipt(%d): %v",
			tradeID, err)
	}

	var items []ReceiptItem
	itemsJSON := extractItemJSONFromBody(string(body))

	for _, itemJSON := range itemsJSON {
		var receiptItem ReceiptItem
		err = json.Unmarshal([]byte(itemJSON), &receiptItem)
		if err != nil {
			return emptyItems, fmt.Errorf("tradeoffer GetItemsFromReceipt(%d): %v \n%s",
				tradeID, itemJSON, err)
		}

		items = append(items, receiptItem)
	}

	return items, nil
}

func extractItemJSONFromBody(bodyS string) []string {

	var itemsJSON []string
	bodyLines := strings.Split(bodyS, "\n")

	for _, line := range bodyLines {
		lineContainingItem := strings.Index(line, `oItem = {`)
		endJSON := strings.LastIndex(line, `};`)
		if lineContainingItem != -1 && endJSON != -1 {
			itemJSON := line[lineContainingItem+8 : endJSON+1]
			itemsJSON = append(itemsJSON, itemJSON)
		}
	}

	return itemsJSON
}
