package tradeoffer

import (
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestAssetDefaults(t *testing.T) {
	asset := Asset{AssetID: 123}
	appID := mangosteam.AppID(730)
	contextID := mangosteam.ContextID(2)

	expectedDefaultAsset :=
		Asset{
			AssetID:   123,
			AppID:     appID,
			Amount:    1,
			ContextID: contextID,
		}

	asset.Defaults(appID)

	if asset != expectedDefaultAsset {
		t.Errorf(
			"Asset.Default(%v), expected %v, got %v",
			appID.String(), expectedDefaultAsset, asset,
		)
	}
}

func TestAssetString(t *testing.T) {
	var assetID AssetID
	assetID = 123
	expectedValue := "123"
	gotValue := assetID.String()
	if expectedValue != gotValue {
		t.Errorf(
			"Asset.String(%v), expected %v, got %v",
			assetID.String(), expectedValue, gotValue,
		)

	}
}
