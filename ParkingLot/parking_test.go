package parkingLot

import "testing"

// Happy Test - car should be parked.

func TestWhenParkingLotIsAvailableThenCarShouldBeParked(t *testing.T) {
	car := Car{}
	lot := ParkingLot{availableSlots: 1}
	got := lot.Park(car)

	if got != true {
		t.Errorf("Parking failed")
	}
}

func TestWhenParkingLotIsUnavailableThenCarShouldNotBeParked(t *testing.T) {
	car := Car{}
	lot := ParkingLot{availableSlots: 0}
	got := lot.Park(car)

	if got != false {
		t.Errorf("ParkingLot Is Unavailable But Car is parked")
	}
}
