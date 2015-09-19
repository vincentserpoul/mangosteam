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

func TestLoginErrLoggedIn(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		time.Sleep(200 * time.Millisecond)
	})
	ts := httptest.NewServer(testMux)
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()

	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "1", "1", "1", "1"}
	err := user.Login(ts.URL)
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestTimeOutLoginGetRSAKey(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})

	testMux.HandleFunc(auth.GetRSAKeyURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	})
	ts := httptest.NewServer(testMux)
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()

	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "1", "1", "1", "1"}
	err := user.Login(ts.URL)
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestLoginGetRSAKey(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})

	testMux.HandleFunc(auth.DoLoginURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockKOLoginGetrsakey())
	})
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "1", "1", "1", "1"}
	err := user.Login(ts.URL)
	if err == nil {
		t.Errorf("Dologin returns no error when login failed, %v", err)
	}
	return
}

func TestLoginEncryptPasswordFail(t *testing.T) {
	testMux := http.NewServeMux()
	// To force the path after isLoggedIn
	testMux.HandleFunc(auth.IsLoggedInURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})
	testMux.HandleFunc(auth.DoLoginURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	})
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	user := User{mangosteam.SteamID(123456789), "1", "1", "1", "1", "", "1", "1", "1"}
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

func TestGetMockKOLoginGetrsakey(t *testing.T) {
	s := getMockKOLoginGetrsakey()

	if len(s) == 0 {
		t.Errorf("getMockKOLoginGetrsakey has an error, please check your mock")
	}
}
