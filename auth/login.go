package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// doLogInResponse is used to store the body of the steam doLogin response
type doLogInResponse struct {
	Success         bool `json:"success"`
	EmailauthNeeded bool `json:"emailauth_needed"`
}

const (
	// DoLoginURL URL used for login
	DoLoginURL string = "/login/dologin"
	// IsLoggedInURL URL used to check if user is logged in
	IsLoggedInURL string = "/actions/GetNotificationCounts"
)

// DoLogin is used to log in the steam account after we got the encrypted password
func DoLogin(
	baseSteamWebURL string,
	client *http.Client,
	username string,
	encryptedPassword string,
	rsatimestamp string,
	emailauthKeyedIn string,
	captchaGID string,
	captchaKeyedIn string,
) (string, string, string, error) {

	var sessionid, steamLogin, steamLoginSecure string

	sessionid, err := getNewSessionID(baseSteamWebURL)
	if err != nil {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): %v", err)
	}

	baseURL, _ := url.Parse(baseSteamWebURL + DoLoginURL)

	// default value set to -1
	if captchaGID == "" {
		captchaGID = "-1"
	}

	// adding query params
	params := url.Values{}
	params.Add("password", encryptedPassword)
	params.Add("username", username)
	params.Add("emailauth", emailauthKeyedIn)
	params.Add("twofactorcode", "")
	params.Add("loginfriendlyname", "")
	params.Add("captchagid", captchaGID)
	params.Add("captcha_text", captchaKeyedIn)
	params.Add("emailsteamid", "")
	params.Add("rsatimestamp", rsatimestamp)
	params.Add("remember_login", "true")

	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", baseURL.String(), nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): bad request %v for %s, ",
			resp.Status, username)
	}

	loginBody := new(doLogInResponse)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(loginBody)

	if err != nil {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): %v", err)
	}

	if loginBody.EmailauthNeeded {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): steamAuth invalid for %s, "+
			" code sent via email", username)
	}

	if !loginBody.Success {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): unknown error for %s", username)
	}

	for _, cookie := range resp.Cookies() {

		if cookie.Name == "steamLogin" {
			steamLogin = cookie.Value
		}
		if cookie.Name == "steamLoginSecure" {
			steamLoginSecure = cookie.Value
		}
	}

	if steamLogin == "" {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): steamLogin cookie is empty")
	}

	if steamLoginSecure == "" {
		return sessionid, steamLogin, steamLoginSecure, fmt.Errorf("auth DoLogin(): steamLoginSecure cookie is empty")
	}

	return sessionid, steamLogin, steamLoginSecure, nil
}

// IsLoggedIn checks if a user is logged in or not
func IsLoggedIn(baseSteamWebURL string, client *http.Client) (bool, error) {
	resp, err := client.Get(baseSteamWebURL + IsLoggedInURL)
	if err != nil {
		return false, fmt.Errorf("auth IsLoggedin(): %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
