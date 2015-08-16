package inventory

import (
	"encoding/json"
	"net/http"

	"github.com/vincentserpoul/mangosteam"
)

// GetUserInventory retrieve the inventory from steam
func GetUserInventory(appID mangosteam.AppID, steamID mangosteam.SteamID) (*Inventory, error) {

	userInventoryURL := getUserInventoryURL(steamID, appID)

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

func getUserInventoryURL(steamID mangosteam.SteamID, appID mangosteam.AppID) string {
	contextID := mangosteam.ContextID(2)
	userInventoryURL := mangosteam.BaseSteamWebURL + "profiles/" +
		steamID.String() + "/inventory/json/" + appID.String() + "/" + contextID.String()

	return userInventoryURL
}
