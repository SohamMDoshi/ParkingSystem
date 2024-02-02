package parkinglot

type ParkingStratefy interface {
	GetNextSlot(slots []Slot) (int, error)
}

type FarthestSlotStrategy struct{}

func (s FarthestSlotStrategy) GetNextSlot(slots []Slot) (int, error) {
	for i := len(slots) - 1; i >= 0; i-- {
		if slots[i].Status == Empty {
			return i, nil
		}
	}
	return -1, ErrParkingLotfull
}
