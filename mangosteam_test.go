package mangosteam

import "testing"

func TestSteamIDString(t *testing.T) {
	steamID := SteamID(1234567890123)
	expectedStrSteamID := "1234567890123"
	gotStrSteamID := steamID.String()

	if gotStrSteamID != expectedStrSteamID {
		t.Errorf(
			"steamID %v .String(), expected %v, got %v",
			steamID, expectedStrSteamID, gotStrSteamID,
		)
	}
}

func TestContextIdString(t *testing.T) {
	contextID := ContextID(1234567890123)
	expectedContextID := "1234567890123"
	gotExpectedContextID := contextID.String()

	if gotExpectedContextID != expectedContextID {
		t.Errorf(
			"contextID %v , expected %v, got %v",
			contextID, expectedContextID, gotExpectedContextID,
		)
	}
}

func TestApIdString(t *testing.T) {
	appID := AppID(123456789)
	expectedAppID := "123456789"
	gotExpectedAppID := appID.String()

	if gotExpectedAppID != expectedAppID {
		t.Errorf(
			"appID %v, expected %v, got %v",
			appID, expectedAppID, gotExpectedAppID,
		)
	}
}

func TestGetAccountIdString(t *testing.T) {
	steamID := SteamID(1234567890123)
	expectedAccountID := "1912276171"
	GotExpectedAccountID := steamID.GetAccountID()

	if GotExpectedAccountID != expectedAccountID {
		t.Errorf(
			"Expected %v, got %v",
			expectedAccountID, GotExpectedAccountID,
		)
	}
}

func TestGetSteamIDFromString(t *testing.T) {

	expectedSteamIDFromString := SteamID(1234567890123)
	GotSteamIDFromString, _ := GetSteamIDFromString("1234567890123")

	if GotSteamIDFromString != expectedSteamIDFromString {
		t.Errorf(
			"Expected %v, got %v",
			expectedSteamIDFromString, GotSteamIDFromString,
		)
	}
}

func TestGetAppIDFromString(t *testing.T) {
	expectedGetAppIDFromString := AppID(123456789)
	GotGetAppIDFromString, _ := GetAppIDFromString("123456789")

	if GotGetAppIDFromString != expectedGetAppIDFromString {
		t.Errorf(
			"Expected %v, got %v",
			expectedGetAppIDFromString, GotGetAppIDFromString,
		)
	}
}

func TestErrorGetSteamIDFromString(t *testing.T) {

	expectedSteamIDFromString := SteamID(12345678901236)
	GotSteamIDFromString, err := GetSteamIDFromString("1234567890123A")

	if err != nil {
		t.Errorf(
			"Should be an error, Expected %v, got %v",
			expectedSteamIDFromString, GotSteamIDFromString,
		)
	}
}

func TestErrorGetAppIDFromString(t *testing.T) {
	expectedGetAppIDFromString := AppID(123456755)
	GotGetAppIDFromString, err := GetAppIDFromString("12345679A")

	if err != nil {
		t.Errorf(
			"Should be an error, Expected %v, got %v",
			expectedGetAppIDFromString, GotGetAppIDFromString,
		)
	}
}
