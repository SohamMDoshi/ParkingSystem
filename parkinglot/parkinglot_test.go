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
	slotNum, err := parkinglot.Park(car)
	assert.NoError(t, err)
	assert.Equal(t, Full, parkinglot.Slots[slotNum].Status)
}

func TestParkCarWhenParkingLotIsFull(t *testing.T) {
	parkinglot := NewParkingLot("PL001", 1)

	car1 := Car{Color: "Red", RegistrationNum: "MH07SS8888"}
	slotNum, err := parkinglot.Park(car1)
	assert.NoError(t, err)
	assert.Equal(t, Full, parkinglot.Slots[slotNum].Status)

	car2 := Car{Color: "Black", RegistrationNum: "JK09GG1012"}
	_, err1 := parkinglot.Park(car2)
	assert.EqualError(t, err1, ErrParkingLotfull.Error())

}
