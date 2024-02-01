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
