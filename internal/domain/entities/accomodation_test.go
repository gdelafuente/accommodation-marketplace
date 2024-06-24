package entities

import "testing"

func Test_NewAccommodation_OK(t *testing.T) {
	id := "1"
	features := Features{
		Capacity: 2,
	}
	coordinates := Coordinates{
		Latitude:  41.388790,
		Longitude: 2.158990,
	}
	address, err := NewAddress("Spain", "08001", "Barcelona", "Carrer de Balmes", "1223", coordinates)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	accommodation, err := NewAccommodation(id, address, features)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if accommodation.id != id {
		t.Fatalf("Expected id to be %s, got %s", id, accommodation.id)
	}

	if accommodation.address != address {
		t.Fatalf("Expected address to be %v, got %v", address, accommodation.address)
	}

	if accommodation.features != features {
		t.Fatalf("Expected features to be %v, got %v", features, accommodation.features)
	}
}

func Test_NewAccommodation_MissingId(t *testing.T) {
	features := Features{
		Capacity: 2,
	}
	coordinates := Coordinates{
		Latitude:  41.388790,
		Longitude: 2.158990,
	}
	address, err := NewAddress("Spain", "08001", "Barcelona", "Carrer de Balmes", "1223", coordinates)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = NewAccommodation("", address, features)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingID)
	}
}

func Test_NewAccommodation_InvalidAddress(t *testing.T) {
	id := "1"
	features := Features{
		Capacity: 2,
	}

	_, err := NewAccommodation(id, Address{}, features)

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrInvalidAddress)
	}
}

func Test_NewAccommodation_MissingFeatures(t *testing.T) {
	id := "1"
	coordinates := Coordinates{
		Latitude:  41.388790,
		Longitude: 2.158990,
	}
	address, err := NewAddress("Spain", "08001", "Barcelona", "Carrer de Balmes", "1223", coordinates)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = NewAccommodation(id, address, Features{})

	if err == nil {
		t.Fatalf("Expected %v, got nil", ErrMissingFeatures)
	}
}
