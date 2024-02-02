package parkinglot

import (
	"github.com/google/go-cmp/cmp"
)

type Ticket struct {
	ParkingLotID    string
	SlotID          int
	CarRegistration string
}

func (t Ticket) Equals(other Ticket) bool {
	return cmp.Equal(t, other)
}
