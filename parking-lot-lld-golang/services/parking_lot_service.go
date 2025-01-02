package services

import (
	"bufio"
	"fmt"
	"main/entities"
	"main/enums"
	"os"
)

// Creates the singleton parking lot with its details
func CreateParkingLot() *entities.ParkingLot {
	fmt.Println("Creating parking lot. Enter details.")

	scanner := bufio.NewScanner(os.Stdin)

	var parkingLotName, parkingLotAddress string
	fmt.Println("Enter parking lot name & address: ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	parkingLotName = scanner.Text()
	scanner.Scan()
	parkingLotAddress = scanner.Text()
	parkingLot := entities.NewParkingLot(1, parkingLotName, parkingLotAddress)

	// Creating gates
	var gatesCount int
	fmt.Println("Enter count of gates: ")
	fmt.Scan(&gatesCount)
	for i := range gatesCount {
		var gateType enums.GateType
		fmt.Println("Enter gate type", i+1, "(0 for both, 1 for entry, 2 for exit): ")
		fmt.Scan(&gateType)
		parkingLot.Gates = append(parkingLot.Gates, *entities.NewGate(i+1, gateType))
	}

	// Creating Floors and Spots
	var floorCount int
	fmt.Println("Enter floors count: ")
	fmt.Scan(&floorCount)
	for i := range floorCount {
		var spotCount int
		fmt.Println("Enter spot count for floor", i+1, ": ")
		fmt.Scan(&spotCount)
		parkingLot.ParkingFloors = append(parkingLot.ParkingFloors, *entities.NewParkingFloor(i + 1))
		for j := range spotCount {
			var vehicleType enums.VehicleType
			fmt.Println("Enter vehicle type for spot", j+1, "(0 - Large, 1 - Medium, 2 - Small): ")
			fmt.Scan(&vehicleType)
			parkingLot.ParkingFloors[i].ParkingSpots = append(parkingLot.ParkingFloors[i].ParkingSpots, *entities.NewParkingSpot(j+1, vehicleType))
		}
		parkingLot.ParkingFloors[i].AvailableCapacity = spotCount
		parkingLot.ParkingFloors[i].TotalCapacity = spotCount
		parkingLot.TotalCapacity += spotCount
		parkingLot.AvailableCapacity += spotCount
	}

	fmt.Println("Parking lot created successfully!")
	fmt.Println()
	return parkingLot
}
