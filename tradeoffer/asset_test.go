package tradeoffer

import (
	"testing"

	"github.com/vincentserpoul/mangosteam"
)

func TestDefaults(t *testing.T) {
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
