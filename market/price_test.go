package market

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/vincentserpoul/mangosteam"
)

func TestGetPrice(t *testing.T) {

	cases := []struct {
		Mock          func() string
		ExpectedErr   error
		ExpectedPrice float64
	}{
		{
			Mock: getMockMarket, ExpectedErr: nil, ExpectedPrice: float64(1.04),
		},
		{
			Mock: getBadFloatMockMarket, ExpectedErr: errors.New("error"), ExpectedPrice: 0,
		},
		{
			Mock: getFailedMockMarket, ExpectedErr: errors.New("error"), ExpectedPrice: 0,
		},
		{
			Mock: getMalformedMockMarket, ExpectedErr: errors.New("error"), ExpectedPrice: 0,
		},
	}

	for _, c := range cases {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, c.Mock())
		}))
		defer ts.Close()

		gotPrice, err := GetPrice(ts.URL, mangosteam.AppID(730), "test")
		if c.ExpectedErr != nil {
			if err == nil {
				t.Errorf("GetPrice(%v) expected an error, didnt get any", c)
			}
			continue
		}
		if gotPrice != c.ExpectedPrice {
			t.Errorf("GetPrice(%v) expected %f, got %f", c, c.ExpectedPrice, gotPrice)
			return
		}
	}
}

func TestGetPriceOverview(t *testing.T) {

	cases := []struct {
		Mock                  func() string
		ExpectedErr           error
		ExpectedPriceOverview PriceOverview
	}{
		{
			Mock:        getMockMarket,
			ExpectedErr: nil,
			ExpectedPriceOverview: PriceOverview{
				Success:     true,
				LowestPrice: "$1.04",
				Volume:      166,
				MedianPrice: "$1.04",
			},
		},
		{
			Mock: getMalformedMockMarket, ExpectedErr: errors.New("error"), ExpectedPriceOverview: PriceOverview{},
		},
	}

loopcases:
	for _, c := range cases {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, c.Mock())
		}))
		defer ts.Close()

		gotPriceOverview, err := getPriceOverview(ts.URL, mangosteam.AppID(730), "test")
		if c.ExpectedErr != nil {
			if err == nil {
				t.Errorf("GetPriceOverview(%v) expected an error, didnt get any", c)
			}
			continue loopcases
		}
		if *gotPriceOverview != c.ExpectedPriceOverview {
			t.Errorf("GetPriceOverview(%v) expected %v, got %v", c, c.ExpectedPriceOverview, gotPriceOverview)
			return
		}
	}
}

func TestTimeoutGetPriceOverview(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()
	ts.Config.WriteTimeout = 20 * time.Millisecond

	_, err := getPriceOverview(ts.URL, mangosteam.AppID(730), "test")
	if err == nil {
		t.Errorf("GetPriceOverview() expected an error, didnt get any")
	}
}

func TestGetPriceURL(t *testing.T) {

	baseSteamWebURL := `http://www.testurl.com`
	appID := mangosteam.AppID(730)
	marketHashName := `$fdgggr | pollux`
	expectedURL := `http://www.testurl.com/market/priceoverview?currency=1&appid=730&format=json&language=en&market_hash_name=%24fdgggr+%7C+pollux`

	gotURL := getPriceURL(baseSteamWebURL, appID, marketHashName)

	gotURLP, _ := url.Parse(gotURL)
	expectedURLP, _ := url.Parse(expectedURL)

	if !reflect.DeepEqual(gotURLP.Query(), expectedURLP.Query()) ||
		expectedURLP.RequestURI() != expectedURLP.RequestURI() {
		t.Errorf("getPriceURL(%s, %d, %s): expected\n%s\ngot\n%s", baseSteamWebURL, appID, marketHashName, expectedURL, gotURL)
	}
}
