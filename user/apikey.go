package steamuser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	getAPIKeyURL      string = "/dev/apikey"
	registerAPIKeyURL string = "/dev/registerkey"
	revokeAPIKeyURL   string = "/dev/revokekey"
)

// GetAPIKey allows to get the API key directly from steam interface
func (user *User) GetAPIKey(baseSteamWebURL string) (string, error) {

	client := user.NewWebSteamClient(baseSteamWebURL)

	resp, err := client.Get(baseSteamWebURL + getAPIKeyURL)
	if err != nil {
		return "", fmt.Errorf("steamuser GetAPIKey(): %v error %v", user.Username, err)
	}
	defer resp.Body.Close()

	APIKey, err := extractAPIKey(resp.Body)
	if err != errKeyNotRegistered && err != nil {
		return "", fmt.Errorf("steamuser GetAPIKey(): %v error %v", user.Username, err)
	}

	if err == errKeyNotRegistered {
		APIKey, err = user.registerAPIKey(baseSteamWebURL)
		if err != nil {
			return "", fmt.Errorf("steamuser GetAPIKey(): %v error %v", user.Username, err)
		}
	}

	if APIKey == "" {
		return "", fmt.Errorf("GetAPIkey : Empty or APIKey length not 32 , for user %v", user.Username)
	}

	user.APIKey = APIKey

	return APIKey, nil
}

// RegisterAPIKey allows to request for an API key
func (user *User) registerAPIKey(baseSteamWebURL string) (string, error) {

	client := user.NewWebSteamClient(baseSteamWebURL)

	baseURL, err := url.Parse(baseSteamWebURL + registerAPIKeyURL)
	if err != nil {
		return "", fmt.Errorf("steamuser RegisterAPIKey(): %v error %v", user.Username, err)
	}

	form := url.Values{}
	form.Add("domain", "localhost")
	form.Add("agreeToTerms", "agreed")
	form.Add("sessionid", user.LastSessionID)
	form.Add("submit", "Register")

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("steamuser RegisterAPIKey(): %v error %v", user.Username, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("steamuser RegisterAPIKey(): %v error %v", user.Username, err)
	}

	defer resp.Body.Close()

	APIKey, err := extractAPIKey(resp.Body)
	if err != nil {
		return "", fmt.Errorf("steamuser RegisterAPIKey(): %v error %v", user.Username, err)
	}
	return APIKey, nil
}

var errKeyNotRegistered = errors.New("key not registered yet")

func extractAPIKey(htmlContent io.Reader) (string, error) {
	var APIKey string

	doc, err := goquery.NewDocumentFromReader(htmlContent)
	if err != nil {
		return "", fmt.Errorf("steamuser extractAPIKey(): %v", err)
	}

	access := doc.Find("#mainContents h2").Text()
	if access == "Access Denied" {
		return "", fmt.Errorf("steamuser GetAPIKey(): error Access Denied")
	}

	title := doc.Find("div#bodyContents_ex h2").Text()

	if title == "Register for a new Steam Web API Key" {
		return "", errKeyNotRegistered
	}

	if title == "Your Steam Web API Key" {
		s := doc.Find("div#bodyContents_ex p").Eq(0)
		node := strings.Split(s.Text(), ":")
		PotentialAPIKey := strings.TrimSpace(node[1])
		if len(PotentialAPIKey) == 32 {
			APIKey = PotentialAPIKey
		}
	}

	return APIKey, nil
}

// RevokeAPIKey cancel Key
func (user *User) RevokeAPIKey(baseSteamWebURL string) error {

	client := user.NewWebSteamClient(baseSteamWebURL)

	baseURL, err := url.Parse(baseSteamWebURL + revokeAPIKeyURL)
	if err != nil {
		return fmt.Errorf("steamuser RevokeAPIKey(): %v error %v", user.Username, err)
	}

	form := url.Values{}
	form.Add("domain", "localhost")
	form.Add("sessionid", "1")
	form.Add("submit", "Revoke My Steam Web API Key")

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("steamuser RevokeAPIKey(): %v error %v", user.Username, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("steamuser RevokeAPIKey(): %v error %v", user.Username, err)
	}

	defer resp.Body.Close()

	return nil
}
