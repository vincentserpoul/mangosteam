package steamuser

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetAPIKey allows to get the API key directly from steam interface
func (user *User) getAPIKey(baseSteamWebURL string) (string, error) {

	client := user.NewWebSteamClient(baseSteamWebURL)

	resp, err := client.Get(baseSteamWebURL + "/dev/apikey")
	if err != nil {
		return "", fmt.Errorf("steamuser GetAPIKey(): %v error %v", user.Username, err)
	}
	defer resp.Body.Close()

	var APIKey string

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("steamuser GetAPIKey(): %v error %v", user.Username, err)
	}

	access := doc.Find("#mainContents h2").Text()

	if access == "Access Denied" {
		return "", fmt.Errorf("steamuser GetAPIKey(): error Access Denied")
	}

	title := doc.Find("div#bodyContents_ex h2").Text()
	if title == "Your Steam Web API Key" {

		s := doc.Find("div#bodyContents_ex p").Eq(0)
		node := strings.Split(s.Text(), ":")
		PotentialAPIKey := strings.TrimSpace(node[1])
		if len(PotentialAPIKey) == 32 {
			APIKey = PotentialAPIKey
		}
	}
	/*
		if APIKey == "" {
			err = user.registerAPIKey(baseSteamWebURL)
			if err != nil {
				return "", fmt.Errorf("steamuser GetAPIKey(): %v error %v", user.Username, err)
			}
			return "", fmt.Errorf("Empty or APIKey  lenght not 32 , for user %v", user.Username) // JJS_TEST
			// return user.getAPIKey(baseSteamWebURL) JJS_TEST Impossible avec la récursion
		}
	*/
	if APIKey == "" {
		return "", fmt.Errorf("GetAPIkey : Empty or APIKey  lenght not 32 , for user %v", user.Username)
	}
	return APIKey, nil
}

// registerAPIKey allows to request for an API key
func (user *User) registerAPIKey(baseSteamWebURL string) error {

	client := user.NewWebSteamClient(baseSteamWebURL)

	baseURL, err := url.Parse(baseSteamWebURL + "/dev/registerkey")
	if err != nil {
		return fmt.Errorf("steamuser registerAPIKey(): %v error %v", user.Username, err)
	}

	form := url.Values{}
	form.Add("domain", "localhost")
	form.Add("agreeToTerms", "agreed")
	form.Add("sessionid", "1")
	form.Add("submit", "Register")

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("steamuser registerAPIKey(): %v error %v", user.Username, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("steamuser registerAPIKey(): %v error %v", user.Username, err)
	}

	defer resp.Body.Close()

	return nil
}

// revokeAPIKey cancel Key
func (user *User) revokeAPIKey(baseSteamWebURL string) error {

	client := user.NewWebSteamClient(baseSteamWebURL)

	baseURL, err := url.Parse(baseSteamWebURL + "/dev/revokekey")
	if err != nil {
		return fmt.Errorf("steamuser revokeAPIKey(): %v error %v", user.Username, err)
	}

	form := url.Values{}
	form.Add("domain", "localhost")
	form.Add("sessionid", "1")
	form.Add("submit", "Revoke My Steam Web API Key")

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("steamuser revokeAPIKey(): %v error %v", user.Username, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("steamuser revokeAPIKey(): %v error %v", user.Username, err)
	}

	defer resp.Body.Close()

	return nil
}
