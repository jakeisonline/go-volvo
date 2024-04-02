package api

import (
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	endpoint := "vehicles"

	expected := "https://api.volvocars.com/connected-vehicle/v2/vehicles"

	if result := GetEndpoint(endpoint, "").Url; result != expected {
		t.Errorf("GetEndpoint(%s, \"\"), didn't return %s", endpoint, expected)
	}
}

func TestReplaceVin(t *testing.T) {
	endpoint := "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands"
	vin := "1GNEK13ZX3R298984"

	expected := "https://api.volvocars.com/connected-vehicle/v2/vehicles/1GNEK13ZX3R298984/commands"

	if got := replaceVin(endpoint, vin); got != expected {
		t.Errorf("replaceVin(%s, %s), didn't return %s", endpoint, vin, expected)
	}
}
