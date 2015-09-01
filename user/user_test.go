package steamuser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestLoginGetRSAKey(t *testing.T) {
	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "1", "1", "1", "1"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockKOLoginGetrsakey())
	}))
	defer ts.Close()
	err := user.Login(ts.URL)
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestLoginEncryptPassword(t *testing.T) {
	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "", "1", "1", "1"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	}))
	defer ts.Close()
	err := user.Login(ts.URL)
	if err == nil {
		t.Errorf("Dologin returns no error when empty password, %v", err)
	}
	return
}

func TestLoginDoLogin(t *testing.T) {
	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "1", "1", "1", "1"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	}))
	defer ts.Close()
	err := user.Login(ts.URL)
	if err == nil {
		t.Errorf("Dologin returns when status not found error , %v", err)
	}
	return
}

func TestOKLogin(t *testing.T) {
	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "1", "1", "1", "1"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	}))
	defer ts.Close()
	err := user.Login(ts.URL)
	if err != nil {
		t.Errorf("Dologin should not failed")
	}
	return
}
