package steamuser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vincentserpoul/mangosteam"
)

func getTestUser() User {
	return User{
		SteamID:  mangosteam.SteamID(123456789),
		Username: "1",
		Password: "1",
		APIKey:   "1",
		Email:    "1",
	}
}

func TestOKGetAPIKey(t *testing.T) {
	user := getTestUser()
	expectedAPIKey := "01234567890123456789012345678901"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockExistingAPIKeyPage())
	}))
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	APIKey, err := user.GetAPIKey()
	if err != nil {
		t.Errorf("Get APIKey should be successful %v", err)
	}

	if APIKey != expectedAPIKey {
		t.Errorf("Register APIKey, expected `%s`, got returned key `%s`", expectedAPIKey, APIKey)
	}

	if user.APIKey != expectedAPIKey {
		t.Errorf("Register APIKey, expected %s, got user key %s", expectedAPIKey, user.APIKey)
	}

}

func TestOKGetAPIKeyNonRegistered(t *testing.T) {
	user := getTestUser()
	expectedAPIKey := "01234567890123456789012345678901"

	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(getAPIKeyURL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockNonExistingAPIKeyPage())
	})

	testMux.HandleFunc(registerAPIKeyURL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockExistingAPIKeyPage())
	})
	ts := httptest.NewServer(testMux)
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	APIKey, err := user.GetAPIKey()
	if err != nil {
		t.Errorf("Get APIKey should be successful %v", err)
	}

	if APIKey != expectedAPIKey {
		t.Errorf("Register APIKey, expected `%s`, got returned key `%s`", expectedAPIKey, APIKey)
	}

	if user.APIKey != expectedAPIKey {
		t.Errorf("Register APIKey, expected %s, got user key %s", expectedAPIKey, user.APIKey)
	}
}

func TestAccessDeniedGetAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockAccessDeniedGetAPIKey())
	}))
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	_, err := user.GetAPIKey()
	if err == nil {
		t.Errorf("Access denied should not be successful, %v", err)
	}
	return
}

func TestOKRegisterAPIKey(t *testing.T) {
	user := getTestUser()
	expectedAPIKey := "01234567890123456789012345678901"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockExistingAPIKeyPage())
	}))
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	APIKey, err := user.registerAPIKey()
	if err != nil {
		t.Errorf("Register APIKey should be successful %v", err)
	}

	if APIKey != expectedAPIKey {
		t.Errorf("Register APIKey, expected `%s`, got returned key `%s`", expectedAPIKey, APIKey)
	}

	return
}

func TestKOClient(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// An error is returned if caused by client policy (such as CheckRedirect),
		// or if there was an HTTP protocol error. A non-2xx response doesn't cause an error.
		time.Sleep(200 * time.Millisecond)
	}))
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()

	mangosteam.BaseSteamWebURL = ts.URL

	_, err := user.registerAPIKey()
	if err == nil {
		t.Errorf("registerAPIKey should return error when timeout")
	}

	return
}
func TestEmptyAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockEmptyAPIKeyPage())
	}))
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	_, err := user.GetAPIKey()
	if err == nil {
		t.Errorf("Empty Apikey should return error")
	}
	return
}

func TestOKRevokeAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginDologin())
	}))
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	err := user.RevokeAPIKey()
	if err != nil {
		t.Errorf("RevokeAPIKey should be successful %v", err)
	}
	return
}

func TestKORevokeAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()
	mangosteam.BaseSteamWebURL = ts.URL

	err := user.RevokeAPIKey()
	if err == nil {
		t.Errorf("Do RevokeAPIKey returns no error when Do PostForm failed")
	}
	return
}
