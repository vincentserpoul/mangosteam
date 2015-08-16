package tradeoffer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/user"
)

const (
	newTradeOfferSendURL        string = mangosteam.BaseSteamWebURL + "tradeoffer/new/send"
	newTradeOfferSendRefererURL string = mangosteam.BaseSteamWebURL + "tradeoffer/new/?partner="
)

// Result is the response body from the tradeoffer create request
type Result struct {
	TradeOfferID SteamTradeOfferID `json:",string"`
}

// CreateSteamTradeOffer sends a new trade offer to the given Steam user.
func CreateSteamTradeOffer(
	client *http.Client,
	sessionID string,
	otherSteamID mangosteam.SteamID,
	accessToken string,
	myItems, theirItems *[]Asset,
	message string,
) (*http.Request, *http.Response, *Result, error) {

	baseURL, _ := url.Parse(newTradeOfferSendURL)

	tradeOfferJSON, err := getJSONTradeOffer(myItems, theirItems)

	if err != nil {
		return nil, nil, nil, err
	}

	tradeOfferCreateParamsJSON, err := getTradeOfferCreateParams(accessToken)

	if err != nil {
		return nil, nil, nil, err
	}

	bodyTradeOffer := getBodyTradeOffer(
		sessionID, otherSteamID, tradeOfferJSON, tradeOfferCreateParamsJSON,
		message)

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(bodyTradeOffer))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	// Headers
	referer := newTradeOfferSendRefererURL + otherSteamID.GetAccountID()
	req.Header.Add("Referer", referer)

	dump, err := httputil.DumpRequest(req, true)

	fmt.Println(string(dump))

	resp, err := client.Do(req)
	defer resp.Body.Close()

	// If we failed, error out
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return req, resp, nil,
			fmt.Errorf("CreateSteamTradeOffer: status code %d. message: %s", resp.StatusCode, body)
	}

	// Load the JSON result into TradeTradeOfferResult
	result := new(Result)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)

	if err != nil {
		return req, resp, nil, err
	}

	return req, resp, result, nil
}

func getJSONTradeOffer(myItems, theirItems *[]Asset) ([]byte, error) {

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
	otherSteamID mangosteam.SteamID,
	user *steamuser.User,
	assetID AssetID,
	accessToken string,
) (string, error) {

	var userAsset []Asset
	asset := Asset{AssetID: assetID}
	asset.Defaults(730)

	tradeOfferJSON, err := getJSONTradeOffer(&userAsset, &[]Asset{asset})
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

	curlString := "curl '" + newTradeOfferSendURL + "'" +
		" -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8'" +
		" -H 'Referer: " + newTradeOfferSendRefererURL + otherSteamID.GetAccountID() + "'" +
		" -H 'Cookie: " +
		"steamMachineAuth" + user.SteamID.String() + "=" + user.SteamMachineAuth + "; " +
		"sessionid=1; " +
		"steamLogin=" + user.SteamLogin + "; " +
		"steamLoginSecure=" + user.SteamLoginSecure + ";'" +
		" --data '" + bodyTradeOffer + "'"

	return curlString, nil
}
