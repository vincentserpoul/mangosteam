package inventory

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestGetUserInventoryURL(t *testing.T) {

	steamID := mangosteam.SteamID(76561198238395094)
	appID := mangosteam.AppID(730)

	inventoryURL := getUserInventoryURL(steamID, appID)
	expectedInventoryURL := mangosteam.BaseSteamWebURL + "profiles/76561198238395094/inventory/json/730/2"

	if inventoryURL != expectedInventoryURL {
		t.Errorf("getUserInventoryURL(%d, %d) expected %s, got %s", steamID, appID,
			expectedInventoryURL, inventoryURL)
	}

}

func getMockJSONInventory(t *testing.T) *Inventory {

	file, err := os.Open("mocks/userInventory.json")
	if err != nil {
		t.Error("getMockJSONInventory() mocks/userInventory.json not found!")
	}

	// dec := gob.NewDecoder(file)
	dec := json.NewDecoder(file)
	var userInventory Inventory

	err = dec.Decode(&userInventory)

	if err != nil {
		t.Error(err)
	}

	return &userInventory

}
