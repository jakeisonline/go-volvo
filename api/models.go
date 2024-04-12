package api

import (
	"time"
)

/* HELPER ResponseS */

type valueTimestamp struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type valueUnitTimestamp struct {
	Value     float32   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}

type successNoStatus struct {
	Vin          string `json:"vin"`
	InvokeStatus string `json:"invokeStatus"`
	Message      string `json:"message"`
}

type successWithStatus struct {
	successNoStatus
	StatusCode int `json:"statusCode"`
}

/* SPECIFIC ResponseS */

type ClimatizationStart struct {
	Data struct {
		successNoStatus
	} `json:"data"`
}

type ClimatizationStop struct {
	Data struct {
		successNoStatus
	} `json:"data"`
}

type CommandsResponse struct {
	Data []struct {
		Command string `json:"command"`
		Href    string `json:"href"`
	} `json:"data"`
}

type CommandAccessibilityResponse struct {
	Data struct {
		AvailabilityStatus struct {
			Value             string    `json:"value"`
			UnavailableReason string    `json:"unavailableReason"`
			Timestamp         time.Time `json:"timestamp"`
		} `json:"availabilityStatus"`
	} `json:"data"`
}

type EngineResponse struct {
	Data struct {
		EngineCoolantLevelWarning valueTimestamp `json:"engineCoolantLevelWarning"`
		OilLevelWarning           valueTimestamp `json:"oilLevelWarning"`
	} `json:"data"`
}

type DiagnosticsResponse struct {
	Data struct {
		ServiceWarning          valueTimestamp     `json:"serviceWarning"`
		EngineHoursToService    valueUnitTimestamp `json:"engineHoursToService"`
		DistanceToService       valueUnitTimestamp `json:"distanceToService"`
		WasherFluidLevelWarning valueTimestamp     `json:"washerFluidLevelWarning"`
		TimeToService           valueUnitTimestamp `json:"timeToService"`
	} `json:"data"`
}

type BrakesResponse struct {
	Data struct {
		BrakeFluidLevelWarning struct {
			Timestamp  time.Time `json:"timestamp"`
			BrakeFluid string    `json:"brakeFluid"`
		} `json:"brakeFluidLevelWarning"`
	} `json:"data"`
}

type WindowsResponse struct {
	Data struct {
		FrontLeftWindow  valueTimestamp `json:"frontLeftWindow"`
		FrontRightWindow valueTimestamp `json:"frontRightWindow"`
		RearLeftWindow   valueTimestamp `json:"rearLeftWindow"`
		RearRightWindow  valueTimestamp `json:"rearRightWindow"`
		Sunroof          valueTimestamp `json:"sunroof"`
	}
}

type DoorsResponse struct {
	Data struct {
		CentralLock    valueTimestamp `json:"centralLock"`
		FrontLeftDoor  valueTimestamp `json:"frontLeftDoor"`
		FrontRightDoor valueTimestamp `json:"frontRightDoor"`
		Hood           valueTimestamp `json:"hood"`
		RearLeftDoor   valueTimestamp `json:"rearLeftDoor"`
		RearRightDoor  valueTimestamp `json:"rearRightDoor"`
		TailGate       valueTimestamp `json:"tailGate"`
		TankLid        valueTimestamp `json:"tankLid"`
	} `json:"data"`
}

type LockResponse struct {
	Data struct {
		successNoStatus
	}
}

type LockReducedGuard struct {
	Data struct {
		successNoStatus
		ReadyToUnlock      bool `json:"readyToUnlock"`
		ReadyToUnlockUntil int  `json:"readyToUnlockUntil"`
	} `json:"data"`
}

type UnlockResponse struct {
	Data struct {
		successWithStatus
		ReadyToUnlock      bool `json:"readyToUnlock"`
		ReadyToUnlockUntil int  `json:"readyToUnlockUntil"`
	} `json:"data"`
}

type EngineStatusResponse struct {
	Data struct {
		EngineStatus valueTimestamp `json:"engineStatus"`
	} `json:"data"`
}

type EngineStartResponse struct {
	Data struct {
		successNoStatus
	} `json:"data"`
}

type EngineStopResponse struct {
	Data struct {
		successNoStatus
	} `json:"data"`
}

type FuelResponse struct {
	Data struct {
		FuelAmount         valueUnitTimestamp `json:"fuelAmount,omitempty"`
		BatteryChargeLevel valueUnitTimestamp `json:"batteryChargeLevel"`
	}
}

type OdometerResponse struct {
	Data struct {
		Odometer valueUnitTimestamp `json:"odometer" mapstructure:",omitempty"`
	}
}

type StatisticsResponse struct {
	Data struct {
		AverageFuelConsumption          valueUnitTimestamp `json:"averageFuelConsumption"`
		AverageEnergyConsumption        valueUnitTimestamp `json:"averageEnergyConsumption"`
		AverageFuelConsumptionAutomatic valueUnitTimestamp `json:"averageFuelConsumptionAutomatic"`
		AverageSpeed                    valueUnitTimestamp `json:"averageSpeed"`
		AverageSpeedAutomatic           valueUnitTimestamp `json:"averageSpeedAutomatic"`
		TripMeterManual                 valueUnitTimestamp `json:"tripMeterManual"`
		TripMeterAutomatic              valueUnitTimestamp `json:"tripMeterAutomatic"`
		DistanceToEmptyTank             valueUnitTimestamp `json:"distanceToEmptyTank"`
		DistanceToEmptyBattery          valueUnitTimestamp `json:"distanceToEmptyBattery"`
	} `json:"data"`
}

type TyresResponse struct {
	Data struct {
		FrontLeft  valueTimestamp `json:"frontLeft"`
		FrontRight valueTimestamp `json:"frontRight"`
		RearLeft   valueTimestamp `json:"rearLeft"`
		RearRight  valueTimestamp `json:"rearRight"`
	} `json:"data"`
}

type VehiclesResponse struct {
	Data []struct {
		Vin string `json:"vin"`
	} `json:"data"`
}

type VehicleResponse struct {
	Data struct {
		Vin                string  `json:"vin"`
		ResponseYear       int     `json:"ResponseYear"`
		Gearbox            string  `json:"gearbox"`
		FuelType           string  `json:"fuelType"`
		ExternalColour     string  `json:"externalColour"`
		BatteryCapacityKWH float64 `json:"batteryCapacityKWH"`
		Images             struct {
			ExteriorImageURL string `json:"exteriorImageUrl"`
			InternalImageURL string `json:"internalImageUrl"`
		} `json:"images"`
		Descriptions struct {
			Response   string `json:"Response"`
			Upholstery string `json:"upholstery"`
			Steering   string `json:"steering"`
		} `json:"descriptions"`
	} `json:"data"`
}

type WarningsResponse struct {
	Data struct {
		BrakeLightLeftWarning           valueTimestamp `json:"brakeLightLeftWarning"`
		BrakeLightCenterWarning         valueTimestamp `json:"brakeLightCenterWarning"`
		BrakeLightRightWarning          valueTimestamp `json:"brakeLightRightWarning"`
		FogLightFrontWarning            valueTimestamp `json:"fogLightFrontWarning"`
		FogLightRearWarning             valueTimestamp `json:"fogLightRearWarning"`
		PositionLightFrontLeftWarning   valueTimestamp `json:"positionLightFrontLeftWarning"`
		PositionLightFrontRightWarning  valueTimestamp `json:"positionLightFrontRightWarning"`
		PositionLightRearLeftWarning    valueTimestamp `json:"positionLightRearLeftWarning"`
		PositionLightRearRightWarning   valueTimestamp `json:"positionLightRearRightWarning"`
		HighBeamLeftWarning             valueTimestamp `json:"highBeamLeftWarning"`
		HighBeamRightWarning            valueTimestamp `json:"highBeamRightWarning"`
		LowBeamLeftWarning              valueTimestamp `json:"lowBeamLeftWarning"`
		LowBeamRightWarning             valueTimestamp `json:"lowBeamRightWarning"`
		DaytimeRunningLightLeftWarning  valueTimestamp `json:"daytimeRunningLightLeftWarning"`
		DaytimeRunningLightRightWarning valueTimestamp `json:"daytimeRunningLightRightWarning"`
		TurnIndicationFrontLeftWarning  valueTimestamp `json:"turnIndicationFrontLeftWarning"`
		TurnIndicationFrontRightWarning valueTimestamp `json:"turnIndicationFrontRightWarning"`
		TurnIndicationRearLeftWarning   valueTimestamp `json:"turnIndicationRearLeftWarning"`
		TurnIndicationRearRightWarning  valueTimestamp `json:"turnIndicationRearRightWarning"`
		RegistrationPlateLightWarning   valueTimestamp `json:"registrationPlateLightWarning"`
		SideMarkLightsWarning           valueTimestamp `json:"sideMarkLightsWarning"`
		HazardLightsWarning             valueTimestamp `json:"hazardLightsWarning"`
		ReverseLightsWarning            valueTimestamp `json:"reverseLightsWarning"`
	} `json:"data"`
}
