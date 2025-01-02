package main

import (
	"fmt"
	"main/entities"
	"main/services"
)

var ParkingLotSingleton *entities.ParkingLot

func main() {
	fmt.Println("----Initializing Parking Lot System----")
	fmt.Println()

	ParkingLotSingleton = services.CreateParkingLot()

}
