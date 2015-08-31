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
	user.Login(ts.URL)
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
	user.Login(ts.URL)
	return
}
