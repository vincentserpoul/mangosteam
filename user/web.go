package steamuser

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/vincentserpoul/mangosteam"
)

// NewWebSteamClient creates a web steam client with the right cookies
// to interact with web steam
func (user *User) NewWebSteamClient() *http.Client {
	client := &http.Client{}
	client.Jar, _ = cookiejar.New(new(cookiejar.Options))
	base, _ := url.Parse(mangosteam.BaseSteamWebURL)

	client.Jar.SetCookies(base, []*http.Cookie{
		&http.Cookie{
			Name:  "steamMachineAuth" + user.SteamID.String(),
			Value: user.SteamMachineAuth,
		},
		&http.Cookie{
			Name:  "steamLogin",
			Value: user.SteamLogin,
		},
		&http.Cookie{
			Name:  "steamLoginSecure",
			Value: user.SteamLoginSecure,
		},
		&http.Cookie{
			Name:  "sessionid",
			Value: user.LastSessionID,
		},
	})

	return client
}
