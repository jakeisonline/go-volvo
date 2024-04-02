package api

import (
	"fmt"
	"strings"
)

type Endpoint struct {
	Method string
	Url    string
}

var endpoints = map[string]Endpoint{
	"climatization-start": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/climatization-start",
		Method: "POST",
	},
	"climatization-stop": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/climatization-stop",
		Method: "POST",
	},
	"commands": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands",
		Method: "GET",
	},
	"command-accessibility": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/command-accessibility",
		Method: "GET",
	},
	"diagnostics": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/engine",
		Method: "GET",
	},
	"windows": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/windows",
		Method: "GET",
	},
	"doors": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/doors",
		Method: "GET",
	},
	"lock": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/lock",
		Method: "POST",
	},
	"lock-reduced-guard": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/lock-reduced-guard",
		Method: "POST",
	},
	"unlock": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/unlock",
		Method: "POST",
	},
	"engine-status": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/engine-status",
		Method: "GET",
	},
	"engine-start": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/engine-start",
		Method: "POST",
	},
	"engine-stop": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/engine-stop",
		Method: "POST",
	},
	"fuel": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/fuel",
		Method: "GET",
	},
	"flash": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/flash",
		Method: "POST",
	},
	"honk": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/honk",
		Method: "POST",
	},
	"honk-flash": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/commands/honk-flash",
		Method: "POST",
	},
	"odometer": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/odometer",
		Method: "GET",
	},
	"statistics": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/statistics",
		Method: "GET",
	},
	"tyres": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/tyres",
		Method: "GET",
	},
	"vehicles": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles",
		Method: "GET",
	},
	"vehicle": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}",
		Method: "GET",
	},
	"warnings": {
		Url:    "https://api.volvocars.com/connected-vehicle/v2/vehicles/{vin}/warnings",
		Method: "GET",
	},
}

func GetEndpoint(endpointName string, vin string) Endpoint {
	endpointMatch := endpoints[endpointName]

	if len(endpointMatch.Url) == 0 {
		errorString := fmt.Sprintf("[go-volvo]: Unable to match endpoint named \"%v\"", endpointName)
		panic(errorString)
	}

	if strings.Contains(endpointMatch.Url, "{vin}") {
		endpointMatch.Url = replaceVin(endpointMatch.Url, vin)
	}

	return endpointMatch
}

func replaceVin(endpointUrl string, vin string) string {
	if len(vin) == 0 {
		errorString := "[go-volvo]: Valid endpoint, but no VIN supplied"
		panic(errorString)
	}

	return strings.Replace(endpointUrl, "{vin}", vin, -1)
}
