package entities

import (
	"main/enums"
	"sync"
)

// This is a singleton and it is instantiated only once per run. The same object is used across the all goroutines
var (
	parkingLotInstance *ParkingLot
	once               sync.Once
)

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

func NewParkingLot(id int, name string, address string) *ParkingLot {
	once.Do(func() {
		parkingLotInstance = &ParkingLot{
			Id:           id,
			Name:         name,
			Address:      address,
			Availability: enums.AVAILABLE,
		}
	})
	return parkingLotInstance
}
