package parkinglot

import "errors"

var (
	ErrParkingLotfull = errors.New("Parking lot is full")
)

type ParkingLot struct {
	ID    string
	Slots []Slot
}

func NewParkingLot(id string, slotCount int) *ParkingLot {
	slots := make([]Slot, slotCount)
	return &ParkingLot{ID: id, Slots: slots}
}

func (p *ParkingLot) Park(car Car) (int, error) {
	for i, slot := range p.Slots {
		if slot.Status == Empty {
			p.Slots[i] = Slot{Status: Full, Car: car}
			return i, nil
		}
	}
	return -1, ErrParkingLotfull
}
