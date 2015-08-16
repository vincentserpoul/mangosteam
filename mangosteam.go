package mangosteam

import (
	"fmt"
	"strconv"
)

// SteamID represent the 64bits identifier within steam network, a steamID
type SteamID uint64

// ContextID represents the context of the items, vary according to the app
type ContextID uint64

// AppID represents the AppID according to steam
type AppID uint32

// BaseSteamWebURL is the steam url used to do requests
const BaseSteamWebURL = "https://steamcommunity.com/"

// BaseSteamWebURL is the steam url used to do requests
const BaseSteamAPIURL = "https://api.steampowered.com/"

// String will turn a steamID into a string
func (steamID SteamID) String() string {
	return strconv.FormatUint(uint64(steamID), 10)
}

// String will turn a steamID into a string
func (contextID ContextID) String() string {
	return strconv.FormatUint(uint64(contextID), 10)
}

// String will turn a steamID into a string
func (appID AppID) String() string {
	return strconv.FormatUint(uint64(appID), 10)
}

// GetAccountID will turn a steamID into an accountID string
func (steamID SteamID) GetAccountID() string {
	accountID := (steamID >> 0) & 0xFFFFFFFF
	return fmt.Sprintf("%d", accountID)
}

// GetSteamIDFromString returns the steamID uint64 according to steamID in string
func GetSteamIDFromString(strSteamID string) (SteamID, error) {
	var steamID SteamID
	uint64SteamID, err := strconv.ParseUint(strSteamID, 10, 64)

	steamID = SteamID(uint64SteamID)

	if err != nil {
		return 0, err
	}

	return steamID, nil
}

// GetAppIDFromString returns the AppID uint32 according to steamID in string
func GetAppIDFromString(strAppID string) (AppID, error) {

	uint64AppID, err := strconv.ParseUint(strAppID, 10, 64)

	AppID := AppID(uint64AppID)

	if err != nil {
		return 0, err
	}

	return AppID, nil
}
