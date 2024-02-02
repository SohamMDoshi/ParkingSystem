package parkinglot

import (
	"errors"

	"github.com/google/go-cmp/cmp"
)

var (
	ErrParkingLotfull    = errors.New("Parking lot is full")
	ErrInValidSlotNumber = errors.New("slot number is Invalid")
)

type ParkingLot struct {
	ID    string
	Slots []Slot
}

func NewParkingLot(id string, slotCount int) ParkingLot {
	slots := make([]Slot, slotCount)
	return ParkingLot{ID: id, Slots: slots}
}

func (p *ParkingLot) IsParkingLotFull() bool {
	for _, slot := range p.Slots {
		if slot.Status == Empty {
			return false
		}
	}
	return true
}

func (p *ParkingLot) IsCarPresent(registration string) bool {
	for _, slot := range p.Slots {
		if slot.Status == Full && slot.Car.RegistrationNum == registration {
			return true
		}
	}
	return false
}

func (p *ParkingLot) Park(car Car) (Ticket, error) {
	if p.IsCarPresent(car.RegistrationNum) {
		return Ticket{}, errors.New("Car with same registration number is already present")
	}
	if p.IsParkingLotFull() {
		return Ticket{}, ErrParkingLotfull
	}
	for i, slot := range p.Slots {
		if slot.Status == Empty {
			p.Slots[i] = Slot{Status: Full, Car: car}
			return Ticket{
				ParkingLotID:    p.ID,
				SlotID:          i,
				CarRegistration: car.RegistrationNum,
			}, nil
		}
	}
	return Ticket{}, ErrParkingLotfull
}

func (p *ParkingLot) Unpark(ticket Ticket) (Car, error) {
	if ticket.SlotID < 0 || ticket.SlotID >= len(p.Slots) {
		return Car{}, ErrInValidSlotNumber
	}
	slot := p.Slots[ticket.SlotID]
	if slot.Status == Empty {
		return Car{}, errors.New("slot is already empty")
	}

	car := slot.Car
	slot.Status = Empty
	slot.Car = Car{}
	return car, nil
}

func (p ParkingLot) Equals(other ParkingLot) bool {
	return cmp.Equal(p, other)
}
