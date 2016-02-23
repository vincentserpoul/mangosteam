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
	client.Jar, _ = cookiejar.New(nil)

	steamURL, _ := url.Parse(mangosteam.BaseSteamWebURL)

	var cookiesToBeSet []*http.Cookie

	if user.SteamLogin != "" {
		cookiesToBeSet = append(cookiesToBeSet,
			&http.Cookie{
				Name:  "steamLogin",
				Value: user.SteamLogin,
			},
		)
	}

	if user.SteamLoginSecure != "" {
		cookiesToBeSet = append(cookiesToBeSet,
			&http.Cookie{
				Name:  "steamLoginSecure",
				Value: user.SteamLoginSecure,
			},
		)
	}

	if user.LastSessionID != "" {
		cookiesToBeSet = append(cookiesToBeSet,
			&http.Cookie{
				Name:  "sessionid",
				Value: user.LastSessionID,
			},
		)
	}

	client.Jar.SetCookies(steamURL, cookiesToBeSet)

	return client
}
