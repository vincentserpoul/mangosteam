package tradeoffer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/user"
)

// State represents the state of the tradeoffer, see constants
type State uint

const (
	// ETradeOfferStateCreated /!\ non steam status, used to know the TO has been created
	ETradeOfferStateCreated State = iota
)

// SteamTradeOfferID is the identifier of the tradeoffer within steam network
type SteamTradeOfferID uint64

// String will turn a steamID into a string
func (steamTradeOfferID SteamTradeOfferID) String() string {
	return strconv.FormatUint(uint64(steamTradeOfferID), 10)
}

const (
	newTradeOfferSendURL        string = "/tradeoffer/new/send"
	newTradeOfferSendRefererURL string = "/tradeoffer/new/?partner="
)

// Result is the response body from the tradeoffer create request
type Result struct {
	TradeOfferID SteamTradeOfferID `json:",string"`
	Error        string            `json:"strError"`
}

// CreateSteamTradeOffer sends a new trade offer to the given Steam user.
func CreateSteamTradeOffer(
	baseSteamWebURL string,
	client *http.Client,
	sessionID string,
	otherSteamID mangosteam.SteamID,
	accessToken string,
	myItems, theirItems []*Asset,
	message string,
) (*Result, error) {

	baseURL, _ := url.Parse(baseSteamWebURL + newTradeOfferSendURL)

	tradeOfferJSON, err := getJSONTradeOffer(myItems, theirItems)

	if err != nil {
		return nil, err
	}

	tradeOfferCreateParamsJSON, err := getTradeOfferCreateParams(accessToken)

	if err != nil {
		return nil, err
	}

	bodyTradeOffer := getBodyTradeOffer(
		sessionID, otherSteamID, tradeOfferJSON, tradeOfferCreateParamsJSON,
		message)

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(bodyTradeOffer))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	// Headers
	referer := baseSteamWebURL + newTradeOfferSendRefererURL + otherSteamID.GetAccountID()
	req.Header.Add("Referer", referer)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// If we failed, error out
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil,
			fmt.Errorf("CreateSteamTradeOffer: status code %d. message: %s", resp.StatusCode, body)
	}

	// Load the JSON result into Result
	result := new(Result)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func getJSONTradeOffer(myItems, theirItems []*Asset) ([]byte, error) {

	var tradeOfferJSON []byte
	var err error

	// json_tradeoffer
	tradeOffer := map[string]interface{}{
		"newversion": true,
		"version":    2,
		"me": map[string]interface{}{
			"assets":   myItems,
			"currency": make([]struct{}, 0),
			"ready":    false,
		},
		"them": map[string]interface{}{
			"assets":   theirItems,
			"currency": make([]struct{}, 0),
			"ready":    false,
		},
	}

	// create the json payload
	tradeOfferJSON, err = json.Marshal(tradeOffer)

	if err != nil {
		return tradeOfferJSON, err
	}

	return tradeOfferJSON, nil
}

func getTradeOfferCreateParams(accessToken string) ([]byte, error) {
	var paramsJSON []byte
	var err error

	// trade_offer_create_params
	params := make(map[string]string)

	if accessToken != "" {
		params["trade_offer_access_token"] = accessToken
	}

	paramsJSON, err = json.Marshal(params)
	if err != nil {
		return paramsJSON, err
	}

	return paramsJSON, nil
}

func getBodyTradeOffer(
	sessionID string,
	otherSteamID mangosteam.SteamID,
	tradeOfferJSON []byte,
	paramsJSON []byte,
	message string,
) string {

	form := url.Values{}
	form.Add("sessionid", sessionID)
	form.Add("serverid", "1")
	form.Add("partner", otherSteamID.String())
	form.Add("tradeoffermessage", message)
	form.Add("captcha", "")
	form.Add("trade_offer_create_params", string(paramsJSON))
	form.Add("json_tradeoffer", string(tradeOfferJSON))

	return form.Encode()
}

// CreateCurlSteamTradeOffer creates a curl tradeoffer, mostly for simple tests
func CreateCurlSteamTradeOffer(
	baseSteamWebURL string,
	otherSteamID mangosteam.SteamID,
	user *steamuser.User,
	assetID AssetID,
	accessToken string,
) (string, error) {

	var userAsset []*Asset
	asset := Asset{AssetID: assetID}
	asset.Defaults(730)

	tradeOfferJSON, err := getJSONTradeOffer(userAsset, []*Asset{&asset})
	tradeOfferCreateParamsJSON, err := getTradeOfferCreateParams(accessToken)

	bodyTradeOffer := getBodyTradeOffer(
		"1",
		otherSteamID,
		tradeOfferJSON,
		tradeOfferCreateParamsJSON,
		"test curl user",
	)
	if err != nil {
		return "", nil
	}

	curlString := "curl '" + baseSteamWebURL + newTradeOfferSendURL + "'" +
		" -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8'" +
		" -H 'Referer: " + baseSteamWebURL + newTradeOfferSendRefererURL + otherSteamID.GetAccountID() +
		"'" + " -H 'Cookie: " +
		"steamMachineAuth" + user.SteamID.String() + "=" + user.SteamMachineAuth + "; " +
		"sessionid=1; " +
		"steamLogin=" + user.SteamLogin + "; " +
		"steamLoginSecure=" + user.SteamLoginSecure + ";'" +
		" --data '" + bodyTradeOffer + "'"

	return curlString, nil
}
