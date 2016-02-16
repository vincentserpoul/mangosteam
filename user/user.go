package steamuser

import (
	"fmt"
	"net/http"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/auth"
)

// User represents a steam user
type User struct {
	SteamID          mangosteam.SteamID `json:"steam_id"`
	Username         string             `json:"username"`
	Password         string
	APIKey           string
	Email            string
	SteamLogin       string
	SteamLoginSecure string
	auth.SteamGuardAccount
	auth.OAuth
}

// Login logs in the bot
func (user *User) Login(baseSteamWebURL string, baseSteamAPIURL string) error {
	isLoggedInclient := user.NewWebSteamClient(baseSteamWebURL)

	isLoggedIn, err := auth.IsLoggedIn(baseSteamWebURL, isLoggedInclient)
	if err != nil {
		return fmt.Errorf("steamuser Login() : %v", err)
	}
	if isLoggedIn {
		fmt.Println("we re in!")
		return nil
	}

	// resetting the login params
	user.SteamLogin = ""
	user.SteamLoginSecure = ""

	client := &http.Client{}
	rsaKey, err := auth.GetRSAKey(baseSteamWebURL, user.Username)
	if err != nil {
		return fmt.Errorf("steamuser Login(): %v", err)
	}

	encryptedPassword, err := auth.EncryptPassword(user.Password, rsaKey)
	if err != nil {
		return fmt.Errorf("steamuser Login(): %v", err)
	}

	steamGuardCode, err := user.GenerateSteamGuardCode(baseSteamAPIURL)
	if err != nil {
		return fmt.Errorf("steamuser Login(): %v", err)
	}

	user.OAuth, err = auth.DoLogin(
		baseSteamWebURL,
		client,
		user.Username,
		encryptedPassword,
		rsaKey.Timestamp,
		"", "", "", steamGuardCode,
	)
	if err != nil {
		return fmt.Errorf("steamuser Login(): %v error %v", user.Username, err)
	}

	user.SteamID = user.OAuth.SteamID
	user.SteamLogin = user.SteamID.String() + "||" + user.OAuth.Token
	user.SteamLoginSecure = user.SteamID.String() + "||" + user.OAuth.TokenSecure

	return nil
}
