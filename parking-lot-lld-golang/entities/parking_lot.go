package entities

import "main/enums"

type ParkingLot struct {
	Id                int
	Name              string
	Address           string
	ParkingFloors     []ParkingFloor
	Gates             []Gate
	Availability      enums.Availability
	TotalCapacity     int
	AvailableCapacity int
}
