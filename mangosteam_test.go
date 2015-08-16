package mangosteam

import "testing"

func TestString(t *testing.T) {
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
