package inventory

import (
	"bytes"
	"encoding/json"

	"github.com/vincentserpoul/mangosteam"
)

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

	Tags Tags
}

// Tags is a list of tags in the description
type Tags []*Tag

// Descriptions for the inventory
type Descriptions map[string]*Description

// UnmarshalJSON for the inventory currencies
func (d *Descriptions) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}
	return json.Unmarshal(data, (*map[string]*Description)(d))
}

// Tag is an unstructured data
type Tag struct {
	InternalName string `json:"internal_name"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	Color        string `json:"color"`
}

// GetTagNameFromCategory returns the tag name for a specific category
func (t *Tags) GetTagNameFromCategory(category string) string {

	for _, tag := range *t {
		if tag.Category == category {
			return tag.Name
		}
	}

	return ""
}
