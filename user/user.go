package steamuser

import (
	"fmt"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/auth"
)

// User represents a steam user
type User struct {
	SteamID          mangosteam.SteamID
	SteamMachineAuth string
	SteamLogin       string
	SteamLoginSecure string
	Username         string
	Password         string
	APIKey           string
	Email            string
	LastSessionID    string
}

// Login logs in the bot
func (user *User) Login(baseSteamWebURL string) error {
	isLoggedInclient := user.NewWebSteamClient(baseSteamWebURL)

	isLoggedIn, err := auth.IsLoggedIn(baseSteamWebURL, isLoggedInclient)
	if err != nil {
		return fmt.Errorf("steamuser Login() : %v", err)
	}
	if isLoggedIn {
		return nil
	}
	fmt.Printf("isloggedin: %t\n", isLoggedIn)
	// resetting the login params
	user.SteamLogin = ""
	user.SteamLoginSecure = ""

	client := user.NewWebSteamClient(baseSteamWebURL)

	rsaKey, err := auth.GetRSAKey(baseSteamWebURL, user.Username)
	if err != nil {
		return fmt.Errorf("steamuser Login(): %v", err)
	}

	encryptedPassword, err := auth.EncryptPassword(user.Password, rsaKey)
	if err != nil {
		return fmt.Errorf("steamuser Login(): %v", err)
	}

	sessionID, steamLogin, steamLoginSecure, err := auth.DoLogin(
		baseSteamWebURL,
		client,
		user.Username,
		encryptedPassword,
		rsaKey.Timestamp,
		"", "", "",
	)
	user.LastSessionID = sessionID
	user.SteamLogin = steamLogin
	user.SteamLoginSecure = steamLoginSecure

	if err != nil {
		return fmt.Errorf("steamuser Login(): %v error %v", user.Username, err)
	}

	return nil
}
