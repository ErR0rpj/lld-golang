package entities

import "main/enums"

type ParkingSpot struct {
	Id           int
	Availability enums.Availability
	VehicleType  enums.VehicleType
}
