package entities

import "time"

type Ticket struct {
	Id          int
	ParkingSpot *ParkingSpot
	Vehicle     *Vehicle
	Gate        *Gate
	EntryTime   time.Time
}

func NewTicket(id int, parkingSpot *ParkingSpot, vehicle *Vehicle, gate *Gate, entryTime time.Time) *Ticket {
	return &Ticket{
		Id:          id,
		ParkingSpot: parkingSpot,
		Vehicle:     vehicle,
		Gate:        gate,
		EntryTime:   entryTime,
	}
}
