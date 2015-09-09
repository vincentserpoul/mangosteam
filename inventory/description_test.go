package inventory

import "testing"

func TestUnmarshalJSON(t *testing.T) {
	var d Descriptions
	b := []byte("[]")
	err := d.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("UnmarshalJSON should return no error if empty description")
	}
}
func TestGetTagNameFromCategory(t *testing.T) {
	cases := []struct {
		tags             Tags
		searchedCategory string
		expectedName     string
	}{
		{
			Tags{
				{
					InternalName: "motus",
					Name:         "craps",
					Category:     "Dedi",
					CategoryName: "Colle",
					Color:        "Pollux",
				},
				{
					InternalName: "modtus",
					Name:         "cradps",
					Category:     "Deddi",
					CategoryName: "Cowwlle",
					Color:        "Polwwlux",
				},
			},
			"poney",
			"",
		},
		{
			Tags{
				{
					InternalName: "motus",
					Name:         "craps",
					Category:     "Dedi",
					CategoryName: "Colle",
					Color:        "Pollux",
				},
				{
					InternalName: "modtus",
					Name:         "cradps",
					Category:     "poney",
					CategoryName: "Cowwlle",
					Color:        "Polwwlux",
				},
			},
			"poney",
			"cradps",
		},
	}

	for _, c := range cases {
		gotName := c.tags.GetTagNameFromCategory(c.searchedCategory)
		if gotName != c.expectedName {
			t.Errorf("GetTagNameFromCategory(%s) == %s, want %s",
				c.searchedCategory, gotName, c.expectedName)
		}
	}

	duplicateTag := Tags{
		{
			InternalName: "ddddwdw",
			Name:         "ddd",
			Category:     "cradps",
			CategoryName: "poney",
			Color:        "dwdwdwd",
		},
		{
			InternalName: "modtus",
			Name:         "craddddps",
			Category:     "poney",
			CategoryName: "Cowwlle",
			Color:        "Polwwlux",
		},
	}

	searchedCategory := "poney"
	gotName := duplicateTag.GetTagNameFromCategory("poney")

	if gotName != "ddd" && gotName != "craddddps" {
		t.Errorf("GetTagNameFromCategory(%s) == %s, want %s or %s",
			searchedCategory, gotName, "ddd", "craddddps")
	}
}
