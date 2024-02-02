package parkinglot

type SlotStatus int

const (
	Empty SlotStatus = iota
	Full
)

type Slot struct {
	ID     int
	Status SlotStatus
	Car    Car
}
