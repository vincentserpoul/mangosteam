package steamuser

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/vincentserpoul/mangosteam"
)

// GetAPIKey allows to get the API key directly from steam interface
func (user *User) getAPIKey() (string, error) {

	client := user.NewWebSteamClient()
	resp, err := client.Get(mangosteam.BaseSteamWebURL + "dev/apikey")
	defer resp.Body.Close()

	var APIKey string

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("steam user GetAPIKey(): %v error %v", user.Username, err)
	}

	access := doc.Find("#mainContents h2").Text()

	if access == "Access Denied" {
		return "", fmt.Errorf("steam user GetAPIKey(): error Access Denied")
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

	if APIKey == "" {
		err = user.registerAPIKey()
		if err != nil {
			return "", fmt.Errorf("steam user GetAPIKey(): %v error %v", user.Username, err)
		}

		return user.getAPIKey()
	}

	return APIKey, nil
}

// registerAPIKey allows to request for an API key
func (user *User) registerAPIKey() error {

	client := user.NewWebSteamClient()
	baseURL, _ := url.Parse(mangosteam.BaseSteamWebURL + "dev/registerkey")

	form := url.Values{}
	form.Add("domain", "localhost")
	form.Add("agreeToTerms", "agreed")
	form.Add("sessionid", "1")
	form.Add("submit", "Register")

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("steam user registerAPIKey(): %v error %v", user.Username, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("steam user registerAPIKey(): %v error %v", user.Username, err)
	}

	defer resp.Body.Close()

	return nil
}