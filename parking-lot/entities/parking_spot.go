package entities

import "fmt"

type ParkingSpotInterface interface {
	ParkVehicle(v VehicleInterface)
	RemoveVehicle(v VehicleInterface)

	GetPrice() int
	GetAvailability() bool
	GetId() int
}

// Concrete Parking spot class
type ParkingSpot struct {
	Id      int
	IsEmpty bool
	Vehicle VehicleInterface
	Price   int
}

func (spot *ParkingSpot) ParkVehicle(v VehicleInterface) {
	spot.Vehicle = v
	spot.IsEmpty = false
}

func (spot *ParkingSpot) RemoveVehicle(v VehicleInterface) {
	spot.Vehicle = nil
	spot.IsEmpty = true
}

func (spot *ParkingSpot) GetPrice() int {
	return spot.Price
}

func (spot *ParkingSpot) GetId() int {
	return spot.Id
}

func (spot *ParkingSpot) GetAvailability() bool {
	return spot.IsEmpty
}

// 2 wheel
type TwoWheelerParkingSpot struct {
	ParkingSpot
}

// 4 wheel
type FourWheelerParkingSpot struct {
	ParkingSpot
}

func CreateParkingSpot(id int, spotType string) (ParkingSpotInterface, error) {
	if spotType == "2" {
		return &TwoWheelerParkingSpot{
			ParkingSpot: ParkingSpot{
				Id:      id,
				IsEmpty: true,
				Price:   10,
			},
		}, nil
	} else if spotType == "4" {
		return &TwoWheelerParkingSpot{
			ParkingSpot: ParkingSpot{
				Id:      id,
				IsEmpty: true,
				Price:   20,
			},
		}, nil
	} else {
		return nil, fmt.Errorf("spot type not defined")
	}
}
