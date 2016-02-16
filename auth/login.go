package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vincentserpoul/mangosteam"
)

// doLogInResponse is used to store the body of the steam doLogin response
type doLogInResponse struct {
	Success         bool   `json:"success"`
	LoginComplete   bool   `json:"login_complete"`
	CaptchaNeeded   bool   `json:"captcha_needed"`
	CaptchaGID      string `json:"captcha_gid"`
	EmailSteamID    uint64 `json:"emailsteamid,string"`
	EmailAuthNeeded bool   `json:"emailauth_needed"`
	TwoFactorNeeded bool   `json:"requires_twofactor"`
	Message         string `json:"message"`
	OAuth           OAuth  `json:"transfer_parameters"`
}

// OAuth is what will be used in future requests to specify we are logged in
type OAuth struct {
	SteamID       mangosteam.SteamID `json:"steamid,string"`
	OAuthToken    string             `json:"auth"`
	Token         string             `json:"token"`
	TokenSecure   string             `json:"token_secure"`
	LastSessionID string             `json:"webcookie"`
}

const (
	// DoLoginURL URL used for login
	DoLoginURL string = "/login/dologin"
	// IsLoggedInURL URL used to check if user is logged in
	IsLoggedInURL string = "/actions/GetNotificationCounts"
)

// ErrTwoFactorNeeded is returned when the user needs to give a two factor code
var ErrTwoFactorNeeded = errors.New("Two factor auth needed")

// ErrEmailAuthNeeded is returned when the login requires an email code
var ErrEmailAuthNeeded = errors.New("Email auth needed")

// ErrCaptchaNeededNeeded is returned when the login requires a captcha
var ErrCaptchaNeededNeeded = errors.New("Captcha needed")

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
	twoFactorCode string,
) (OAuth, error) {
	var oAuth OAuth

	baseURL, _ := url.Parse(baseSteamWebURL + DoLoginURL)

	// default value set to -1
	if captchaGID == "" {
		captchaGID = "-1"
	}

	// adding query params
	params := url.Values{}
	params.Add("password", encryptedPassword)
	params.Add("username", username)
	params.Add("twofactorcode", twoFactorCode)
	params.Add("emailauth", emailauthKeyedIn)
	params.Add("loginfriendlyname", "")
	params.Add("captchagid", captchaGID)
	params.Add("captcha_text", captchaKeyedIn)
	params.Add("emailsteamid", "")
	params.Add("rsatimestamp", rsatimestamp)

	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", baseURL.String(), nil)

	resp, err := client.Do(req)
	if err != nil {
		return oAuth, fmt.Errorf("auth DoLogin(): %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return oAuth, fmt.Errorf("auth DoLogin(): bad request %v for %s, ",
			resp.Status, username)
	}

	loginBody := new(doLogInResponse)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(loginBody)

	if err != nil {
		return oAuth, fmt.Errorf("auth DoLogin(): %v", err)
	}

	if loginBody.EmailAuthNeeded {
		return oAuth, ErrEmailAuthNeeded
	}

	if !loginBody.Success {
		if loginBody.TwoFactorNeeded {
			return oAuth, ErrTwoFactorNeeded
		}
		return oAuth, fmt.Errorf("auth DoLogin(): unknown error for %s", username)
	}

	return loginBody.OAuth, nil
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
