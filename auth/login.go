package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/vincentserpoul/mangosteam"
)

// doLogInResponse is used to store the body of the steam doLogin response
type doLogInResponse struct {
	Success         bool `json:"success"`
	EmailauthNeeded bool `json:"emailauth_needed"`
}

// DoLogin is used to log in the steam account after we got the encrypted password
func DoLogin(
	client *http.Client,
	username string,
	encryptedPassword string,
	rsatimestamp string,
	emailauthKeyedIn string,
	captchaGID string,
	captchaKeyedIn string,
) error {

	baseURL, _ := url.Parse(mangosteam.BaseSteamWebURL + "login/dologin")

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

	baseURLCookie, _ := url.Parse(mangosteam.BaseSteamWebURL)

	logrus.WithFields(logrus.Fields{
		"username": username,
	}).Debug("doLogin: cookiesJar contains ", client.Jar.Cookies(baseURLCookie))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("auth doLogin(): %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("auth doLogin(): bad request %v for %s, ",
			resp.Status, username)
	}

	loginBody := new(doLogInResponse)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(loginBody)

	if err != nil {
		return fmt.Errorf("auth doLogin(): %v", err)
	}

	if loginBody.EmailauthNeeded {
		return fmt.Errorf("auth doLogin(): steamAuth invalid for %s, "+
			" code sent via email", username)
	}

	if !loginBody.Success {
		return fmt.Errorf("auth doLogin(): unknown error for %s", username)
	}

	return nil
}
