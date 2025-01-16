package entities

import (
	"fmt"
	"main/enums"
)

type VehicleInterface interface {
	GetVehicleType() enums.VehicleType
}

type Vehicle struct {
	VehicleNumber int
	VehicleType   enums.VehicleType
}

func (vehicle *Vehicle) GetVehicleType() enums.VehicleType {
	return vehicle.VehicleType
}

func CreateVehicle(vehicleType enums.VehicleType, vehicleNumber int) (VehicleInterface, error) {
	switch vehicleType {
	case enums.FourWheelerVehicle:
		return &Vehicle{
				VehicleNumber: vehicleNumber,
				VehicleType:   vehicleType,
			},
			nil
	case enums.TwoWheelerVehicle:
		return &Vehicle{
			VehicleNumber: vehicleNumber,
			VehicleType:   vehicleType,
		}, nil
	default:
		return nil, fmt.Errorf("vehicle type undefined")
	}
}
