package api

import (
	"fmt"
	"strings"
)

type Endpoint struct {
	Method string
	Url    string
	Scopes []string
}

var prefix string = "https://api.volvocars.com/connected-vehicle/v2/vehicles"

var endpoints = map[string]Endpoint{
	"climatization-start": {
		Url:    prefix + "/{vin}/commands/climatization-start",
		Method: "POST",
		Scopes: []string{"openid", "conve:climatization_start_stop"},
	},
	"climatization-stop": {
		Url:    prefix + "/{vin}/commands/climatization-stop",
		Method: "POST",
		Scopes: []string{"openid", "conve:climatization_start_stop"},
	},
	"commands": {
		Url:    prefix + "/{vin}/commands",
		Method: "GET",
		Scopes: []string{"openid", "conve:commands"},
	},
	"command-accessibility": {
		Url:    prefix + "/{vin}/command-accessibility",
		Method: "GET",
		Scopes: []string{"openid", "conve:command_accessibility"},
	},
	"engine": {
		Url:    prefix + "/{vin}/engine",
		Method: "GET",
		Scopes: []string{"openid", "conve:diagnostics_engine_status"},
	},
	"diagnostics": {
		Url:    prefix + "/{vin}/diagnostics",
		Method: "GET",
		Scopes: []string{"openid", "conve:diagnostics_workshop"},
	},
	"brakes": {
		Url:    prefix + "/{vin}/brakes",
		Method: "GET",
		Scopes: []string{"openid", "conve:brake_status"},
	},
	"windows": {
		Url:    prefix + "/{vin}/windows",
		Method: "GET",
		Scopes: []string{"openid", "conve:windows_status"},
	},
	"doors": {
		Url:    prefix + "/{vin}/doors",
		Method: "GET",
		Scopes: []string{"openid", "conve:doors_status", "conve:lock_status"},
	},
	"lock": {
		Url:    prefix + "/{vin}/commands/lock",
		Method: "POST",
		Scopes: []string{"openid", "conve:lock"},
	},
	"lock-reduced-guard": {
		Url:    prefix + "/{vin}/commands/lock-reduced-guard",
		Method: "POST",
		Scopes: []string{"openid", "conve:lock"},
	},
	"unlock": {
		Url:    prefix + "/{vin}/commands/unlock",
		Method: "POST",
		Scopes: []string{"openid", "conve:unlock"},
	},
	"engine-status": {
		Url:    prefix + "/{vin}/engine-status",
		Method: "GET",
		Scopes: []string{"openid", "conve:engine_status"},
	},
	"engine-start": {
		Url:    prefix + "/{vin}/commands/engine-start",
		Method: "POST",
		Scopes: []string{"openid", "conve:engine_start_stop"},
	},
	"engine-stop": {
		Url:    prefix + "/{vin}/commands/engine-stop",
		Method: "POST",
		Scopes: []string{"openid", "conve:engine_start_stop"},
	},
	"fuel": {
		Url:    prefix + "/{vin}/fuel",
		Method: "GET",
		Scopes: []string{"openid", "conve:fuel_status", "conve:battery_charge_level"},
	},
	"flash": {
		Url:    prefix + "/{vin}/commands/flash",
		Method: "POST",
		Scopes: []string{"openid", "conve:honk_flash"},
	},
	"honk": {
		Url:    prefix + "/{vin}/commands/honk",
		Method: "POST",
		Scopes: []string{"openid", "conve:honk_flash"},
	},
	"honk-flash": {
		Url:    prefix + "/{vin}/commands/honk-flash",
		Method: "POST",
		Scopes: []string{"openid", "conve:honk_flash"},
	},
	"odometer": {
		Url:    prefix + "/{vin}/odometer",
		Method: "GET",
		Scopes: []string{"openid", "conve:odometer_status"},
	},
	"statistics": {
		Url:    prefix + "/{vin}/statistics",
		Method: "GET",
		Scopes: []string{"openid", "conve:trip_statistics"},
	},
	"tyres": {
		Url:    prefix + "/{vin}/tyres",
		Method: "GET",
		Scopes: []string{"openid", "conve:tyre_status"},
	},
	"vehicles": {
		Url:    prefix,
		Method: "GET",
		Scopes: []string{"openid", "conve:vehicle_relation"},
	},
	"vehicle": {
		Url:    prefix + "/{vin}",
		Method: "GET",
		Scopes: []string{"openid", "conve:vehicle_relation"},
	},
	"warnings": {
		Url:    prefix + "/{vin}/warnings",
		Method: "GET",
		Scopes: []string{"openid", "conve:warnings"},
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
