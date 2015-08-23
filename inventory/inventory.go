package inventory

import (
	"encoding/json"
	"net/http"

	"github.com/vincentserpoul/mangosteam"
)

// InvChecker is the user interface
type InvChecker interface {
	AreItemsWithinUserInventory(items []*Item) bool
	AreItemsDataSimilarUserInventory(items []*Item) bool
}

// GetUserWebInventory returns the inventory of the user, if available
func GetUserWebInventory(baseSteamWebURL string, appID mangosteam.AppID, steamID mangosteam.SteamID) (*Inventory, error) {
	userInventoryURL := getUserInventoryURL(baseSteamWebURL, steamID, appID)

	resp, err := http.Get(userInventoryURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var inventory Inventory
	err = json.NewDecoder(resp.Body).Decode(&inventory)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func getUserInventoryURL(baseSteamWebURL string, steamID mangosteam.SteamID, appID mangosteam.AppID) string {
	contextID := mangosteam.ContextID(2)
	userInventoryURL := baseSteamWebURL + "/profiles/" +
		steamID.String() + "/inventory/json/" + appID.String() + "/" + contextID.String() +
		"?l=english"

	return userInventoryURL
}

// AreItemsWithinUserInventory returns true or false accordign to the presence of items
func (userInventory *Inventory) AreItemsWithinUserInventory(items []*Item) bool {

	if len(userInventory.Items) == 0 {
		return false
	}

	for _, item := range items {
		_, present := userInventory.Items[item.ID.String()]

		if !present {
			return false
		}
	}

	return true
}

// AreItemsDataSimilarUserInventory returns true or false according to the data of the items
func (userInventory *Inventory) AreItemsDataSimilarUserInventory(items []*Item) bool {

	if len(userInventory.Items) == 0 {
		return false
	}

	for _, item := range items {
		presentItem, present := userInventory.Items[item.ID.String()]
		if present {
			if item.ID != presentItem.ID ||
				item.ClassID != presentItem.ClassID ||
				item.InstanceID != presentItem.InstanceID {

				return false
			}
		}
	}

	return true
}
