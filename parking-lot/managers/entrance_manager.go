package manager

import (
	"fmt"
	"main/entities"
	"main/enums"
	"time"
)

type EntranceMangerInterface interface {
	BookParkingSpot(entities.Vehicle) (*entities.Ticket, error)
}

type EntranceManger struct {
	EntranceNumber int
}

func (entranceManage *EntranceManger) BookParkingSpot(vehicle entities.VehicleInterface) (*entities.Ticket, error) {
	parkingSpotManager, _ := GetParkingSpotManager(vehicle.GetVehicleType())
	parkingSpot, err := parkingSpotManager.ParkVehicle(vehicle)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ticket := &entities.Ticket{
		EntryTime:   time.Now(),
		Vehicle:     vehicle,
		ParkingSpot: parkingSpot,
	}

	fmt.Println("parking spot booked at", parkingSpot.GetId())
	return ticket, nil
}

func AddParkingSpots() {
	TwoparkingSpotManager, _ := GetParkingSpotManager(enums.TwoWheelerVehicle)
	TwoparkingSpotManager.AddParkingSpot(1)
	TwoparkingSpotManager.AddParkingSpot(2)

	FourWheelerParkingSpotManager, _ := GetParkingSpotManager(enums.FourWheelerVehicle)
	FourWheelerParkingSpotManager.AddParkingSpot(1)
	FourWheelerParkingSpotManager.AddParkingSpot(2)
}
