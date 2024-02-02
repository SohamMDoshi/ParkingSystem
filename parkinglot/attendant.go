package parkinglot

import "errors"

type Attendant struct {
	Name         string
	AssignedLots map[string]ParkingLot
}

func NewAttendant(name string) Attendant {
	return Attendant{
		Name:         name,
		AssignedLots: make(map[string]ParkingLot),
	}
}

func (a *Attendant) Assign(parkinglot ParkingLot) {
	a.AssignedLots[parkinglot.ID] = parkinglot
}

func (a *Attendant) Park(car Car) (Ticket, error) {
	for _, lot := range a.AssignedLots {
		ticket, err := lot.Park(car)
		if err == nil {
			return ticket, nil
		}
	}
	return Ticket{}, ErrParkingLotfull
}

func (a *Attendant) Unpark(ticket Ticket) (Car, error) {
	lot, found := a.AssignedLots[ticket.ParkingLotID]
	if !found {
		return Car{}, errors.New("Parking lot not assinged to this attendant")
	}
	return lot.Unpark(ticket)
}
