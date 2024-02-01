package parkinglot

type SlotStatus int

const (
	Empty SlotStatus = iota
	Full
)

type Slot struct {
	Status SlotStatus
	Car    Car
}
