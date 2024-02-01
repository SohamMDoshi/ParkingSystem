package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptySlot(t *testing.T) {
	emptySlot := Slot{Status: Empty, Car: Car{}}
	assert.Equal(t, Empty, emptySlot.Status)
}

func TestFullSlot(t *testing.T) {
	fullSlot := Slot{Status: Full, Car: Car{Color: "Black", RegistrationNum: "MH12DD000"}}
	assert.Equal(t, Full, fullSlot.Status)
}
