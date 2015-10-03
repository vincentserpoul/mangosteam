package tradeoffer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ReceiptItem represents an item in the receipt page
type ReceiptItem struct {
	ItemID         uint64 `json:"id,string"`
	OwnerSteamID   uint64 `json:"owner,string"`
	ClassID        uint64 `json:"classid,string"`
	InstanceID     uint64 `json:"instanceid,string"`
	MarketHashName string `json:"market_hash_name"`
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
		return emptyItems, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return emptyItems, fmt.Errorf("tradeoffer GetItemsFromReceipt(%d): http error %d steam",
			tradeID, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyItems, err
	}

	var items []ReceiptItem
	itemsJSON := extractItemJSONFromBody(string(body))

	for _, itemJSON := range itemsJSON {
		var receiptItem ReceiptItem
		err = json.Unmarshal([]byte(itemJSON), &receiptItem)
		if err != nil {
			return emptyItems, err
		}

		items = append(items, receiptItem)
	}

	return items, nil
}

func extractItemJSONFromBody(bodyS string) []string {

	var itemsJSON []string
	bodyLines := strings.Split(bodyS, "\n")

	for _, line := range bodyLines {
		lineContainingItem := strings.Index(line, "oItem = ")
		if lineContainingItem != -1 {
			itemJSON := line[lineContainingItem+8 : len(line)-1]

			itemsJSON = append(itemsJSON, itemJSON)
		}
	}

	return itemsJSON
}
