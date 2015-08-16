package auth

import "testing"

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
	username := "steambot_test"

	key, err := GetRSAKey(username)

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
