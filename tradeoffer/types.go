package tradeoffer

import (
	"strconv"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/inventory"
)

// State represents the state of the tradeoffer, see constants
type State uint

const (
	// ETradeOfferStateCreated /!\ non steam status, used to know the TO has been created
	ETradeOfferStateCreated State = iota
	// ETradeOfferStateInvalid Invalid
	ETradeOfferStateInvalid
	// ETradeOfferStateActive This trade offer has been sent, neither party has acted on it yet.
	ETradeOfferStateActive
	// ETradeOfferStateAccepted The trade offer was accepted by the recipient and items were exchanged.
	ETradeOfferStateAccepted
	// ETradeOfferStateCountered The recipient made a counter offer
	ETradeOfferStateCountered
	// ETradeOfferStateExpired The trade offer was not accepted before the expiration date
	ETradeOfferStateExpired
	// ETradeOfferStateCanceled The sender cancelled the offer
	ETradeOfferStateCanceled
	// ETradeOfferStateDeclined The recipient declined the offer
	ETradeOfferStateDeclined
	// ETradeOfferStateInvalidItems Some of the items in the offer are no longer available
	// (indicated by the missing flag in the output)
	ETradeOfferStateInvalidItems
	// ETradeOfferStateEmailPending The offer hasn't been sent yet and is awaiting email confirmation
	ETradeOfferStateEmailPending
	// ETradeOfferStateEmailCanceled The receiver cancelled the offer via email
	ETradeOfferStateEmailCanceled
)

// CEconAsset represents an asset in steam web api
type CEconAsset struct {
	AppID      mangosteam.AppID     `json:",string"`
	ContextID  mangosteam.ContextID `json:",string"`
	AssetID    AssetID              `json:",string"`
	CurrencyID uint64               `json:",string"`
	ClassID    inventory.ClassID    `json:",string"`
	InstanceID inventory.InstanceID `json:",string"`
	Amount     uint64               `json:",string"`
	Missing    bool
}

// CEconTradeOffer represent the to from the steam API
type CEconTradeOffer struct {
	TradeOfferID   SteamTradeOfferID  `json:",string"`
	OtherAccountID mangosteam.SteamID `json:"accountid_other"`
	Message        string
	ExpirationTime uint32        `json:"expiraton_time"`
	State          State         `json:"trade_offer_state"`
	ToGive         []*CEconAsset `json:"items_to_give"`
	ToReceive      []*CEconAsset `json:"items_to_receive"`
	IsOurs         bool          `json:"is_our_offer"`
	TimeCreated    uint32        `json:"time_created"`
	TimeUpdated    uint32        `json:"time_updated"`
}

// CEconTradeOffers contains a list of tradeoffers, sent and received
type CEconTradeOffers struct {
	Sent         []*CEconTradeOffer     `json:"trade_offers_sent"`
	Received     []*CEconTradeOffer     `json:"trade_offers_received"`
	Descriptions inventory.Descriptions `json:"descriptions"`
}

// SteamTradeOfferID is the identifier of the tradeoffer within steam network
type SteamTradeOfferID uint64

// String will turn a steamID into a string
func (steamTradeOfferID SteamTradeOfferID) String() string {
	return strconv.FormatUint(uint64(steamTradeOfferID), 10)
}
