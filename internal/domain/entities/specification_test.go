package entities

import (
	"testing"
	"time"
)

func Test_NewSpecification_OK(t *testing.T) {
	coordinates := Coordinates{1.0, 2.0}
	maxKmsAway := 20
	availableFrom := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location())
	availableUntil := time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location())
	minPrice := 100
	maxPrice := 200
	features := Features{
		Capacity: 2,
	}

	spec, err := NewSpecification(
		coordinates,
		maxKmsAway,
		WithAvailableFrom(availableFrom),
		WithAvailableUntil(availableUntil),
		WithMinPrice(minPrice),
		WithMaxPrice(maxPrice),
		WithFeatures(features),
	)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if spec.coordinates != coordinates {
		t.Fatalf("Expected coordinates to be %v, got %v", coordinates, spec.coordinates)
	}

	if spec.maxKmsAway != maxKmsAway {
		t.Fatalf("Expected maxKmsAway to be %d, got %d", maxKmsAway, spec.maxKmsAway)
	}

	if spec.availableFrom != availableFrom {
		t.Fatalf("Expected availableFrom to be %v, got %v", availableFrom, spec.availableFrom)
	}

	if spec.availableUntil != availableUntil {
		t.Fatalf("Expected availableUntil to be %v, got %v", availableUntil, spec.availableUntil)
	}

	if spec.minPrice != minPrice {
		t.Fatalf("Expected minPrice to be %d, got %d", minPrice, spec.minPrice)
	}

	if spec.maxPrice != maxPrice {
		t.Fatalf("Expected maxPrice to be %d, got %d", maxPrice, spec.maxPrice)
	}

	if spec.features != features {
		t.Fatalf("Expected features to be %v, got %v", features, spec.features)
	}

}

func Test_NewSpecification_MissingCoordinates(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{},
		20,
	)

	if err != ErrMissingCoordinates {
		t.Fatalf("Expected error to be %v, got %v", ErrMissingCoordinates, err)
	}

}

func Test_NewSpecification_NegativeMaxKmsAway(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{1.0, 2.0},
		-1,
	)

	if err != ErrInvalidMaxKmsAway {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidMaxKmsAway, err)
	}

}

func Test_NewSpecification_TooHighMaxKmsAway(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{1.0, 2.0},
		101,
	)

	if err != ErrInvalidMaxKmsAway {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidMaxKmsAway, err)
	}

}

func Test_NewSpecification_InvalidTimeRange(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{1.0, 2.0},
		20,
		WithAvailableFrom(time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location())),
		WithAvailableUntil(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location())),
	)

	if err != ErrInvalidTimeRange {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidTimeRange, err)
	}

}

func Test_NewSpecification_InvalidMinPrice(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{1.0, 2.0},
		20,
		WithMinPrice(-1),
	)

	if err != ErrInvalidMinPrice {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidMinPrice, err)
	}

}

func Test_NewSpecification_InvalidMaxPrice(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{1.0, 2.0},
		20,
		WithMaxPrice(-1),
	)

	if err != ErrInvalidMaxPrice {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidMaxPrice, err)
	}

}

func Test_NewSpecification_InvalidPriceRange(t *testing.T) {
	_, err := NewSpecification(
		Coordinates{1.0, 2.0},
		20,
		WithMinPrice(200),
		WithMaxPrice(100),
	)

	if err != ErrInvalidPriceRange {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidPriceRange, err)
	}

}
