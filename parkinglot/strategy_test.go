package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFarthestStrategy(t *testing.T) {
	farthestStrategy := FarthestSlotStrategy{}
	slots := []Slot{{Status: Full}, {Status: Empty}, {Status: Full}, {Status: Empty}}
	expectedSlotNumber := 3

	atualSlotNumber, err := farthestStrategy.GetNextSlot(slots)

	assert.NoError(t, err)
	assert.Equal(t, expectedSlotNumber, atualSlotNumber)
}

// func TestFarthestStrategyWhenParkingLotIsFull(t *testing.T) {
// 	farthestSlotStrategy := FarthestSlotStrategy{}
// 	slots := []
// }
