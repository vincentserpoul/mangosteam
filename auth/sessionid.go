package auth

import (
	"fmt"
	"net/http"
)

func getNewSessionID(baseSteamWebURL string) (string, error) {

	resp, err := http.Get(baseSteamWebURL)
	if err != nil {
		return "", fmt.Errorf("auth getNewSessionID(%s): %v", baseSteamWebURL, err)
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "sessionid" {
			return cookie.Value, nil
		}
	}

	return "", fmt.Errorf("auth getNewSessionID(%s) contains no sessionid cookie", baseSteamWebURL)
}
