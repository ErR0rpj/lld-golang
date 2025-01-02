package enums

type VehicleType int

const (
	LARGE_VEHICLE  VehicleType = iota // Truck, trailers, etc
	MEDIUM_VEHICLE                    // SUVs, Luxury Cars, Tempo, etc
	SMALL_VEHICLE                     // Mini SUVs,  Hatchback, Sedan etc
)
