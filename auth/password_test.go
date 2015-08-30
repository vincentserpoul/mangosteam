package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEncryptPasswordEmptyPassword(t *testing.T) {
	testRSAKey := &RSAKey{
		"123456789",
		"123456789",
		"123456789",
		"123456789",
	}
	_, err := EncryptPassword("", testRSAKey)

	if err == nil {
		t.Error("Encrypt password doesn't return error when password is empty")
	}
}

func TestEncryptPasswordNilRSAKey(t *testing.T) {
	password := "testpass"
	_, err := EncryptPassword(password, nil)

	if err == nil {
		t.Errorf("Encrypt password doesn't return error when RSAKey is nil")
	}
}

func TestEncryptPasswordExpTooBig(t *testing.T) {
	testRSAKey := &RSAKey{
		"123456789123456789123456789123456789123456789",
		"123456789",
		"123456789",
		"123456789",
	}
	password := "testpassword"
	_, err := EncryptPassword(password, testRSAKey)

	if err == nil {
		t.Errorf("EncryptPassword() doesn't returns error whereas PublicKeyExponent %s is too big", testRSAKey.PublicKeyExponent)
	}
}

func TestEncryptPasswordBadMod(t *testing.T) {
	testRSAKey := &RSAKey{
		"123456789123456789123456789123456789",
		"1234",
		"123456789",
		"123456789",
	}
	password := "testpassword"
	_, err := EncryptPassword(password, testRSAKey)

	if err == nil {
		t.Errorf("EncryptPassword() doesn't returns error whereas PublicKeyModulus %s is too small", testRSAKey.PublicKeyModulus)
	}
}

func TestEncryptPasswordExpNotInt(t *testing.T) {
	testRSAKey := &RSAKey{
		"123456789123456789123456789123456789",
		"12FG",
		"123456789",
		"123456789",
	}
	password := "testpassword"
	_, err := EncryptPassword(password, testRSAKey)

	if err == nil {
		t.Errorf("EncryptPassword() doesn't returns error whereas PublicKeyModulus %s is not int parsable", testRSAKey.PublicKeyExponent)
	}
}

func TestEncryptPassword(t *testing.T) {
	testRSAKey := &RSAKey{
		"123456789123456789123456789123456789123456789",
		"1234",
		"123456789",
		"123456789",
	}
	password := "testpassword"

	_, err := EncryptPassword(password, testRSAKey)

	if err != nil {
		t.Errorf("EncryptPassword(%v, %v) returns error %v", password, testRSAKey, err)
	}
}

func TestExtractRSAKeyFromJSON(t *testing.T) {
	wantedRSAKey := &RSAKey{
		"123456789123456789123456789123456789123456789",
		"1234",
		"123456789",
		"123456789",
	}

	testJSONString := "{\"publickey_mod\":\"123456789123456789123456789123456789123456789\",\"publickey_exp\":\"1234\",\"timestamp\":\"123456789\",\"steamid\":\"123456789\"}"

	gotRSAkey, err := extractRSAKeyFromJSON([]byte(testJSONString))

	if err != nil {
		t.Errorf("extractRSAKeyFromJSON(%v) returns error %v", testJSONString, err)
	} else if *gotRSAkey != *wantedRSAKey {
		t.Errorf("extractRSAKeyFromJSON(%v) ==  %v, wanted %v", testJSONString, gotRSAkey, wantedRSAKey)
	}
}

func TestExtractRSAKeyFromJSONMissingField(t *testing.T) {

	testJSONString := "{\"publickey_exp\": \"1234\", \"timestamp\": \"123456789\", \"steamid\":\"123456789\"}"

	_, err := extractRSAKeyFromJSON([]byte(testJSONString))

	if err == nil {
		t.Errorf("extractRSAKeyFromJSON(%v)  doesn't trigger an error whereas the JSON string is missing a field", testJSONString)
	}
}

func TestExtractRSAKeyFromJSONBadJSON(t *testing.T) {

	testJSONString := "{publickey_exp\": \"1234\", \"timestamp\": \"123456789\", \"steamid\":\"123456789\"}"

	_, err := extractRSAKeyFromJSON([]byte(testJSONString))

	if err == nil {
		t.Errorf("extractRSAKeyFromJSON(%v)  doesn't trigger an error whereas the JSON string is wrong", testJSONString)
	}
}

func TestMockOKGetRSAKey(t *testing.T) {
	username := "mangosteam"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKLoginGetrsakey())
	}))
	defer ts.Close()

	key, err := GetRSAKey(ts.URL, username)

	if err != nil {
		t.Errorf("GetRSAkey httptest is failing: %v", err)
	}

	if key.PublicKeyExponent == "" ||
		key.PublicKeyModulus == "" ||
		key.SteamID == "" ||
		key.Timestamp == "" {
		t.Errorf("GetRSAkey from mocks is not returning a complete key but %v instead", key)
	}
}

func TestMockKOSteamGetRSAKey(t *testing.T) {
	username := "mangosteam"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockKOLoginGetrsakey())
	}))
	defer ts.Close()

	_, err := GetRSAKey(ts.URL, username)

	if err == nil {
		t.Errorf("GetRSAkey httptest should failing")
	}
}

func TestMockKOGetRSAKey(t *testing.T) {
	username := "mangosteam"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	key, err := GetRSAKey(ts.URL, username)

	if key != nil || err == nil {
		t.Errorf("GetRSAkey failing but not showing errors")
	}
}

func TestGetMockOKLoginGetrsakey(t *testing.T) {
	str := getMockOKLoginGetrsakey()

	if len(str) == 0 {
		t.Errorf("MockOKLoginGetrsakey is not working anymore")
	}
}

func TestGetMockKOLoginGetrsakey(t *testing.T) {
	str := getMockKOLoginGetrsakey()

	if len(str) == 0 {
		t.Errorf("getMockKOLoginGetrsakey is not working anymore")
	}
}
func TestKODoLoginPostForm(t *testing.T) {
	username := "mangosteam"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// An error is returned if caused by client policy (such as CheckRedirect),
		// or if there was an HTTP protocol error. A non-2xx response doesn't cause an error.
		time.Sleep(200 * time.Millisecond)
	}))
	ts.Config.ReadTimeout = 20 * time.Millisecond
	ts.Config.WriteTimeout = 20 * time.Millisecond
	defer ts.Close()
	_, err := GetRSAKey(ts.URL, username)
	if err == nil {
		t.Errorf("Dologin returns no error when DoLogin PostForm failed")
	}

	return
}
