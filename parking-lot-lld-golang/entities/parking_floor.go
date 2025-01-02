package entities

import "main/enums"

type ParkingFloor struct {
	Id                int
	ParkingSpots      []ParkingSpot
	Availability      enums.Availability
	TotalCapacity     int
	AvailableCapacity int
}

func NewParkingFloor(id int) *ParkingFloor {
	return &ParkingFloor{
		Id: id,
	}
}
