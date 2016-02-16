package tradeoffer

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/auth"
)

// Confirmation is a tradeoffer confirmation
type Confirmation struct {
	ConfirmationID          string
	ConfirmationKey         string
	ConfirmationDescription string
}

// FetchConfirmations returns a list of tradeoffers to confirm
func FetchConfirmations(
	client *http.Client,
	steamID mangosteam.SteamID,
	baseSteamAPIURL string,
	baseSteamWebURL string,
	sga *auth.SteamGuardAccount,
) ([]*Confirmation, error) {

	req, err := getFetchConfirmationsRequest(
		steamID,
		baseSteamAPIURL,
		baseSteamWebURL,
		sga, "conf",
	)
	if err != nil {
		return nil, fmt.Errorf("tradeoffer FetchConfirmations: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("tradeoffer FetchConfirmations: %v", err)
	}

	bodyStr := string(body)
	// Nothing to confirm
	if strings.Contains(bodyStr, "<div>Nothing to confirm</div>") {
		return nil, nil
	}
	if strings.Contains(bodyStr, "<div>There was a problem loading the confirmations page. Please try your request again later.</div>") {
		return nil,
			fmt.Errorf("tradeoffer FetchConfirmations: problem loading the confirmations page. Please try your request again later")
	}

	// Try to parse response
	confIDs := regexp.MustCompile("data-confid=\"(\\d+)\"").
		FindAllStringSubmatch(bodyStr, -1)
	confKeys := regexp.MustCompile("data-key=\"(\\d+)\"").
		FindAllStringSubmatch(bodyStr, -1)
	confDescs := regexp.MustCompile("<div>((Confirm|Trade with|Sell -) .+)</div>").
		FindAllStringSubmatch(bodyStr, -1)

	if confIDs == nil || confKeys == nil || confDescs == nil {
		return nil, fmt.Errorf("tradeoffer FetchConfirmations: failed to parse response")
	}

	if len(confIDs) != len(confKeys) || len(confIDs) != len(confDescs) {
		return nil, fmt.Errorf("tradeoffer FetchConfirmations: unexpected response format: number of ids, keys and descriptions are not the same")
	}

	// Create confirmations slice
	var confirmations []*Confirmation
	for index := range confIDs {
		cn := &Confirmation{
			ConfirmationID:          confIDs[index][1],
			ConfirmationKey:         confKeys[index][1],
			ConfirmationDescription: confDescs[index][1],
		}
		confirmations = append(confirmations, cn)
	}

	return confirmations, nil
}

func getFetchConfirmationsRequest(
	steamID mangosteam.SteamID,
	baseSteamAPIURL string,
	baseSteamWebURL string,
	sga *auth.SteamGuardAccount,
	tag string,
) (*http.Request, error) {
	queryParams, err := generateConfirmationQueryParams(
		baseSteamAPIURL,
		steamID,
		sga,
		tag,
	)
	if err != nil {
		return nil, err
	}

	return http.NewRequest("GET", baseSteamWebURL+"/mobileconf/conf?"+queryParams.Encode(), nil)
}

// AcceptConfirmation will accept confirmation
func AcceptConfirmation(
	client *http.Client,
	steamID mangosteam.SteamID,
	baseSteamAPIURL string,
	baseSteamWebURL string,
	sga *auth.SteamGuardAccount,
	cn *Confirmation,
) error {
	return sendConfirmation(
		client,
		steamID,
		baseSteamAPIURL,
		baseSteamWebURL,
		sga,
		cn,
		"allow",
	)
}

// DenyConfirmation will deny the tradeoffer confirmation
func DenyConfirmation(
	client *http.Client,
	steamID mangosteam.SteamID,
	baseSteamAPIURL string,
	baseSteamWebURL string,
	sga *auth.SteamGuardAccount,
	cn *Confirmation,
) error {
	return sendConfirmation(
		client,
		steamID,
		baseSteamAPIURL,
		baseSteamWebURL,
		sga,
		cn,
		"cancel",
	)
}

func sendConfirmation(
	client *http.Client,
	steamID mangosteam.SteamID,
	baseSteamAPIURL string,
	baseSteamWebURL string,
	sga *auth.SteamGuardAccount,
	cn *Confirmation,
	op string,
) error {

	req, err := getSendConfirmationsRequest(
		baseSteamAPIURL,
		baseSteamWebURL,
		steamID,
		sga,
		cn,
		op,
	)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r := sendConfirmationResponse{}
	if err = json.Unmarshal(body, &r); err != nil {
		return err
	}
	if !r.Success {
		return errors.New("steam returned success false")
	}

	return nil
}

func getSendConfirmationsRequest(
	baseSteamAPIURL string,
	baseSteamWebURL string,
	steamID mangosteam.SteamID,
	sga *auth.SteamGuardAccount,
	cn *Confirmation,
	tag string,
) (*http.Request, error) {
	queryParams, err := generateConfirmationQueryParams(
		baseSteamAPIURL,
		steamID,
		sga,
		tag,
	)
	if err != nil {
		return nil, err
	}

	queryParams.Set("op", tag)
	queryParams.Set("cid", cn.ConfirmationID)
	queryParams.Set("ck", cn.ConfirmationKey)

	return http.NewRequest("GET", baseSteamWebURL+"/mobileconf/ajaxop?"+queryParams.Encode(), nil)
}

func generateConfirmationQueryParams(
	baseSteamAPIURL string,
	steamID mangosteam.SteamID,
	sga *auth.SteamGuardAccount,
	tag string,
) (url.Values, error) {
	if sga.DeviceID == "" {
		return nil, errors.New("Device ID is empty")
	}
	t := auth.GetSteamTime(baseSteamAPIURL)
	queryParams := url.Values{}
	queryParams.Set("p", sga.DeviceID)
	queryParams.Set("a", steamID.String())
	queryParams.Set("k", generateConfirmationHashForTime(sga, t, tag))
	queryParams.Set("t", strconv.FormatInt(t, 10))
	queryParams.Set("m", "android")
	queryParams.Set("tag", tag)
	return queryParams, nil
}

func generateConfirmationHashForTime(
	sga *auth.SteamGuardAccount,
	t int64,
	tag string,
) string {
	identitySecretBytes, err := base64.StdEncoding.DecodeString(sga.IdentitySecret)
	if err != nil {
		// TODO: maybe we shall panic or return error up the chain
		return ""
	}

	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(t))
	tagBytes := []byte(tag)
	if len(tagBytes) > 32 {
		// maximum tag length is 32 bytes
		tagBytes = tagBytes[:32]
	}
	data = append(data, tagBytes...)

	// Generate hmac
	hmacGenerator := hmac.New(sha1.New, identitySecretBytes)
	hmacGenerator.Write(data)
	mac := hmacGenerator.Sum(nil)

	return base64.StdEncoding.EncodeToString(mac)
}

type sendConfirmationResponse struct {
	Success bool `json:"success"`
}
