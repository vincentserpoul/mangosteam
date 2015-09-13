package tradeoffer

import (
	"io/ioutil"
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestSteamTradeOfferIDString(t *testing.T) {
	steamTradeOfferID := SteamTradeOfferID(123)
	expectedValue := "123"
	gotValue := steamTradeOfferID.String()
	if expectedValue != gotValue {
		t.Errorf(
			"SteamTradeOfferID.String(%v), expected %v, got %v",
			steamTradeOfferID.String(), expectedValue, gotValue,
		)
	}
}

func TestGetCreateSteamTradeOfferRequest(t *testing.T) {

	baseSteamWebURL := `http://mockymocky.com`
	sessionID := "1234abcde"
	otherSteamID := mangosteam.SteamID(1234567890)
	accessToken := `Er_owt`
	myItems := []*Asset{&Asset{AssetID: 124}, &Asset{AssetID: 125}, &Asset{AssetID: 126}}
	theirItems := []*Asset{&Asset{AssetID: 221}, &Asset{AssetID: 222}, &Asset{AssetID: 223}}
	message := `Mock me over and over!`

	expectedReqURL := `http://mockymocky.com/tradeoffer/new/send`
	expectedReqBody := `captcha=&json_tradeoffer=%7B%22me%22%3A%7B%22assets%22%3A%5B%7B%22assetid%22%3A%22124%22%2C%22appid%22%3A0%2C%22contextid%22%3A%220%22%2C%22amount%22%3A0%7D%2C%7B%22assetid%22%3A%22125%22%2C%22appid%22%3A0%2C%22contextid%22%3A%220%22%2C%22amount%22%3A0%7D%2C%7B%22assetid%22%3A%22126%22%2C%22appid%22%3A0%2C%22contextid%22%3A%220%22%2C%22amount%22%3A0%7D%5D%2C%22currency%22%3A%5B%5D%2C%22ready%22%3Afalse%7D%2C%22newversion%22%3Atrue%2C%22them%22%3A%7B%22assets%22%3A%5B%7B%22assetid%22%3A%22221%22%2C%22appid%22%3A0%2C%22contextid%22%3A%220%22%2C%22amount%22%3A0%7D%2C%7B%22assetid%22%3A%22222%22%2C%22appid%22%3A0%2C%22contextid%22%3A%220%22%2C%22amount%22%3A0%7D%2C%7B%22assetid%22%3A%22223%22%2C%22appid%22%3A0%2C%22contextid%22%3A%220%22%2C%22amount%22%3A0%7D%5D%2C%22currency%22%3A%5B%5D%2C%22ready%22%3Afalse%7D%2C%22version%22%3A2%7D&partner=1234567890&serverid=1&sessionid=1234abcde&trade_offer_create_params=%7B%22trade_offer_access_token%22%3A%22Er_owt%22%7D&tradeoffermessage=Mock+me+over+and+over%21`
	expectedReferer := `http://mockymocky.com/tradeoffer/new/?partner=1234567890`
	expectedContentHeader := `application/x-www-form-urlencoded; charset=UTF-8`

	req, err := getCreateSteamTradeOfferRequest(
		baseSteamWebURL,
		sessionID,
		otherSteamID,
		accessToken,
		myItems, theirItems,
		message,
	)

	if err != nil {
		t.Errorf("getCreateSteamTradeOfferRequest threw an error where it shouldn't: %v", err)
		return
	}

	if req.URL.String() != expectedReqURL {
		t.Errorf("getCreateSteamTradeOfferRequest expected URL '%s', got '%s'",
			expectedReqURL,
			req.URL.String(),
		)
		return
	}

	if req.Referer() != expectedReferer {
		t.Errorf("getCreateSteamTradeOfferRequest expected Referer '%s', got '%s'",
			expectedReferer,
			req.Referer(),
		)
		return
	}

	if req.Header.Get(`Content-Type`) != expectedContentHeader {
		t.Errorf("getCreateSteamTradeOfferRequest expected header '%s', got '%s'",
			expectedContentHeader,
			req.Header.Get(`Content-Type`),
		)
		return
	}

	body, _ := ioutil.ReadAll(req.Body)

	if string(body) != expectedReqBody {
		t.Errorf("getCreateSteamTradeOfferRequest expected Body '%s', got '%s'",
			expectedReqBody,
			body,
		)
		return
	}

}
