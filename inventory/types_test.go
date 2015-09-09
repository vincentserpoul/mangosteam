package inventory

import "testing"

func TestTypesEmpty(t *testing.T) {
	var d Items
	b := []byte("[]")
	err := d.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("Items UnmarshalJSON should return no error if empty description")
	}
}

func TestItemIdString(t *testing.T) {
	var d ItemID
	d = 1
	v := ItemID.String(d)
	if v != "1" {
		t.Errorf("ItemsId String should return a string '1' no %v", v)
	}
}

func TestClassIdString(t *testing.T) {
	var d ClassID
	d = 1
	v := ClassID.String(d)
	if v != "1" {
		t.Errorf("ClassId String should return a string '1' no %v", v)
	}
}

func TestInstanceIDString(t *testing.T) {
	var d InstanceID
	d = 1
	v := InstanceID.String(d)
	if v != "1" {
		t.Errorf("InstanceId String should return a string '1' no %v", v)
	}
}
