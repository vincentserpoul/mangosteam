package auth

// all props go to https://github.com/YellowOrWhite/go-steam-mobileauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
)

// SteamGuardAccount represents the data needed to generate a code
type SteamGuardAccount struct {
	SharedSecret   string `json:"shared_secret"`
	SerialNumber   string `json:"serial_number"`
	RevocationCode string `json:"revocation_code"`
	URI            string `json:"uri"`
	ServerTime     int64  `json:"server_time,string"`
	AccountName    string `json:"account_name"`
	TokenGID       string `json:"token_gid"`
	IdentitySecret string `json:"identity_secret"`
	Secret1        string `json:"secret_1"`
	Status         int32  `json:"status"`
	DeviceID       string `json:"device_id"`
	// Set to true if the authenticator has actually been applied to the account.
	FullyEnrolled bool `json:"fully_enrolled"`
}

// GenerateSteamGuardCode gives a steamgaurd code for login
func (a *SteamGuardAccount) GenerateSteamGuardCode(baseSteamAPIURL string) (string, error) {
	return a.GenerateSteamGuardCodeForTime(GetSteamTime(baseSteamAPIURL))
}

// GenerateSteamGuardCodeForTime according to the server time
func (a *SteamGuardAccount) GenerateSteamGuardCodeForTime(t int64) (string, error) {
	if a.SharedSecret == "" {
		return "", errors.New("shared secret not set")
	}

	// Shared secret is our key
	sharedSecretBytes, err := base64.StdEncoding.DecodeString(a.SharedSecret)
	if err != nil {
		return "", err
	}

	// Time for code
	t = t / 30 // TODO: why we are doing this???
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(t))

	// Generate hmac
	hmacGenerator := hmac.New(sha1.New, sharedSecretBytes)
	hmacGenerator.Write(timeBytes)
	mac := hmacGenerator.Sum(nil)

	// the last 4 bits of the mac say where the code starts
	// (e.g. if last 4 bit are 1100, we start at byte 12)
	start := int(mac[19] & 0x0f)

	// extract code - 4 bytes
	codeBytes := make([]byte, 4)
	copy(codeBytes, mac[start:])
	fullCode := binary.BigEndian.Uint32(codeBytes)
	fullCode = fullCode & 0x7fffffff

	// character set for authenticator code
	chars := []byte{50, 51, 52, 53, 54, 55, 56, 57, 66, 67, 68, 70, 71, 72, 74, 75, 77, 78, 80, 81, 82, 84, 86, 87, 88, 89}

	// build the alphanumeric code
	var textCodeBytes []byte
	for i := 0; i < 5; i++ {
		textCodeBytes = append(textCodeBytes, chars[fullCode%uint32(len(chars))])
		fullCode = fullCode / uint32(len(chars))
	}

	return string(textCodeBytes[:]), nil
}
