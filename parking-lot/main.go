package main

import (
	"main/entities"
	"main/enums"
	manager "main/managers"
)

func main() {
	vehicle1, _ := entities.CreateVehicle(enums.TwoWheelerVehicle, 1)
	vehicle2, _ := entities.CreateVehicle(enums.TwoWheelerVehicle, 2)
	vehicle3, _ := entities.CreateVehicle(enums.TwoWheelerVehicle, 3)
	vehicle4, _ := entities.CreateVehicle(enums.FourWheelerVehicle, 4)

	entrance := manager.EntranceManger{
		EntranceNumber: 1,
	}

	manager.AddParkingSpots()

	entrance.BookParkingSpot(vehicle1)
	entrance.BookParkingSpot(vehicle2)
	entrance.BookParkingSpot(vehicle3)
	entrance.BookParkingSpot(vehicle4)

}
