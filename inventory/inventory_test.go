package inventory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/vincentserpoul/mangosteam"
)

func TestAreItemsDataSimilarUserInventoryNotPresent(t *testing.T) {

	var d Inventory
	var inv, inv2 Item
	var listInv []*Item
	inv = Item{1234567890, 730, 3, 100, 1}
	inv2 = Item{12567890, 73, 30, 10, 10}
	listInv = append(listInv, &inv)
	d.Items = make(map[string]*Item)
	d.Items["1234567890"] = &inv2
	r := d.AreItemsDataSimilarUserInventory(listInv)
	if r != false {
		t.Errorf("AreItemsDataSimilarUserInventory should return false")
	}
}

func TestAreItemsDataSimilarUserInventoryOK(t *testing.T) {

	var d Inventory
	var inv Item
	var listInv []*Item
	inv = Item{1234567890, 730, 3, 100, 1}
	listInv = append(listInv, &inv)
	d.Items = make(map[string]*Item)
	d.Items["1234567890"] = &inv
	r := d.AreItemsDataSimilarUserInventory(listInv)
	if r != true {
		t.Errorf("AreItemsDataSimilarUserInventory should return true")
	}
}

func TestAreItemsWithinUserInventoryNotPresent(t *testing.T) {

	var d Inventory
	var inv Item
	var listInv []*Item
	inv = Item{1234567890, 730, 3, 100, 1}
	listInv = append(listInv, &inv)
	d.Items = make(map[string]*Item)
	d.Items["1"] = &inv
	r := d.AreItemsWithinUserInventory(listInv)
	if r != false {
		t.Errorf("AreItemsWithinUserInventory should return false")
	}
}

func TestAreItemsWithinUserInventoryOK(t *testing.T) {

	var d Inventory
	var inv Item
	var listInv []*Item
	inv.ID = 1234567890
	inv.ClassID = 730
	inv.InstanceID = 1
	inv.Amount = 100
	inv.Pos = 1
	listInv = append(listInv, &inv)
	d.Items = make(map[string]*Item)
	d.Items["1234567890"] = &inv
	r := d.AreItemsWithinUserInventory(listInv)
	if r != true {
		t.Errorf("AreItemsWithinUserInventory should return true")
	}
}

func TestAreItemsWithinUserInventoryEmpty(t *testing.T) {
	var d Inventory
	r := d.AreItemsWithinUserInventory(nil)
	if r != false {
		t.Errorf("AreItemsWithinUserInventory should return false")
	}
}

func TestAreItemsDataSimilarUserInventoryEmpty(t *testing.T) {
	var d Inventory
	r := d.AreItemsDataSimilarUserInventory(nil)
	if r != false {
		t.Errorf("AreItemsWithinUserInventory should return false")
	}
}
func TestGetUserWebInventoryEmpty(t *testing.T) {
	steamID := mangosteam.SteamID(1234567890)
	appID := mangosteam.AppID(730)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetMockGetUserWebInventoryEmpty())
	}))
	defer ts.Close()
	_, err := GetUserWebInventory(ts.URL, appID, steamID)
	if err == nil {
		t.Errorf("GetUserWebInventory should return error with no inventory")
	}
}

func TestGetUserWebInventoryTimeout(t *testing.T) {
	steamID := mangosteam.SteamID(1234567890)
	appID := mangosteam.AppID(730)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()
	ts.Config.WriteTimeout = 20 * time.Millisecond
	_, err := GetUserWebInventory(ts.URL, appID, steamID)
	if err == nil {
		t.Errorf("GetUserWebInventory should return error with timeout")
	}

}

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

func TestGetUserWebInventory(t *testing.T) {
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
				Tags: Tags{
					{
						InternalName: "CSGO_Type_Collectible",
						Name:         "Collectible",
						Category:     "Type",
						CategoryName: "Type",
					},
					{
						InternalName: "Rarity_Ancient",
						Name:         "Extraordinary",
						Category:     "Rarity",
						Color:        "eb4b4b",
						CategoryName: "Quality",
					},
				},
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
				Tags: Tags{
					{
						InternalName: "CSGO_Type_Shotgun",
						Name:         "Shotgun",
						Category:     "Type",
						CategoryName: "Type",
					},
					{
						InternalName: "Rarity_Common_Weapon",
						Name:         "Consumer Grade",
						Category:     "Rarity",
						Color:        "b0c3d9",
						CategoryName: "Quality",
					},
					{
						InternalName: "WearCategory2",
						Name:         "Field-Tested",
						Category:     "Exterior",
						CategoryName: "Exterior",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedUserInventory, userInventory) {
		t.Errorf("GetUserWebInventory(%s, %v, %v): got %#v, expected %#v",
			ts.URL, appID, steamID, userInventory, expectedUserInventory)
	}

}
