package tradeoffer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/vincentserpoul/mangosteam"
	"github.com/vincentserpoul/mangosteam/inventory"
)

func TestGetItemsFromReceipt(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockOKReceipt())
	}))
	defer ts.Close()
	client := http.Client{}
	mangosteam.BaseSteamWebURL = ts.URL

	receiptItems, _ := GetItemsFromReceipt(&client, uint64(123))

	expectedItems := []ReceiptItem{
		ReceiptItem{
			ItemID:         inventory.ItemID(1234),
			OwnerSteamID:   mangosteam.SteamID(12345678999),
			ClassID:        inventory.ClassID(1234567),
			InstanceID:     inventory.InstanceID(0),
			MarketHashName: "MOllusk 12 | 345",
		},
		ReceiptItem{
			ItemID:         inventory.ItemID(1235),
			OwnerSteamID:   mangosteam.SteamID(12345678999),
			ClassID:        inventory.ClassID(1234568),
			InstanceID:     inventory.InstanceID(0),
			MarketHashName: "AKlove 12 | 3",
		},
		ReceiptItem{
			ItemID:         inventory.ItemID(1236),
			OwnerSteamID:   mangosteam.SteamID(12345678999),
			ClassID:        inventory.ClassID(1234569),
			InstanceID:     inventory.InstanceID(0),
			MarketHashName: "TG43 | Elpiero",
		},
		ReceiptItem{
			ItemID:         inventory.ItemID(1237),
			OwnerSteamID:   mangosteam.SteamID(12345678999),
			ClassID:        inventory.ClassID(12345670),
			InstanceID:     inventory.InstanceID(7855747),
			MarketHashName: "Mallo | 441",
		},
	}

	if !reflect.DeepEqual(receiptItems, expectedItems) {
		t.Errorf("GetItemsFromReceipt(), \nexpected: %v\n got: %v\n",
			expectedItems, receiptItems)
	}

}

func TestKOGetItemsFromReceipt(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockKOReceipt())
	}))
	defer ts.Close()
	client := http.Client{}
	mangosteam.BaseSteamWebURL = ts.URL

	_, err := GetItemsFromReceipt(&client, uint64(123))
	if err == nil {
		t.Errorf("GetItemsFromReceipt() should return an error if the json is not correct")
	}
}

func TestStatus500GetItemsFromReceipt(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()
	client := http.Client{}
	mangosteam.BaseSteamWebURL = ts.URL

	_, err := GetItemsFromReceipt(&client, uint64(123))
	if err == nil {
		t.Errorf("GetItemsFromReceipt() should return an error if the server doesnt reply")
	}
}

func TestTimeOutGetItemsFromReceipt(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()
	ts.Config.WriteTimeout = 20 * time.Millisecond
	client := http.Client{}
	mangosteam.BaseSteamWebURL = ts.URL

	_, err := GetItemsFromReceipt(&client, uint64(123))
	if err == nil {
		t.Errorf("GetItemsFromReceipt() should return an error if the server doesnt reply")
	}
}

func TestExtractItemJSONFromBody(t *testing.T) {
	fakeBody := getMockOKReceipt()

	itemsJSON := extractItemJSONFromBody(fakeBody)

	expectedItemsJSON := []string{
		`{"id":"1234","owner":"12345678999","classid":"1234567","instanceid":"0","icon_url":"-9a81dl","icon_url_large":"-9a81dlWMRkL5","icon_drag_url":"","name":"MOllusk 12 | 345","market_hash_name":"MOllusk 12 | 345","market_name":"MOllusk 12 | 345","name_color":"D2D2D2","background_color":"","type":"Mesh","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"html","value":"dfdsfsdf"},{"type":"html","value":" "},{"type":"html","value":"fsefesf"},{"type":"html","value":" "},{"type":"html","value":"Collection Bank","color":"34324324","app_data":{"def_index":"432432","is_itemset_name":1}},{"type":"html","value":" ","app_data":{"def_index":"3432432"}},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"fsefes","name":"PM","category":"Type","category_name":"Type"},{"internal_name":"efsefff","name":"esfs","category":"esfsef","category_name":"seffesf"},{"internal_name":"sefgggse","name":"efse","category":"fesf","category_name":"esfsef"},{"internal_name":"fsegge","name":"efsef","category":"fsefes","category_name":"fesfse"},{"internal_name":"fesfgggsgses","name":"esfesf","category":"34324","color":"b0c3d9","category_name":"3243242"}],"pos":1,"appid":730,"contextid":2}`,
		`{"id":"1235","owner":"12345678999","classid":"1234568","instanceid":"0","icon_url":"-9a81dlVdw8","icon_url_large":"-9a81dEbQ","icon_drag_url":"","name":"AKlove 12 | 3","market_hash_name":"AKlove 12 | 3","market_name":"AKlove 12 | 3","name_color":"32323DD","background_color":"","type":"AKlove 12 | 3","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"wdaw","value":"wdawd"},{"type":"html","value":" "},{"type":"html","value":""},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"fsfsef","name":"fsfsef","category":"gesgs","category_name":"htdthr"}],"pos":2,"appid":730,"contextid":2}`,
		`{"id":"1236","owner":"12345678999","classid":"1234569","instanceid":"0","icon_url":"-9a8-E_","icon_url_large":"-9a81dlWc","icon_drag_url":"","name":"TG43 | Elpiero","market_hash_name":"TG43 | Elpiero","market_name":"TG43 | Elpiero","name_color":"rwerwr","background_color":"","type":"TG43 | Elpiero","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"html","value":"awdawdwadawdawd"},{"type":"html","value":" "},{"type":"html","value":"wadawdawdwa"},{"type":"html","value":" "},{"type":"html","value":"awdawdawda","color":"dawawd","app_data":{"def_index":"awdawd","is_itemset_name":1}},{"type":"html","value":" ","app_data":{"def_index":"dwada"}},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"wdaawd","name":"dwdawdawd","category":"wdawwa","category_name":"dawawd"},{"internal_name":"dwdwd","name":"dwdwdwd","category":"daww","category_name":"dawdaw"}],"pos":3,"appid":730,"contextid":2}`,
		`{"id":"1237","owner":"12345678999","classid":"12345670","instanceid":"7855747","icon_url":"-bgvvnWI1RoN","icon_url_large":"-9qBUup_Omyd","icon_drag_url":"","name":"Mallo | 441","market_hash_name":"Mallo | 441","market_name":"Mallo | 441","name_color":"FRGT","background_color":"","type":"Mallo | 441","tradable":1,"marketable":1,"commodity":0,"market_tradable_restriction":"7","descriptions":[{"type":"html","value":"fsesefse"},{"type":"html","value":" "}],"owner_descriptions":"","tags":[{"internal_name":"gdfgdfg","name":"dfgdfgd","category":"dfgdfgdf","category_name":"gdfgdf"},{"internal_name":"gfgdfgdfgdfg","name":"fgdfgdfgd","category":"dfgdfgdf","category_name":"gdfgdfgdf"}],"pos":4,"appid":730,"contextid":2}`,
	}

	if !reflect.DeepEqual(itemsJSON, expectedItemsJSON) {
		t.Errorf("extractItemJSONFromBody() unexpected response")
	}
}
