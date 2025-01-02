package entities

import "main/enums"

type Vehicle struct {
	VehicleNumber string
	VehicleType   enums.VehicleType
}

func NewVehicle(vehicleNumber string, vehicleType enums.VehicleType) *Vehicle {
	return &Vehicle{
		VehicleNumber: vehicleNumber,
		VehicleType:   vehicleType,
	}
}
