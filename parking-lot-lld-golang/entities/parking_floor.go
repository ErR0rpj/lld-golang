package entities

import "main/enums"

type ParkingFloor struct {
	Id                string
	ParkingSpots      []ParkingSpot
	Availability      enums.Availability
	TotalCapacity     int
	AvailableCapacity int
}
