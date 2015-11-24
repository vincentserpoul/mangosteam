package tradeoffer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/vincentserpoul/mangosteam"
)

const (
	cancelTradeOfferURL        string = "/tradeoffer/%d/cancel"
	cancelTradeOfferRefererURL string = "/profiles/%d/tradeoffers/sent"
)

// CancelSteamTradeOffer will cancel the specific tradeoffer, make sure the client is the right steam account
func CancelSteamTradeOffer(
	baseSteamWebURL string,
	client *http.Client,
	sessionID string,
	creatorSteamID mangosteam.SteamID,
	steamTradeOfferID SteamTradeOfferID,
) error {

	req, err := getCancelSteamTradeOfferRequest(
		baseSteamWebURL,
		sessionID,
		creatorSteamID,
		steamTradeOfferID,
	)

	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// If we failed, error out
	if resp.StatusCode != 200 && resp.StatusCode != 500 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("CancelSteamTradeOffer: status code %d. message: %s", resp.StatusCode, body)
	}

	// Load the JSON result into Result
	result := new(Result)

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)

	if err != nil {
		return err
	}

	// 16 means the TO has been cancelled, 11 means it has already been cancelled in the past
	if result.Success != 16 && result.Success != 11 && result.TradeOfferID == 0 {
		return fmt.Errorf("CancelSteamTradeOffer: status code %d. message: %#v", resp.StatusCode, result)
	}

	return nil

}

func getCancelSteamTradeOfferRequest(
	baseSteamWebURL string,
	sessionID string,
	creatorSteamID mangosteam.SteamID,
	steamTradeOfferID SteamTradeOfferID,
) (*http.Request, error) {
	if (baseSteamWebURL == "") || (sessionID == "") || (creatorSteamID.String() == "") || (steamTradeOfferID.String() == "") {
		return nil, fmt.Errorf("getCancelSteamTradeOfferRequest: Empty baseSteamURL or sessionID or creatorSteamID or steamTradeOfferID")
	}
	baseURL, _ := url.Parse(baseSteamWebURL + fmt.Sprintf(cancelTradeOfferURL, steamTradeOfferID))

	form := url.Values{}
	form.Add("sessionid", sessionID)

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	// Referer
	referer := baseSteamWebURL + fmt.Sprintf(cancelTradeOfferRefererURL, creatorSteamID)
	req.Header.Add("Referer", referer)

	return req, nil
}
