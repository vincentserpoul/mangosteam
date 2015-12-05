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
		SteamID:          mangosteam.SteamID(123456789),
		SteamMachineAuth: "1",
		SteamLogin:       "1",
		SteamLoginSecure: "1",
		Username:         "1",
		Password:         "1",
		APIKey:           "1",
		Email:            "1",
		LastSessionID:    "1",
	}
}

func TestOKGetAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockExistingAPIKeyPage())
	}))
	defer ts.Close()
	_, err := user.getAPIKey(ts.URL)
	if err != nil {
		t.Errorf("Get APIKey should be successful %v", err)
	}
	return
}

func TestAccessDeniedGetAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockAccessDeniedgetAPIKey())
	}))
	defer ts.Close()
	_, err := user.getAPIKey(ts.URL)
	if err == nil {
		t.Errorf("Access denied should not be successful, %v", err)
	}
	return
}

func TestOKregisterAPIKey(t *testing.T) {
	user := getTestUser()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginDologin())
	}))
	defer ts.Close()
	err := user.registerAPIKey(ts.URL)
	if err != nil {
		t.Errorf("Register APIKey should be successful %v", err)
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
	err := user.registerAPIKey(ts.URL)
	if err == nil {
		t.Errorf("DoClient should return error")
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
	_, err := user.getAPIKey(ts.URL)
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
	err := user.revokeAPIKey(ts.URL)
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
	err := user.revokeAPIKey(ts.URL)
	if err == nil {
		t.Errorf("Do RevokeAPIKey returns no error when Do PostForm failed")
	}
	return
}
