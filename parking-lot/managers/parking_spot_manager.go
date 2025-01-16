package manager

import (
	"fmt"
	"main/entities"
	"main/enums"
	"main/strategy"
)

var TwoWheelerParkingSpotManagerSingleton ParkingSpotManagerInterface
var FourWheelerParkingSpotManagerSingleton ParkingSpotManagerInterface

type ParkingSpotManagerInterface interface {
	AddParkingSpot(int)
	RemoveParkingSpot()
	ParkVehicle(entities.VehicleInterface) (entities.ParkingSpotInterface, error)
	RemoveVehicle(entities.ParkingSpotInterface, entities.VehicleInterface)

	GetCount() int
}

type ParkingSpotManager struct {
	ParkingSpotList     []entities.ParkingSpotInterface
	ParkingSpotStrategy strategy.ParkingStrategyInterface
}

func (spotManager *ParkingSpotManager) GetCount() int {
	return len(spotManager.ParkingSpotList)
}

func (spotManager *ParkingSpotManager) RemoveParkingSpot() {
	index := len(spotManager.ParkingSpotList) - 1
	spotManager.ParkingSpotList = append(spotManager.ParkingSpotList[:index], spotManager.ParkingSpotList[index+1:]...)
}

func (spotManager *ParkingSpotManager) ParkVehicle(v entities.VehicleInterface) (entities.ParkingSpotInterface, error) {
	parkingSpot, err := spotManager.ParkingSpotStrategy.FindParkingSpot(spotManager.ParkingSpotList)
	if err != nil {
		return nil, err
	}

	parkingSpot.ParkVehicle(v)
	return parkingSpot, err
}

func (spotManager *ParkingSpotManager) RemoveVehicle(parkingSpot entities.ParkingSpotInterface, v entities.VehicleInterface) {
	parkingSpot.RemoveVehicle(v)
}

// 2 wheeler manager
type TwoWheelerParkingSpotManager struct {
	ParkingSpotManager
}

func (spotManager *TwoWheelerParkingSpotManager) AddParkingSpot(id int) {
	spot, _ := entities.CreateParkingSpot(id, "2")
	spotManager.ParkingSpotList = append(spotManager.ParkingSpotList, spot)
}

// 4 wheeler manager
type FourWheelerParkingSpotManager struct {
	ParkingSpotManager
}

func (spotManager *FourWheelerParkingSpotManager) AddParkingSpot(id int) {
	spot, _ := entities.CreateParkingSpot(id, "4")
	spotManager.ParkingSpotList = append(spotManager.ParkingSpotList, spot)
}

// Factory to create manager
func GetParkingSpotManager(vechicleType enums.VehicleType) (ParkingSpotManagerInterface, error) {
	if TwoWheelerParkingSpotManagerSingleton == nil {
		TwoWheelerParkingSpotManagerSingleton = &TwoWheelerParkingSpotManager{
			ParkingSpotManager: ParkingSpotManager{
				ParkingSpotStrategy: &strategy.DefaultStrategy{},
			},
		}
	}
	if FourWheelerParkingSpotManagerSingleton == nil {
		FourWheelerParkingSpotManagerSingleton = &FourWheelerParkingSpotManager{
			ParkingSpotManager: ParkingSpotManager{
				ParkingSpotStrategy: &strategy.NearestToEntranceStrategy{},
			},
		}
	}

	switch vechicleType {
	case enums.TwoWheelerVehicle:
		return TwoWheelerParkingSpotManagerSingleton, nil
	case enums.FourWheelerVehicle:
		return FourWheelerParkingSpotManagerSingleton, nil
	default:
		return nil, fmt.Errorf("vehicle type undefined")
	}

}
