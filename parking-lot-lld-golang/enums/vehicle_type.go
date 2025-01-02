package enums

type VehicleType int

const (
	LARGE_VEHICLE  VehicleType = iota // Truck, trailers, etc
	MEDIUM_VEHICLE                    // SUVs, Luxury Cars, Tempo, etc
	SMALL_VEHICLE                     // Mini SUVs,  Hatchback, Sedan etc
)

func (a VehicleType) String() string {
	switch a {
	case LARGE_VEHICLE:
		return "Large"
	case MEDIUM_VEHICLE:
		return "Medium"
	case SMALL_VEHICLE:
		return "Small"
	}
	return ""
}
