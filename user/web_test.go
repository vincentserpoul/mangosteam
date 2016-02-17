package steamuser

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestOKNewWebSteamClient(t *testing.T) {
	user := User{
		SteamID:           mangosteam.SteamID(123456789),
		Username:          "1",
		Password:          "1",
		APIKey:            "1",
		Email:             "1",
		SteamLogin:        "123",
		SteamLoginSecure:  "1234",
		SteamGuardAccount: getTestSteamGuardAccount(),
		OAuth:             getTestOAuth(),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()
	client := user.NewWebSteamClient(ts.URL)

	u, _ := url.Parse(ts.URL)

	foundSessionIDCookie := false
	foundSteamLoginCookie := false
	foundSteamLoginSecureCookie := false
	for _, cookie := range client.Jar.Cookies(u) {
		if cookie.Name == "sessionid" && cookie.Value == "1" {
			foundSessionIDCookie = true
		}
		if cookie.Name == "steamLogin" && cookie.Value == "123" {
			foundSteamLoginCookie = true
		}
		if cookie.Name == "steamLoginSecure" && cookie.Value == "1234" {
			foundSteamLoginSecureCookie = true
		}
	}

	if !foundSessionIDCookie || !foundSteamLoginCookie || !foundSteamLoginSecureCookie {
		t.Errorf("NewWebSteamClient(%s) doesn't contain the cookies it should contain: \nsessionid(%t), steamLogin(%t), steamLoginSecure(%t)",
			ts.URL, foundSessionIDCookie, foundSteamLoginCookie, foundSteamLoginSecureCookie)
	}
}
