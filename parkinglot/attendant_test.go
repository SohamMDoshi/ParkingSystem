package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssigningParkingLotToAttendant(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLot := NewParkingLot("PL001", 2)

	attendant.Assign(parkingLot)

	assert.Equal(t, parkingLot.ID, attendant.AssignedLots["PL001"].ID)
}

func TestAssigningInvaildParkingLotToAttendant(t *testing.T) {
	attendat := NewAttendant("ABC")
	parkinglot := NewParkingLot("PL001", 2)
	parkinglotTwo := NewParkingLot("PL002", 2)

	attendat.Assign(parkinglot)

	assert.NotEqual(t, parkinglotTwo.ID, attendat.AssignedLots["PL001"].ID)
}

func TestParkingCar(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLot := NewParkingLot("PL001", 2)

	attendant.Assign(parkingLot)
	assert.Equal(t, parkingLot.ID, attendant.AssignedLots["PL001"].ID)

	car, err := NewCar("Black", "MH12AA3003")
	assert.NoError(t, err)
	assert.True(t, car.Equals(Car{"Black", "MH12AA3003"}))

	ticket, err1 := attendant.Park(car)
	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH12AA3003"}
	assert.NoError(t, err1)
	assert.True(t, expectedTicket.Equals(ticket))
}

func TestParkingCarWhenParkingLotIsFull(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLot := NewParkingLot("PL001", 1)

	attendant.Assign(parkingLot)
	assert.Equal(t, parkingLot.ID, attendant.AssignedLots["PL001"].ID)

	carOne, err := NewCar("Black", "MH12AA3003")
	assert.NoError(t, err)
	assert.True(t, carOne.Equals(Car{"Black", "MH12AA3003"}))

	carTwo, err := NewCar("White", "JK77DD1111")
	assert.NoError(t, err)
	assert.True(t, carTwo.Equals(Car{"White", "JK77DD1111"}))

	ticket, err1 := attendant.Park(carOne)
	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH12AA3003"}
	assert.NoError(t, err1)
	assert.True(t, expectedTicket.Equals(ticket))

	_, err2 := attendant.Park(carTwo)
	assert.Error(t, err2, ErrParkingLotfull)
}

func TestAssigningMultipleParkingLotToOneAttendant(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLotOne := NewParkingLot("PL001", 1)
	parkingLotTwo := NewParkingLot("PL002", 1)

	attendant.Assign(parkingLotOne)
	assert.True(t, parkingLotOne.Equals(attendant.AssignedLots["PL001"]))

	attendant.Assign(parkingLotTwo)
	assert.True(t, parkingLotTwo.Equals(attendant.AssignedLots["PL002"]))
}

func TestParkingCarWhenPakringLotOneIsFullAndOtherParkingLotIsAvailable(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLotOne := NewParkingLot("PL001", 1)
	parkingLotTwo := NewParkingLot("PL002", 1)

	attendant.Assign(parkingLotOne)
	assert.True(t, parkingLotOne.Equals(attendant.AssignedLots["PL001"]))

	attendant.Assign(parkingLotTwo)
	assert.True(t, parkingLotTwo.Equals(attendant.AssignedLots["PL002"]))

	carOne, err := NewCar("Black", "MH12AA3003")
	assert.NoError(t, err)
	assert.True(t, carOne.Equals(Car{"Black", "MH12AA3003"}))

	carTwo, err := NewCar("White", "JK77DD1111")
	assert.NoError(t, err)
	assert.True(t, carTwo.Equals(Car{"White", "JK77DD1111"}))

	ticket, err1 := attendant.Park(carOne)
	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH12AA3003"}
	assert.NoError(t, err1)
	assert.True(t, expectedTicket.Equals(ticket))

	ticketOfCarTwo, err2 := attendant.Park(carTwo)
	expectedTicketOfCarTwo := Ticket{ParkingLotID: "PL002", SlotID: 0, CarRegistration: "JK77DD1111"}
	assert.NoError(t, err2)
	assert.True(t, expectedTicketOfCarTwo.Equals(ticketOfCarTwo))
}

func TestParkingCarWhenAllParkingLotsAreFull(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLotOne := NewParkingLot("PL001", 1)
	parkingLotTwo := NewParkingLot("PL002", 1)

	attendant.Assign(parkingLotOne)
	assert.True(t, parkingLotOne.Equals(attendant.AssignedLots["PL001"]))

	attendant.Assign(parkingLotTwo)
	assert.True(t, parkingLotTwo.Equals(attendant.AssignedLots["PL002"]))

	carOne, err := NewCar("Black", "MH12AA3003")
	assert.NoError(t, err)
	assert.True(t, carOne.Equals(Car{"Black", "MH12AA3003"}))

	carTwo, err := NewCar("White", "JK77DD1111")
	assert.NoError(t, err)
	assert.True(t, carTwo.Equals(Car{"White", "JK77DD1111"}))

	carThree, err := NewCar("Red", "KA88JH9999")
	assert.NoError(t, err)
	assert.True(t, carThree.Equals(Car{"Red", "KA88JH9999"}))

	ticket, err1 := attendant.Park(carOne)
	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH12AA3003"}
	assert.NoError(t, err1)
	assert.True(t, expectedTicket.Equals(ticket))

	ticketOfCarTwo, err2 := attendant.Park(carTwo)
	expectedTicketOfCarTwo := Ticket{ParkingLotID: "PL002", SlotID: 0, CarRegistration: "JK77DD1111"}
	assert.NoError(t, err2)
	assert.True(t, expectedTicketOfCarTwo.Equals(ticketOfCarTwo))

	_, err3 := attendant.Park(carTwo)
	assert.Error(t, err3, ErrParkingLotfull)
}

func TestUnparkingCar(t *testing.T) {
	attendant := NewAttendant("John")
	parkingLot := NewParkingLot("PL001", 2)

	attendant.Assign(parkingLot)
	assert.Equal(t, parkingLot.ID, attendant.AssignedLots["PL001"].ID)

	car, err := NewCar("Black", "MH12AA3003")
	assert.NoError(t, err)
	assert.True(t, car.Equals(Car{"Black", "MH12AA3003"}))

	ticket, err1 := attendant.Park(car)
	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH12AA3003"}
	assert.NoError(t, err1)
	assert.True(t, expectedTicket.Equals(ticket))

	unparkedCar, err2 := attendant.Unpark(ticket)
	assert.NoError(t, err2)
	assert.True(t, unparkedCar.Equals(car))
}
