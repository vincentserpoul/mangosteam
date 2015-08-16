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
func (user *User) Login() error {
	client := user.NewWebSteamClient()

	rsaKey, err := auth.GetRSAKey(user.Username)
	if err != nil {
		return fmt.Errorf("user Login(): %v", err)
	}
	encryptedPassword, err := auth.EncryptPassword(user.Password, rsaKey)
	if err != nil {
		return fmt.Errorf("user Login(): %v", err)
	}

	err = auth.DoLogin(
		client,
		user.Username,
		encryptedPassword,
		rsaKey.Timestamp,
		"", "", "",
	)

	if err != nil {
		return fmt.Errorf("user Login(): %v error %v", user.Username, err)
	}

	return nil
}
