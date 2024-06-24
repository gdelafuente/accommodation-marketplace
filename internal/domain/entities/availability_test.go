package entities

import (
	"testing"
	"time"
)

func Test_NewAvailability_OK(t *testing.T) {
	availability, err := NewAvailability()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(availability.slots) != 0 {
		t.Fatalf("Expected no slots, got %d", len(availability.slots))
	}
}

func Test_NewAvailabilitySlot_OK(t *testing.T) {
	startTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location())
	endTime := time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location())
	pricePerNight := 100.0

	slot, err := NewAvailabilitySlot(startTime, endTime, pricePerNight)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if slot.startTime != startTime {
		t.Fatalf("Expected startTime to be %v, got %v", startTime, slot.startTime)
	}

	if slot.endTime != endTime {
		t.Fatalf("Expected endTime to be %v, got %v", endTime, slot.endTime)
	}

	if slot.pricePerNight != pricePerNight {
		t.Fatalf("Expected pricePerNight to be %f, got %f", pricePerNight, slot.pricePerNight)
	}
}

func Test_NewAvailabilitySlot_MissingStartTime(t *testing.T) {
	_, err := NewAvailabilitySlot(
		time.Time{},
		time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location()),
		100.0,
	)

	if err != ErrMissingStartTime {
		t.Fatalf("Expected error to be %v, got %v", ErrMissingStartTime, err)
	}
}

func Test_NewAvailabilitySlot_MissingEndTime(t *testing.T) {
	_, err := NewAvailabilitySlot(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		time.Time{},
		100.0,
	)

	if err != ErrMissingEndTime {
		t.Fatalf("Expected error to be %v, got %v", ErrMissingEndTime, err)
	}
}

func Test_NewAvailabilitySlot_InvalidTimeRange(t *testing.T) {
	_, err := NewAvailabilitySlot(
		time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location()),
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		100.0,
	)

	if err != ErrInvalidTimeRange {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidTimeRange, err)
	}
}

func Test_NewAvailabilitySlot_InvalidPricePerNight(t *testing.T) {
	_, err := NewAvailabilitySlot(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location()),
		0.0,
	)

	if err != ErrInvalidPricePerNight {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidPricePerNight, err)
	}

	_, err = NewAvailabilitySlot(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		time.Date(2021, 1, 2, 0, 0, 0, 0, time.Now().Location()),
		-1.0,
	)

	if err != ErrInvalidPricePerNight {
		t.Fatalf("Expected error to be %v, got %v", ErrInvalidPricePerNight, err)
	}
}
