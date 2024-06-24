package entities

import "errors"

var (
	ErrMissingCountry      = errors.New("the country is required")
	ErrMissingPostalCode   = errors.New("the postal code is required")
	ErrMissingCity         = errors.New("the city is required")
	ErrMissingStreet       = errors.New("the street is required")
	ErrMissingStreetNumber = errors.New("the street number is required")
	ErrMissingCoordinates  = errors.New("the coordinates are required")
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Private fields ensure the constructor is used, so validation is always performed
type Address struct {
	coordinates  Coordinates
	country      string
	postalCode   string
	city         string
	street       string
	streetNumber string
}

func (a Address) validate() error {
	if a.country == "" {
		return ErrMissingCountry
	}
	if a.postalCode == "" {
		return ErrMissingPostalCode
	}
	if a.city == "" {
		return ErrMissingCity
	}
	if a.street == "" {
		return ErrMissingStreet
	}
	if a.streetNumber == "" {
		return ErrMissingStreetNumber
	}
	if a.coordinates == (Coordinates{}) {
		return ErrMissingCoordinates
	}
	return nil
}

func NewAddress(country string, postalCode string, city string, street string, streetNumber string, coordinates Coordinates) (Address, error) {
	// TODO: ensure the coordinates match the other fields. Have a factory to generate them calling a geospatial service?
	a := Address{
		country:      country,
		postalCode:   postalCode,
		city:         city,
		street:       street,
		streetNumber: streetNumber,
		coordinates:  coordinates,
	}

	return a, a.validate()
}
