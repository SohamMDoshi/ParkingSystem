package parkinglot

type ParkingLot struct {
	ID    string
	Slots []Slot
}

func NewParkingLot(id string, slotCount int) *ParkingLot {
	slots := make([]Slot, slotCount)
	return &ParkingLot{ID: id, Slots: slots}
}
