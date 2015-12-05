package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestKOGetNewSessionID(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()
	ts.Config.WriteTimeout = 20 * time.Millisecond

	_, err := getNewSessionID(ts.URL)

	if err == nil {
		t.Errorf("getNewSessionID should return an error if the server is down")
	}

}

func TestMissingSessionIDGetNewSessionID(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	_, err := getNewSessionID(ts.URL)

	if err == nil {
		t.Errorf("getNewSessionID should return an error if sessionid is missing")
	}

}

func TestOKSessionIDGetNewSessionID(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: `sessionid`, Value: `1234`})
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	sessionid, err := getNewSessionID(ts.URL)

	if err != nil {
		t.Errorf("getNewSessionID returns an error whereas it shouldn't")
	}
	if sessionid != `1234` {
		t.Errorf("getNewSessionID should return `1234` but returned `%s`", sessionid)
	}

}
