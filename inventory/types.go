package inventory

import (
	"bytes"
	"encoding/json"
	"strconv"
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
