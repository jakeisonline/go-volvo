package api

import (
	"time"
)

/* HELPER MODELS */

type valueTimestamp struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type valueUnitTimestamp struct {
	Value     float32   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}

type successNoStatusModel struct {
	Vin          string `json:"vin"`
	InvokeStatus string `json:"invokeStatus"`
	Message      string `json:"message"`
}

type successWithStatusModel struct {
	successNoStatusModel
	StatusCode int `json:"statusCode"`
}

/* SPECIFIC MODELS */

type ClimatizationStart struct {
	Data struct {
		successNoStatusModel
	} `json:"data"`
}

type ClimatizationStop struct {
	Data struct {
		successNoStatusModel
	} `json:"data"`
}

type CommandsModel struct {
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
		EngineCoolantLevelWarning valueTimestamp `json:"engineCoolantLevelWarning"`
		OilLevelWarning           valueTimestamp `json:"oilLevelWarning"`
	} `json:"data"`
}

type DiagnosticsModel struct {
	Data struct {
		ServiceWarning          valueTimestamp     `json:"serviceWarning"`
		EngineHoursToService    valueUnitTimestamp `json:"engineHoursToService"`
		DistanceToService       valueUnitTimestamp `json:"distanceToService"`
		WasherFluidLevelWarning valueTimestamp     `json:"washerFluidLevelWarning"`
		TimeToService           valueUnitTimestamp `json:"timeToService"`
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
		FrontLeftWindow  valueTimestamp `json:"frontLeftWindow"`
		FrontRightWindow valueTimestamp `json:"frontRightWindow"`
		RearLeftWindow   valueTimestamp `json:"rearLeftWindow"`
		RearRightWindow  valueTimestamp `json:"rearRightWindow"`
		Sunroof          valueTimestamp `json:"sunroof"`
	}
}

type DoorsModel struct {
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

type LockModel struct {
	Data struct {
		successNoStatusModel
	}
}

type LockReducedGuard struct {
	Data struct {
		successNoStatusModel
		ReadyToUnlock      bool `json:"readyToUnlock"`
		ReadyToUnlockUntil int  `json:"readyToUnlockUntil"`
	} `json:"data"`
}

type UnlockModel struct {
	Data struct {
		successNoStatusModel
		ReadyToUnlock      bool `json:"readyToUnlock"`
		ReadyToUnlockUntil int  `json:"readyToUnlockUntil"`
	} `json:"data"`
}

type EngineStatusModel struct {
	Data struct {
		EngineStatus valueTimestamp `json:"engineStatus"`
	} `json:"data"`
}

type EngineStartModel struct {
	Data struct {
		successNoStatusModel
	} `json:"data"`
}

type EngineStopModel struct {
	Data struct {
		successNoStatusModel
	} `json:"data"`
}

type FuelModel struct {
	Data struct {
		FuelAmount         valueUnitTimestamp `json:"fuelAmount,omitempty"`
		BatteryChargeLevel valueUnitTimestamp `json:"batteryChargeLevel"`
	}
}

type OdometerModel struct {
	Data struct {
		Odometer valueUnitTimestamp `json:"odometer" mapstructure:",omitempty"`
	}
}

type StatisticsModel struct {
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

type TyresModel struct {
	Data struct {
		FrontLeft  valueTimestamp `json:"frontLeft"`
		FrontRight valueTimestamp `json:"frontRight"`
		RearLeft   valueTimestamp `json:"rearLeft"`
		RearRight  valueTimestamp `json:"rearRight"`
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
			InternalImageURL string `json:"internalImageUrl"`
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
