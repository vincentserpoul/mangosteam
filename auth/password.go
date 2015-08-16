package auth

// RSAKey is the encoding params for the auth
import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/vincentserpoul/mangosteam"
)

// RSAKey represent the key for encoding password
type RSAKey struct {
	PublicKeyModulus  string `json:"publickey_mod"`
	PublicKeyExponent string `json:"publickey_exp"`
	Timestamp         string `json:"timestamp"`
	SteamID           string `json:"steamid"`
}

// GetRSAKey queries steam to get the key needed to encode the password
func GetRSAKey(username string) (*RSAKey, error) {

	logrus.WithFields(logrus.Fields{
		"username": username,
	}).Debug("calling ", mangosteam.BaseSteamWebURL+"login/getrsakey")

	resp, err := http.PostForm(mangosteam.BaseSteamWebURL+"login/getrsakey",
		url.Values{"username": {username}})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	key, err := extractRSAKeyFromJSON(body)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"username": username,
	}).Debug("got RSA key from Steam")

	return key, nil
}

func extractRSAKeyFromJSON(JSONBytes []byte) (*RSAKey, error) {
	key := RSAKey{}
	if err := json.Unmarshal(JSONBytes, &key); err != nil {
		return nil, err
	}

	if key.PublicKeyExponent == "" ||
		key.PublicKeyModulus == "" ||
		key.SteamID == "" ||
		key.Timestamp == "" {
		return nil, fmt.Errorf(
			"websteam extractRSAKeyFromJSON: incomplete RSAKey unmarshalled")
	}

	return &key, nil
}

// EncryptPassword will RSA encode the password with the key provided
func EncryptPassword(password string, rsaKey *RSAKey) (string, error) {

	if password == "" {
		return "", fmt.Errorf("websteam EncryptPassword(password: %v, rsaKey: %v) "+
			"requires a non empty password as an argument", password, rsaKey)
	}
	if rsaKey == nil {
		return "", fmt.Errorf("websteam EncryptPassword(password: %v, rsaKey: %v) "+
			"requires a non empty rsaKey as an argument", password, rsaKey)
	}

	// convert the hex string to int
	pubkeyExpInt, err := strconv.ParseInt(rsaKey.PublicKeyExponent, 16, 64)

	if err != nil {
		return "", err
	}

	// convert the hex string to a big int (we can't use ParseInt)
	pubkeyModInt := new(big.Int)
	pubkeyModInt.SetString(rsaKey.PublicKeyModulus, 16)

	var realKey rsa.PublicKey

	realKey.E = int(pubkeyExpInt)
	realKey.N = pubkeyModInt

	encryptedPass, err := rsa.EncryptPKCS1v15(rand.Reader, &realKey, []byte(password))

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedPass), nil
}
