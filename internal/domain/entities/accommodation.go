package entities

import (
	"errors"
)

var (
	ErrMissingID       = errors.New("ID is required")
	ErrMissingFeatures = errors.New("features are required")
	ErrInvalidAddress  = errors.New("address is invalid")
)

// Private fields ensure the constructor is used, so validation is always performed
type Accommodation struct {
	id           string
	address      Address
	features     Features
	availability Availability
}

func (a Accommodation) validate() error {
	if a.id == "" {
		return ErrMissingID
	}
	if a.features == (Features{}) {
		return ErrMissingFeatures
	}
	err := a.address.validate()
	if err != nil {
		return ErrInvalidAddress
	}
	return nil
}

func NewAccommodation(id string, address Address, features Features) (Accommodation, error) {
	a := Accommodation{
		id:       id,
		address:  address,
		features: features,
	}

	return a, a.validate()
}

type Features struct {
	Capacity         int
	Kitchen          bool
	Parking          bool
	Wifi             bool
	AirCon           bool
	Heating          bool
	WheelchairAccess bool
	Garden           bool
	Pool             bool
	PetFriendly      bool
}
