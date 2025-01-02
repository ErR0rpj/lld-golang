package entities

import "main/enums"

type ParkingSpot struct {
	Id           int
	Availability enums.Availability
	VehicleType  enums.VehicleType
}

// Creates new parking spot
func NewParkingSpot(id int, vehicleType enums.VehicleType) *ParkingSpot {
	return &ParkingSpot{
		Id:           id,
		Availability: enums.AVAILABLE,
		VehicleType:  vehicleType,
	}
}
