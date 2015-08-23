package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoLogin(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockKOLoginDologin())
	}))
	defer ts.Close()

	client := http.Client{}
	username := "mangosteam"
	encryptedPassword := "123"
	rsatimestamp := "123"
	emailauthKeyedIn := ""
	captchaGID := ""
	captchaKeyedIn := ""

	err := DoLogin(
		ts.URL,
		&client,
		username,
		encryptedPassword,
		rsatimestamp,
		emailauthKeyedIn,
		captchaGID,
		captchaKeyedIn,
	)

	if err == nil {
		t.Errorf("Dologin returns no error when login is not successful")
	}

	return
}
