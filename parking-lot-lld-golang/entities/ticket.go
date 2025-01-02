package entities

import "time"

type Ticket struct {
	Id          int
	ParkingSpot *ParkingSpot
	Vehicle     *Vehicle
	Gate        *Gate
	EntryTime   time.Time
}
