package entities

import (
	"time"
)

type Ticket struct {
	EntryTime   time.Time
	Vehicle     VehicleInterface
	ParkingSpot ParkingSpotInterface
}
