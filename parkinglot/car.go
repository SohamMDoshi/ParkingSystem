package parkinglot

import (
	"errors"
	"regexp"

	"github.com/google/go-cmp/cmp"
)

type Car struct {
	Color           string
	RegistrationNum string
}

func (t Car) Equals(other Car) bool {
	return cmp.Equal(t, other)
}

func IsValidRegistrationNum(registartionNum string) bool {
	pattern := `^[A-Z]{2}\d{2}[A-Z]{2}\d{4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(registartionNum)
}

func NewCar(color, registartionNum string) (Car, error) {
	if color == "" {
		return Car{}, errors.New("color cannot be empty")
	}
	if registartionNum == "" {
		return Car{}, errors.New("registration number cannot be empty")
	}
	if !IsValidRegistrationNum(registartionNum) {
		return Car{}, errors.New("Invalid registration number")
	}

	return Car{
		Color:           color,
		RegistrationNum: registartionNum,
	}, nil
}
