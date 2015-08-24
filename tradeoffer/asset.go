package tradeoffer

import (
	"strconv"

	"github.com/vincentserpoul/mangosteam"
)

// AssetID is the identifier for an asset in an inventory
type AssetID uint64

// Asset represents an asset in an inventory
type Asset struct {
	AssetID   AssetID              `json:"assetid,string"`
	AppID     mangosteam.AppID     `json:"appid"`
	ContextID mangosteam.ContextID `json:"contextid,string"`
	Amount    uint64               `json:"amount"`
}

// Defaults update assets with default values
func (asset *Asset) Defaults(appID mangosteam.AppID) {

	asset.AppID = appID
	// arbitrary
	asset.Amount = 1
	asset.ContextID = 2

	return
}

// String will turn a AssetID into a string
func (assetID AssetID) String() string {
	return strconv.FormatUint(uint64(assetID), 10)
}
