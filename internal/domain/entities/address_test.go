package entities

import (
	"testing"
)

func Test_NewAddress_OK(t *testing.T) {
	country := "Spain"
	postalCode := "08001"
	city := "Barcelona"
	street := "Carrer de Balmes"
	streetNumber := "1223"
	coordinates := Coordinates{
		Latitude:  41.388790,
		Longitude: 2.158990,
	}

	address, err := NewAddress(country, postalCode, city, street, streetNumber, coordinates)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if address.country != country {
		t.Fatalf("Expected country to be %s, got %s", country, address.country)
	}

	if address.postalCode != postalCode {
		t.Fatalf("Expected postalCode to be %s, got %s", postalCode, address.postalCode)
	}

	if address.city != city {
		t.Fatalf("Expected city to be %s, got %s", city, address.city)
	}

	if address.street != street {
		t.Fatalf("Expected street to be %s, got %s", street, address.street)
	}

	if address.streetNumber != streetNumber {
		t.Fatalf("Expected streetNumber to be %s, got %s", streetNumber, address.streetNumber)
	}

	if address.coordinates != coordinates {
		t.Fatalf("Expected coordinates to be %v, got %v", coordinates, address.coordinates)
	}
}

func Test_NewAddress_MissingCountry(t *testing.T) {
	_, err := NewAddress(
		"",
		"08001",
		"Barcelona",
		"Carrer de Balmes",
		"123",
		Coordinates{
			Latitude:  41.388790,
			Longitude: 2.158990,
		},
	)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingCountry)
	}
}

func Test_NewAddress_MissingPostalCode(t *testing.T) {
	_, err := NewAddress(
		"Spain",
		"",
		"Barcelona",
		"Carrer de Balmes",
		"123",
		Coordinates{
			Latitude:  41.388790,
			Longitude: 2.158990,
		},
	)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingPostalCode)
	}
}

func Test_NewAddress_MissingCity(t *testing.T) {
	_, err := NewAddress(
		"Spain",
		"08001",
		"",
		"Carrer de Balmes",
		"123",
		Coordinates{
			Latitude:  41.388790,
			Longitude: 2.158990,
		},
	)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingCity)
	}
}

func Test_NewAddress_MissingStreet(t *testing.T) {
	_, err := NewAddress(
		"Spain",
		"08001",
		"Barcelona",
		"",
		"123",
		Coordinates{
			Latitude:  41.388790,
			Longitude: 2.158990,
		},
	)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingStreet)
	}
}

func Test_NewAddress_MissingStreetNumber(t *testing.T) {
	_, err := NewAddress(
		"Spain",
		"08001",
		"Barcelona",
		"Carrer de Balmes",
		"",
		Coordinates{
			Latitude:  41.388790,
			Longitude: 2.158990,
		},
	)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingStreetNumber)
	}
}

func Test_NewAddress_MissingCoordinates(t *testing.T) {
	_, err := NewAddress("Spain", "08001", "Barcelona", "Carrer de Balmes", "123", Coordinates{})

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingCoordinates)
	}
}
