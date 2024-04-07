package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jakeisonline/go-volvo/helpers"
)

func Call(endpointName string, vin string) (interface{}, error) {
	resp, err := request(endpointName, vin)
	return resp, err
}

func request(endpointName string, vin string) (interface{}, error) {
	endpointMap := GetEndpoint(endpointName, vin)

	req, err := http.NewRequest(endpointMap.Method, endpointMap.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+GetAccessToken())
	req.Header.Add("VCC-API-Key", GetApiKey())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		switch endpointName {
		case "climitization-start":
			var response ClimitizationStart
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "climitization-stop":
			var response ClimitizationStop
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "commands":
			var response CommandsModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "command-accessibility":
			var response CommandAccessibilityModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "engine-model":
			var response EngineModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "diagnostics":
			var response DiagnosticsModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "brakes":
			var response BrakesModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "doors":
			var response DoorsModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "lock":
			var response LockModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "lock-reduced-guard":
			var response LockReducedGuard
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "unlock":
			var response UnlockModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "engine-status":
			var response EngineStatusModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "engine-start":
			var response EngineStartModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "engine-stop":
			var response EngineStopModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "fuel":
			var response FuelModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "odometer":
			var response OdometerModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "statistics":
			var response StatisticsModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "tyres":
			var response TyresModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "vehicle":
			var response VehicleModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "vehicles":
			var response VehiclesModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
		case "warnings":
			var response WarningsModel
			err = json.Unmarshal(body, &response)
			if err != nil {
				return nil, err
			}
			return response, nil
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
