package tradeoffer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vincentserpoul/mangosteam"
)

func TestCancelSteamTradeOffer(t *testing.T) {

	sessionID := "1234abcde"
	creatorSteamID := mangosteam.SteamID(1234567890)
	steamTradeOfferID := SteamTradeOfferID(1098765432)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, cancelMockSteamTradeOffer())
	}))
	defer ts.Close()
	client := http.Client{}

	_, err := CancelSteamTradeOffer(
		ts.URL,
		&client,
		sessionID,
		creatorSteamID,
		steamTradeOfferID,
	)
	if err != nil {
		t.Errorf("CancelSteamTradeOffer threw an error where it shouldn't: %v", err)
		return
	}

}

func TestTimeOutCancelSteamTradeOffer(t *testing.T) {

	sessionID := "1234abcde"
	creatorSteamID := mangosteam.SteamID(1234567890)
	steamTradeOfferID := SteamTradeOfferID(1098765432)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()
	ts.Config.WriteTimeout = 20 * time.Millisecond
	client := http.Client{}

	_, err := CancelSteamTradeOffer(
		ts.URL,
		&client,
		sessionID,
		creatorSteamID,
		steamTradeOfferID,
	)
	if err == nil {
		t.Errorf("CreateSteamTradeOffer validate where it shouldn't: %v", err)
		return
	}

}

func TestNotFoundSteamTradeOffer(t *testing.T) {

	sessionID := "1234abcde"
	creatorSteamID := mangosteam.SteamID(1234567890)
	steamTradeOfferID := SteamTradeOfferID(1098765432)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()
	client := http.Client{}

	_, err := CancelSteamTradeOffer(
		ts.URL,
		&client,
		sessionID,
		creatorSteamID,
		steamTradeOfferID,
	)
	if err == nil {
		t.Errorf("CreateSteamTradeOffer validate where it shouldn't: %v", err)
		return
	}

}

func TestNoBodySteamTradeOffer(t *testing.T) {

	sessionID := "1234abcde"
	creatorSteamID := mangosteam.SteamID(1234567890)
	steamTradeOfferID := SteamTradeOfferID(1098765432)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()
	client := http.Client{}

	_, err := CancelSteamTradeOffer(
		ts.URL,
		&client,
		sessionID,
		creatorSteamID,
		steamTradeOfferID,
	)
	if err == nil {
		t.Errorf("CreateSteamTradeOffer validate where it shouldn't: %v", err)
		return
	}

}

func TestGetCancelSteamTradeOfferRequest(t *testing.T) {

	baseSteamWebURL := `http://mockymocky.com`
	sessionID := "1234abcde"
	creatorSteamID := mangosteam.SteamID(1234567890)
	steamTradeOfferID := SteamTradeOfferID(1098765432)

	expectedReqURL := `http://mockymocky.com/tradeoffer/1098765432/decline`
	expectedContentHeader := `application/x-www-form-urlencoded; charset=UTF-8`
	expectedReferer := `http://mockymocky.com/profiles/1234567890/tradeoffers/`
	expectedReqBody := `sessionid=1234abcde`

	req, err := getCancelSteamTradeOfferRequest(
		baseSteamWebURL,
		sessionID,
		creatorSteamID,
		steamTradeOfferID,
	)

	if err != nil {
		t.Errorf("getCancelSteamTradeOfferRequest threw an error where it shouldn't: %v", err)
		return
	}

	if req.URL.String() != expectedReqURL {
		t.Errorf("getCancelSteamTradeOfferRequest expected URL '%s', got '%s'",
			expectedReqURL,
			req.URL.String(),
		)
		return
	}

	if req.Referer() != expectedReferer {
		t.Errorf("getCancelSteamTradeOfferRequest expected Referer '%s', got '%s'",
			expectedReferer,
			req.Referer(),
		)
		return
	}

	if req.Header.Get(`Content-Type`) != expectedContentHeader {
		t.Errorf("getCancelSteamTradeOfferRequest expected header '%s', got '%s'",
			expectedContentHeader,
			req.Header.Get(`Content-Type`),
		)
		return
	}

	body, _ := ioutil.ReadAll(req.Body)

	if string(body) != expectedReqBody {
		t.Errorf("getCancelSteamTradeOfferRequest expected Body '%s', got '%s'",
			expectedReqBody,
			body,
		)
		return
	}

}
