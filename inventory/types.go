package inventory

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/vincentserpoul/mangosteam"
)

// ItemID represents the unique identifier for an item
type ItemID uint64

// String will turn a ItemID into a string
func (itemID ItemID) String() string {
	return strconv.FormatUint(uint64(itemID), 10)
}

// ClassID and InstanceID represent the unique identifier for an item type
type ClassID uint64

// String will turn a ClassID into a string
func (classID ClassID) String() string {
	return strconv.FormatUint(uint64(classID), 10)
}

// InstanceID and ClassID represent the unique identifier for an item type
type InstanceID uint64

// String will turn a ClassID into a string
func (instanceID InstanceID) String() string {
	return strconv.FormatUint(uint64(instanceID), 10)
}

// Inventory represents the inventory of a user
type Inventory struct {
	Items        Items        `json:"rgInventory"`
	Descriptions Descriptions `json:"rgDescriptions"`
}

// Item represents an item in the inventory
type Item struct {
	ID         ItemID     `json:",string"`
	ClassID    ClassID    `json:",string"`
	InstanceID InstanceID `json:",string"`
	Amount     uint64     `json:",string"`
	Pos        uint64
}

// Items is a map of items in the inventory
type Items map[string]*Item

// UnmarshalJSON for the inventory items
func (i *Items) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}
	return json.Unmarshal(data, (*map[string]*Item)(i))
}

// Description contains the market hash name
type Description struct {
	AppID      mangosteam.AppID `json:",string"`
	ClassID    ClassID          `json:",string"`
	InstanceID InstanceID       `json:",string"`

	IconURL     string `json:"icon_url"`
	IconDragURL string `json:"icon_drag_url"`

	Name           string
	MarketHashName string `json:"market_hash_name"`
	MarketName     string `json:"market_name"`

	// Colors in hex, for example `B2B2B2`
	NameColor       string `json:"name_color"`
	BackgroundColor string `json:"background_color"`

	Type string

	Tradable   int
	Marketable int
	Commodity  int
}

// Descriptions for the inventory
type Descriptions map[string]*Description

// UnmarshalJSON for the inventory currencies
func (d *Descriptions) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}
	return json.Unmarshal(data, (*map[string]*Description)(d))
}
