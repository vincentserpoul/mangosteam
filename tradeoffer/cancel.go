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

const cancelTradeOfferURL string = "/tradeoffer/%d/decline"
const cancelTradeOfferRefererURL string = "/profiles/%d/tradeoffers/"

// CancelTradeOffer will cancel the specific tradeoffer, make sure the client is the right steam account
func CancelTradeOffer(
	baseSteamWebURL string,
	client *http.Client,
	sessionID string,
	creatorSteamID mangosteam.SteamID,
	steamTradeOfferID SteamTradeOfferID,
) (*Result, error) {

	baseURL, _ := url.Parse(baseSteamWebURL + fmt.Sprintf(cancelTradeOfferURL, steamTradeOfferID))

	cancelTOBody := struct {
		SteamTradeOfferID
	}{
		steamTradeOfferID,
	}

	cancelTOBodyJSON, err := json.Marshal(cancelTOBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL.String(), strings.NewReader(string(cancelTOBodyJSON)))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	// Referer
	referer := baseSteamWebURL + fmt.Sprintf(cancelTradeOfferRefererURL, creatorSteamID)
	req.Header.Add("Referer", referer)

	if err != nil {
		return nil, err
	}

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
