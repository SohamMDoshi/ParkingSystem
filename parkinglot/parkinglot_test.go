package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParkingLot(t *testing.T) {
	id := "PL001"
	slotCount := 5
	parkingLot := NewParkingLot(id, slotCount)

	assert.Equal(t, id, parkingLot.ID)
	assert.Len(t, parkingLot.Slots, slotCount)
}

func TestParkCar(t *testing.T) {
	parkinglot := NewParkingLot("PL001", 1)

	car := Car{Color: "Red", RegistrationNum: "MH07SS8888"}
	ticket, err := parkinglot.Park(car)
	assert.NoError(t, err)
	assert.Equal(t, Full, parkinglot.Slots[ticket.SlotID].Status)
}

func TestParkCarWhenParkingLotIsFull_ExpectError(t *testing.T) {
	parkinglot := NewParkingLot("PL001", 1)

	car := Car{Color: "Red", RegistrationNum: "MH07SS8888"}
	ticket, err := parkinglot.Park(car)
	assert.NoError(t, err)
	assert.Equal(t, Full, parkinglot.Slots[ticket.SlotID].Status)

	car2 := Car{Color: "Black", RegistrationNum: "JK09GG1012"}
	_, err1 := parkinglot.Park(car2)
	assert.EqualError(t, err1, ErrParkingLotfull.Error())
}

func TestUnParkCarWithVaildTicket(t *testing.T) {
	parkinglot := NewParkingLot("PL001", 1)

	car := Car{Color: "Red", RegistrationNum: "MH07SS8888"}
	ticket, err := parkinglot.Park(car)
	assert.NoError(t, err)
	assert.Equal(t, Full, parkinglot.Slots[ticket.SlotID].Status)

	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH07SS8888"}
	assert.True(t, expectedTicket.Equals(ticket))

	unparkedCar, err := parkinglot.Unpark(ticket)
	assert.True(t, unparkedCar.Equals(car))
}

func TestUnParkCarWithInvaildTicket(t *testing.T) {
	parkinglot := NewParkingLot("PL001", 1)

	car := Car{Color: "Red", RegistrationNum: "MH07SS8888"}
	ticket, err := parkinglot.Park(car)
	assert.NoError(t, err)
	assert.Equal(t, Full, parkinglot.Slots[ticket.SlotID].Status)
	expectedTicket := Ticket{ParkingLotID: "PL001", SlotID: 0, CarRegistration: "MH07SS8888"}
	assert.True(t, expectedTicket.Equals(ticket))

	inVaildTicket := Ticket{ParkingLotID: "PL003", SlotID: 3, CarRegistration: "JK07AA3333"}
	_, err1 := parkinglot.Unpark(inVaildTicket)
	assert.EqualError(t, err1, ErrInValidSlotNumber.Error())
}
