package steamuser

import (
	"net/url"
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestOKNewWebSteamClient(t *testing.T) {
	baseSteamWebURL := "http://www.test.com"

	user := User{
		SteamID:          mangosteam.SteamID(76561198264159435),
		Username:         "github_mangosteam",
		Password:         "mangosteam_test",
		LastSessionID:    "1",
		SteamLogin:       "123",
		SteamLoginSecure: "1234",
		SteamMachineAuth: "12345",
	}

	client := user.NewWebSteamClient(baseSteamWebURL)

	u, _ := url.Parse(baseSteamWebURL)

	foundSessionIDCookie := false
	foundSteamLoginCookie := false
	foundSteamLoginSecureCookie := false
	foundSteamMachineAuthCookie := false
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
		if cookie.Name == "steamMachineAuth"+user.SteamID.String() && cookie.Value == "12345" {
			foundSteamMachineAuthCookie = true
		}
	}

	if !foundSessionIDCookie || !foundSteamLoginCookie || !foundSteamLoginSecureCookie || !foundSteamMachineAuthCookie {
		t.Errorf("NewWebSteamClient(%s) doesn't contain the cookies it should contain: \nsessionid(%t), steamLogin(%t), steamLoginSecure(%t),steamMachineAuth(%t)",
			baseSteamWebURL, foundSessionIDCookie, foundSteamLoginCookie, foundSteamLoginSecureCookie, foundSteamMachineAuthCookie)
	}
}
