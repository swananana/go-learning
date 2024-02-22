package parkingLot

// As a driver, I should be able to park my car in the parking lot so that I can catch my flight.

type Car struct {
}

type ParkingLot struct {
	availableSlots int
}

func (l *ParkingLot) Park(car Car) bool {
	if l.availableSlots > 0 {
		l.availableSlots--
		return true
	}
	return false
}
