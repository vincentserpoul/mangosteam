package inventory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestGetUserWebInventoryURL(t *testing.T) {

	steamID := mangosteam.SteamID(1234567890)
	appID := mangosteam.AppID(730)
	baseSteamWebURL := "https://steam"

	inventoryURL := getUserWebInventoryURL(baseSteamWebURL, steamID, appID)
	expectedInventoryURL := "https://steam/profiles/1234567890/inventory/json/730/2?l=english"

	if inventoryURL != expectedInventoryURL {
		t.Errorf("getUserInventoryURL(%d, %d) expected %s, got %s", steamID, appID,
			expectedInventoryURL, inventoryURL)
	}

}

func TestGetUserWebInventoryMock(t *testing.T) {
	steamID := mangosteam.SteamID(1234567890)
	appID := mangosteam.AppID(730)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetMockOKProfilesInventory())
	}))
	defer ts.Close()

	userInventory, err := GetUserWebInventory(ts.URL, appID, steamID)

	if err != nil {
		t.Errorf("GetUserWebInventory(%s, %v, %v) error: %v", ts.URL, appID, steamID, err)
	}
	// expectedUserInventory := Inventory{
	// 	Items: inventory.Items{
	// 		"8742038": &Item{
	// 			ID: 8742038,
	// 			ClassID: 77838,
	// 			InstanceID: 0,
	// 			Amount:    1,
	// 			Pos:  1 ,
	// 		},
	// 		"172795": &Item{
	// 			ID: 8742038,
	// 			ClassID: 77838,
	// 			InstanceID: 0,
	// 			Amount:    1,
	// 			Pos:  1 ,
	// 		}
	// 	},
	// 	Descriptions: inventory.Descriptions{
	// 		"2107773_0": (*inventory.Description)(0xc208032ea0),
	// 		"77838_0": (*inventory.Description)(0xc208032b60)
	// 	},
	// 	AppInfo: (*inventory.AppInfo)(nil)}
	fmt.Printf("%#v", userInventory)
}
