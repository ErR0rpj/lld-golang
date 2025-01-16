package strategy

import (
	"fmt"
	"main/entities"
)

type ParkingStrategyInterface interface {
	FindParkingSpot(parkingSpotList []entities.ParkingSpotInterface) (entities.ParkingSpotInterface, error)
}

type NearestToEntranceStrategy struct {
}

func (s *NearestToEntranceStrategy) FindParkingSpot(parkingSpotList []entities.ParkingSpotInterface) (entities.ParkingSpotInterface, error) {
	fmt.Println("using nearest strategy")
	for i := range parkingSpotList {
		if parkingSpotList[i].GetAvailability() {
			return parkingSpotList[i], nil
		}
	}
	return nil, fmt.Errorf("all parking spots are full")
}

type DefaultStrategy struct {
}

func (s *DefaultStrategy) FindParkingSpot(parkingSpotList []entities.ParkingSpotInterface) (entities.ParkingSpotInterface, error) {
	fmt.Println("using default strategy")
	for i := range parkingSpotList {
		if parkingSpotList[i].GetAvailability() {
			return parkingSpotList[i], nil
		}
	}
	return nil, fmt.Errorf("all parking spots are full")
}
