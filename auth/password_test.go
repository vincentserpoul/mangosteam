package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestGetRSAKey(t *testing.T) {
	username := "mangosteam"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"success":true,"publickey_mod":"BAC1D6C489FE24770D198A3F06903080902195934FCF2F0F44DB7DDF275B68F951520010F3B07BFDC9E339A1865A6C33257723493E52C84D81D2F488EB3EF0ABC3CB7617432CB52EB7ED934AE5B99241F936F651BAD78BD0A18EC6CC5F5211EC41814C9D9C5BC19E2267DA81DEA17B998E0D04FC346E1F3E7AFF657CD18296B3863A60E08B66B8E6446C6BD5C4417E5F75A9685E65E47A0A5AFC0168B66909FE4F894F93312528EC3F272D5D1587C0AD458F259198114CE4EF2E234B39D3311136679A4E8D37C56D28C4F506C3790C652F4DC84AD01F82C47AB15B1710255949EC897B552B577C0ED1EE63F26F649AD4EBED6A0A8874A3D4F517558B78D2DC69","publickey_exp":"010001","timestamp":"583882850000","steamid":"76561198063143983","token_gid":"8213f42ccc2b4c7"}`)
	}))
	defer ts.Close()

	key, err := GetRSAKey(ts.URL+"/", username)

	if err != nil {
		t.Errorf("WARNING GetRSAkey from http://127.0.0.1:9090 is not working anymore: %v", err)
		return
	}

	if key.PublicKeyExponent == "" ||
		key.PublicKeyModulus == "" ||
		key.SteamID == "" ||
		key.Timestamp == "" {
		t.Errorf("GetRSAkey from steam is not returning a complete key but %v instead", key)
	}
}

func TestGetRSAKeySteam(t *testing.T) {
	username := "mangosteam"
	baseSteamURL := "https://steamcommunity.com/"

	key, err := GetRSAKey(baseSteamURL, username)

	if err != nil {
		t.Errorf("WARNING GetRSAkey from steam is not working anymore or %s has been deactivated", username)
	}
	if key.PublicKeyExponent == "" ||
		key.PublicKeyModulus == "" ||
		key.SteamID == "" ||
		key.Timestamp == "" {
		t.Errorf("GetRSAkey from steam is not returning a complete key but %v instead", key)
	}
}
