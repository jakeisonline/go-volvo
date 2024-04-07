package api

import (
	"time"
)

/* HELPER MODELS */

type ValueTimestamp struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type ValueUnitTimestamp struct {
	Value     float32   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}

type SuccessNoStatusModel struct {
	Vin          string `json:"vin"`
	InvokeStatus string `json:"invokeStatus"`
	Message      string `json:"message"`
}

type SuccessWithStatusModel struct {
	SuccessNoStatusModel
	StatusCode int `json:"statusCode"`
}

/* SPECIFIC MODELS */

type ClimatizationStart struct {
	Data struct {
		SuccessNoStatusModel
	} `json:"data"`
}

type ClimatizationStop struct {
	Data struct {
		SuccessNoStatusModel
	} `json:"data"`
}

type CommandsModel []struct {
	Data []struct {
		Command string `json:"command"`
		Href    string `json:"href"`
	} `json:"data"`
}

type CommandAccessibilityModel struct {
	Data struct {
		AvailabilityStatus struct {
			Value             string    `json:"value"`
			UnavailableReason string    `json:"unavailableReason"`
			Timestamp         time.Time `json:"timestamp"`
		} `json:"availabilityStatus"`
	} `json:"data"`
}

type EngineModel struct {
	Data struct {
		EngineCoolantLevelWarning ValueTimestamp `json:"engineCoolantLevelWarning"`
		OilLevelWarning           ValueTimestamp `json:"oilLevelWarning"`
	} `json:"data"`
}

type DiagnosticsModel struct {
	Data struct {
		ServiceWarning       ValueTimestamp     `json:"serviceWarning"`
		ServiceTrigger       ValueTimestamp     `json:"serviceTrigger"`
		EngineHoursToService ValueUnitTimestamp `json:"engineHoursToService"`
		DistanceToService    ValueUnitTimestamp `json:"distanceToService"`
		TimeToService        ValueUnitTimestamp `json:"timeToService"`
	} `json:"data"`
}

type BrakesModel struct {
	Data struct {
		BrakeFluidLevelWarning struct {
			Timestamp  time.Time `json:"timestamp"`
			BrakeFluid string    `json:"brakeFluid"`
		} `json:"brakeFluidLevelWarning"`
	} `json:"data"`
}

type WindowsModel struct {
	Data struct {
		FrontLeftWindow  ValueTimestamp `json:"frontLeftWindow"`
		FrontRightWindow ValueTimestamp `json:"frontRightWindow"`
		RearLeftWindow   ValueTimestamp `json:"rearLeftWindow"`
		RearRightWindow  ValueTimestamp `json:"data"`
	}
}

type DoorsModel struct {
	Data struct {
		CarLocked      ValueTimestamp `json:"carLocked"`
		FrontLeftDoor  ValueTimestamp `json:"frontLeftDoor"`
		FrontRightDoor ValueTimestamp `json:"frontRightDoor"`
		Hood           ValueTimestamp `json:"hood"`
		RearLeftDoor   ValueTimestamp `json:"rearLeftDoor"`
		RearRightDoor  ValueTimestamp `json:"rearRightDoor"`
		TailGate       ValueTimestamp `json:"tailGate"`
		TankLid        ValueTimestamp `json:"tankLid"`
	} `json:"data"`
}

type LockModel struct {
	Data struct {
		SuccessNoStatusModel
	}
}

type LockReducedGuard struct {
	Data struct {
		SuccessNoStatusModel
		ReadyToUnlock      bool `json:"readyToUnlock"`
		ReadyToUnlockUntil int  `json:"readyToUnlockUntil"`
	} `json:"data"`
}

type UnlockModel struct {
	Data struct {
		SuccessNoStatusModel
		ReadyToUnlock      bool `json:"readyToUnlock"`
		ReadyToUnlockUntil int  `json:"readyToUnlockUntil"`
	} `json:"data"`
}

type EngineStatusModel struct {
	Data struct {
		EngineStatus ValueTimestamp `json:"engineStatus"`
	} `json:"data"`
}

type EngineStartModel struct {
	Data struct {
		SuccessNoStatusModel
	} `json:"data"`
}

type EngineStopModel struct {
	Data struct {
		SuccessNoStatusModel
	} `json:"data"`
}

type FuelModel struct {
	FuelAmount         ValueUnitTimestamp `json:"fuelAmount,omitempty"`
	BatteryChargeLevel ValueUnitTimestamp `json:"batteryChargeLevel"`
}

type OdometerModel struct {
	Odometer ValueUnitTimestamp `json:"odometer" mapstructure:",omitempty"`
}

type StatisticsModel struct {
	Data struct {
		AverageFuelConsumption              ValueUnitTimestamp `json:"averageFuelConsumption"`
		AverageEnergyConsumption            ValueUnitTimestamp `json:"averageEnergyConsumption"`
		AverageFuelConsumptionAutomatic     ValueUnitTimestamp `json:"averageFuelConsumptionAutomatic"`
		AverageEnergyConsumptionAutomatic   ValueUnitTimestamp `json:"averageEnergyConsumptionAutomatic"`
		AverageEnergyConsumptionSinceCharge ValueUnitTimestamp `json:"averageEnergyConsumptionSinceCharge"`
		AverageSpeed                        ValueUnitTimestamp `json:"averageSpeed"`
		AverageSpeedAutomatic               ValueUnitTimestamp `json:"averageSpeedAutomatic"`
		TripMeterManual                     ValueUnitTimestamp `json:"tripMeterManual"`
		TripMeterAutomatic                  ValueUnitTimestamp `json:"tripMeterAutomatic"`
		DistanceToEmptyTank                 ValueUnitTimestamp `json:"distanceToEmptyTank"`
		DistanceToEmptyBattery              ValueUnitTimestamp `json:"distanceToEmptyBattery"`
	} `json:"data"`
}

type TyresModel struct {
	Data struct {
		FrontLeft  ValueTimestamp `json:"frontLeft"`
		FrontRight ValueTimestamp `json:"frontRight"`
		RearLeft   ValueTimestamp `json:"rearLeft"`
		RearRight  ValueTimestamp `json:"rearRight"`
	} `json:"data"`
}

type VehiclesModel struct {
	Data []struct {
		Vin string `json:"vin"`
	} `json:"data"`
}

type VehicleModel struct {
	Data struct {
		Vin                string  `json:"vin"`
		ModelYear          int     `json:"modelYear"`
		Gearbox            string  `json:"gearbox"`
		FuelType           string  `json:"fuelType"`
		ExternalColour     string  `json:"externalColour"`
		BatteryCapacityKWH float64 `json:"batteryCapacityKWH"`
		Images             struct {
			ExteriorImageURL string `json:"exteriorImageUrl"`
			InteriorImageURL string `json:"interiorImageUrl"`
		} `json:"images"`
		Descriptions struct {
			Model      string `json:"model"`
			Upholstery string `json:"upholstery"`
			Steering   string `json:"steering"`
		} `json:"descriptions"`
	} `json:"data"`
}

type WarningsModel struct {
	Data struct {
		BrakeLightLeftWarning           ValueTimestamp `json:"brakeLightLeftWarning"`
		BrakeLightCenterWarning         ValueTimestamp `json:"brakeLightCenterWarning"`
		BrakeLightRightWarning          ValueTimestamp `json:"brakeLightRightWarning"`
		FogLightFrontWarning            ValueTimestamp `json:"fogLightFrontWarning"`
		FogLightRearWarning             ValueTimestamp `json:"fogLightRearWarning"`
		PositionLightFrontLeftWarning   ValueTimestamp `json:"positionLightFrontLeftWarning"`
		PositionLightFrontRightWarning  ValueTimestamp `json:"positionLightFrontRightWarning"`
		PositionLightRearLeftWarning    ValueTimestamp `json:"positionLightRearLeftWarning"`
		PositionLightRearRightWarning   ValueTimestamp `json:"positionLightRearRightWarning"`
		HighBeamLeftWarning             ValueTimestamp `json:"highBeamLeftWarning"`
		HighBeamRightWarning            ValueTimestamp `json:"highBeamRightWarning"`
		LowBeamLeftWarning              ValueTimestamp `json:"lowBeamLeftWarning"`
		LowBeamRightWarning             ValueTimestamp `json:"lowBeamRightWarning"`
		DaytimeRunningLightLeftWarning  ValueTimestamp `json:"daytimeRunningLightLeftWarning"`
		DaytimeRunningLightRightWarning ValueTimestamp `json:"daytimeRunningLightRightWarning"`
		TurnIndicationFrontLeftWarning  ValueTimestamp `json:"turnIndicationFrontLeftWarning"`
		TurnIndicationFrontRightWarning ValueTimestamp `json:"turnIndicationFrontRightWarning"`
		TurnIndicationRearLeftWarning   ValueTimestamp `json:"turnIndicationRearLeftWarning"`
		TurnIndicationRearRightWarning  ValueTimestamp `json:"turnIndicationRearRightWarning"`
		RegistrationPlateLightWarning   ValueTimestamp `json:"registrationPlateLightWarning"`
		SideMarkLightsWarning           ValueTimestamp `json:"sideMarkLightsWarning"`
		HazardLightsWarning             ValueTimestamp `json:"hazardLightsWarning"`
		ReverseLightsWarning            ValueTimestamp `json:"reverseLightsWarning"`
	} `json:"data"`
}
