package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jakeisonline/go-volvo/helpers"
)

func unmarshalAndCheck(body []byte, v interface{}) (interface{}, error) {
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func handleResponse(endpointName string, resp *http.Response) (interface{}, error) {
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		switch endpointName {
		case "climatization-start":
			var response ClimatizationStart
			return unmarshalAndCheck(body, &response)
		case "climatization-stop":
			var response ClimatizationStop
			return unmarshalAndCheck(body, &response)
		case "commands":
			var response CommandsModel
			return unmarshalAndCheck(body, &response)
		case "command-accessibility":
			var response CommandAccessibilityModel
			return unmarshalAndCheck(body, &response)
		case "engine":
			var response EngineModel
			return unmarshalAndCheck(body, &response)
		case "diagnostics":
			var response DiagnosticsModel
			return unmarshalAndCheck(body, &response)
		case "brakes":
			var response BrakesModel
			return unmarshalAndCheck(body, &response)
		case "windows":
			var response WindowsModel
			return unmarshalAndCheck(body, &response)
		case "doors":
			var response DoorsModel
			return unmarshalAndCheck(body, &response)
		case "lock":
			var response LockModel
			return unmarshalAndCheck(body, &response)
		case "lock-reduced-guard":
			var response LockReducedGuard
			return unmarshalAndCheck(body, &response)
		case "unlock":
			var response UnlockModel
			return unmarshalAndCheck(body, &response)
		case "engine-status":
			var response EngineStatusModel
			return unmarshalAndCheck(body, &response)
		case "engine-start":
			var response EngineStartModel
			return unmarshalAndCheck(body, &response)
		case "engine-stop":
			var response EngineStopModel
			return unmarshalAndCheck(body, &response)
		case "fuel":
			var response FuelModel
			return unmarshalAndCheck(body, &response)
		case "odometer":
			var response OdometerModel
			return unmarshalAndCheck(body, &response)
		case "statistics":
			var response StatisticsModel
			return unmarshalAndCheck(body, &response)
		case "tyres":
			var response TyresModel
			return unmarshalAndCheck(body, &response)
		case "vehicle":
			var response VehicleModel
			return unmarshalAndCheck(body, &response)
		case "vehicles":
			var response VehiclesModel
			return unmarshalAndCheck(body, &response)
		case "warnings":
			var response WarningsModel
			return unmarshalAndCheck(body, &response)
		default:
			return nil, helpers.Error(
				fmt.Errorf("unable to find relevant Unmarshal type for endpoint named: %v", endpointName),
			)
		}
	} else {
		return nil, helpers.Error(
			fmt.Errorf("status code received: %v\n\n%v", resp.StatusCode, resp),
		)
	}
}
