package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptySlot(t *testing.T) {
	emptySlot := Slot{Status: Empty, Car: Car{}}
	assert.Equal(t, Empty, emptySlot.Status)
}
