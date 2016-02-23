package steamuser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/auth"
)

func getTestSteamGuardAccount() auth.SteamGuardAccount {
	return auth.SteamGuardAccount{
		SharedSecret:   "1",
		SerialNumber:   "1",
		RevocationCode: "1",
		URI:            "1",
		ServerTime:     0,
		AccountName:    "1",
		TokenGID:       "1",
		IdentitySecret: "1",
		Secret1:        "1",
		Status:         0,
		DeviceID:       "1",
		FullyEnrolled:  true,
	}
}

func getTestOAuth() auth.OAuth {
	return auth.OAuth{
		SteamID:       mangosteam.SteamID(123456789),
		OAuthToken:    "1",
		Token:         "1",
		TokenSecure:   "1",
		LastSessionID: "1",
	}
}
func TestLoginErrLoggedIn(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		time.Sleep(200 * time.Millisecond)
	})
	ts := httptest.NewServer(testMux)
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()
	mangosteam.BaseSteamAPIURL = ts.URL
	mangosteam.BaseSteamWebURL = ts.URL

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

	err := user.Login()
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestTimeOutLoginGetRSAKey(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})

	testMux.HandleFunc(auth.GetRSAKeyURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	})
	ts := httptest.NewServer(testMux)
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()
	mangosteam.BaseSteamAPIURL = ts.URL
	mangosteam.BaseSteamWebURL = ts.URL

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

	err := user.Login()
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestLoginGetRSAKey(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})

	testMux.HandleFunc(auth.DoLoginURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockKOLoginGetrsakey())
	})
	ts := httptest.NewServer(testMux)
	defer ts.Close()
	mangosteam.BaseSteamAPIURL = ts.URL
	mangosteam.BaseSteamWebURL = ts.URL

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
	err := user.Login()
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestLoginEncryptPasswordFail(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})
	testMux.HandleFunc(auth.DoLoginURI, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	})
	ts := httptest.NewServer(testMux)
	defer ts.Close()
	mangosteam.BaseSteamAPIURL = ts.URL
	mangosteam.BaseSteamWebURL = ts.URL

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

	err := user.Login()
	if err == nil {
		t.Errorf("Dologin returns no error when empty password, %v", err)
	}
	return
}

func TestLoginDoLogin(t *testing.T) {
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	}))
	defer ts.Close()
	mangosteam.BaseSteamAPIURL = ts.URL
	mangosteam.BaseSteamWebURL = ts.URL

	err := user.Login()
	if err == nil {
		t.Errorf("Dologin returns when status not found error , %v", err)
	}
	return
}

func TestOKLogin(t *testing.T) {
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
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	}))
	defer ts.Close()
	mangosteam.BaseSteamAPIURL = ts.URL
	mangosteam.BaseSteamWebURL = ts.URL

	err := user.Login()
	if err != nil {
		t.Errorf("Dologin should not failed")
	}
	return
}

func TestGetMockKOLoginGetrsakey(t *testing.T) {
	s := getMockKOLoginGetrsakey()

	if len(s) == 0 {
		t.Errorf("getMockKOLoginGetrsakey has an error, please check your mock")
	}
}
