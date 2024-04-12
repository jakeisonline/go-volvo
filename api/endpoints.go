package api

import (
	"fmt"
	"strings"
)

type Endpoint struct {
	Method string
	Url    string
}

var prefix string = "https://api.volvocars.com/connected-vehicle/v2/vehicles"

var endpoints = map[string]Endpoint{
	"climatization-start": {
		Url:    prefix + "/{vin}/commands/climatization-start",
		Method: "POST",
	},
	"climatization-stop": {
		Url:    prefix + "/{vin}/commands/climatization-stop",
		Method: "POST",
	},
	"commands": {
		Url:    prefix + "/{vin}/commands",
		Method: "GET",
	},
	"command-accessibility": {
		Url:    prefix + "/{vin}/command-accessibility",
		Method: "GET",
	},
	"engine": {
		Url:    prefix + "/{vin}/engine",
		Method: "GET",
	},
	"diagnostics": {
		Url:    prefix + "/{vin}/diagnostics",
		Method: "GET",
	},
	"brakes": {
		Url:    prefix + "/{vin}/brakes",
		Method: "GET",
	},
	"windows": {
		Url:    prefix + "/{vin}/windows",
		Method: "GET",
	},
	"doors": {
		Url:    prefix + "/{vin}/doors",
		Method: "GET",
	},
	"lock": {
		Url:    prefix + "/{vin}/commands/lock",
		Method: "POST",
	},
	"lock-reduced-guard": {
		Url:    prefix + "/{vin}/commands/lock-reduced-guard",
		Method: "POST",
	},
	"unlock": {
		Url:    prefix + "/{vin}/commands/unlock",
		Method: "POST",
	},
	"engine-status": {
		Url:    prefix + "/{vin}/engine-status",
		Method: "GET",
	},
	"engine-start": {
		Url:    prefix + "/{vin}/commands/engine-start",
		Method: "POST",
	},
	"engine-stop": {
		Url:    prefix + "/{vin}/commands/engine-stop",
		Method: "POST",
	},
	"fuel": {
		Url:    prefix + "/{vin}/fuel",
		Method: "GET",
	},
	"flash": {
		Url:    prefix + "/{vin}/commands/flash",
		Method: "POST",
	},
	"honk": {
		Url:    prefix + "/{vin}/commands/honk",
		Method: "POST",
	},
	"honk-flash": {
		Url:    prefix + "/{vin}/commands/honk-flash",
		Method: "POST",
	},
	"odometer": {
		Url:    prefix + "/{vin}/odometer",
		Method: "GET",
	},
	"statistics": {
		Url:    prefix + "/{vin}/statistics",
		Method: "GET",
	},
	"tyres": {
		Url:    prefix + "/{vin}/tyres",
		Method: "GET",
	},
	"vehicles": {
		Url:    prefix,
		Method: "GET",
	},
	"vehicle": {
		Url:    prefix + "/{vin}",
		Method: "GET",
	},
	"warnings": {
		Url:    prefix + "/{vin}/warnings",
		Method: "GET",
	},
}

func GetEndpoint(endpointName string, vin string) Endpoint {
	endpointMatch := endpoints[endpointName]

	if len(endpointMatch.Url) == 0 {
		throwPanic(fmt.Sprintf("Unable to match endpoint named \"%v\"", endpointName))
	}

	if strings.Contains(endpointMatch.Url, "{vin}") {
		endpointMatch.Url = replaceVin(endpointMatch.Url, vin)
	}

	return endpointMatch
}

func replaceVin(endpointUrl string, vin string) string {
	if len(vin) == 0 {
		throwPanic("Valid endpoint, but no VIN supplied")
	}

	return strings.Replace(endpointUrl, "{vin}", vin, -1)
}
