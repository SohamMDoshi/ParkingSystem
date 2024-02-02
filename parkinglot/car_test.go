package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingNewCarWithValidRegistratrion(t *testing.T) {
	car, err := NewCar("Black", "MH12AA3003")
	assert.NoError(t, err)
	assert.True(t, car.Equals(Car{"Black", "MH12AA3003"}))
}

func TestCreatingNewCarWithInvalidRegistratrion(t *testing.T) {
	_, err := NewCar("Black", "AA3003")
	assert.Error(t, err)
}
