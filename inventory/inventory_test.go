package inventory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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
		t.Errorf("getUserWebInventoryURL(%d, %d) expected %s, got %s", steamID, appID,
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

	expectedUserInventory := &Inventory{
		Items: Items{
			"172795": &Item{
				ID:         ItemID(172795),
				ClassID:    ClassID(2107773),
				InstanceID: InstanceID(0),
				Amount:     uint64(1),
				Pos:        uint64(2),
			},
			"8742038": &Item{
				ID:         ItemID(8742038),
				ClassID:    ClassID(77838),
				InstanceID: InstanceID(0),
				Amount:     uint64(1),
				Pos:        uint64(1),
			},
		},
		Descriptions: Descriptions{
			"77838_0": &Description{
				AppID:           mangosteam.AppID(730),
				ClassID:         ClassID(77838),
				InstanceID:      InstanceID(0),
				IconURL:         "-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9Q1LO5kNoBhSQl-fEv2o1t3QXFR6a1wE4uOkKlFm0qvJd2gSvYS3x9nbwfXyZrqBxDkCvZYmjurEpomlilL6ux07YtuiRwA",
				IconDragURL:     "",
				Name:            "5 Year Veteran Coin",
				MarketHashName:  "5 Year Veteran Coin",
				MarketName:      "5 Year Veteran Coin",
				NameColor:       "D2D2D2",
				BackgroundColor: "",
				Type:            "Extraordinary Collectible",
				Tradable:        0,
				Marketable:      0,
				Commodity:       0,
			},
			"2107773_0": &Description{
				AppID:           mangosteam.AppID(730),
				ClassID:         ClassID(2107773),
				InstanceID:      InstanceID(0),
				IconURL:         "-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgporrf0e1Y07ODHTjBN_8-JmYWPnuL5feuJwjlVscQhj7rH9tzw2wXmqkNlYG-hJNWSegc9Zl-E_QK9xbjr08Si_MOejgzGL-s",
				IconDragURL:     "",
				Name:            "XM1014 | Blue Spruce",
				MarketHashName:  "XM1014 | Blue Spruce (Field-Tested)",
				MarketName:      "XM1014 | Blue Spruce (Field-Tested)",
				NameColor:       "D2D2D2",
				BackgroundColor: "",
				Type:            "Consumer Grade Shotgun",
				Tradable:        1,
				Marketable:      1,
				Commodity:       0,
			},
		},
	}

	if !reflect.DeepEqual(expectedUserInventory, userInventory) {
		t.Errorf("GetUserWebInventory(%s, %v, %v): got %#v, expected %#v",
			ts.URL, appID, steamID, userInventory, expectedUserInventory)
	}

}
