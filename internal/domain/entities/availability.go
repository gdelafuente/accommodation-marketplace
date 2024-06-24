package entities

import (
	"errors"
	"time"
)

var (
	ErrMissingStartTime     = errors.New("start time is required")
	ErrMissingEndTime       = errors.New("end time is required")
	ErrInvalidTimeRange     = errors.New("start time must be before end time")
	ErrInvalidPricePerNight = errors.New("price per night must be greater than 0")
)

type Availability struct {
	slots []AvailabilitySlot
}

func (a Availability) validate() error {
	return nil
}

func NewAvailability() (Availability, error) {
	a := Availability{}
	return a, a.validate()
}

type AvailabilitySlot struct {
	startTime     time.Time
	endTime       time.Time
	pricePerNight float64 // TODO: this is naive, pricing should be calculated based on user input and different strategies
}

func (s AvailabilitySlot) validate() error {
	if s.startTime.IsZero() {
		return ErrMissingStartTime
	}
	if s.endTime.IsZero() {
		return ErrMissingEndTime
	}
	if s.startTime.After(s.endTime) {
		return ErrInvalidTimeRange
	}
	if s.pricePerNight <= 0 {
		return ErrInvalidPricePerNight
	}

	return nil
}

func NewAvailabilitySlot(startTime time.Time, endTime time.Time, pricePerNight float64) (AvailabilitySlot, error) {
	s := AvailabilitySlot{
		startTime:     startTime,
		endTime:       endTime,
		pricePerNight: pricePerNight,
	}
	return s, s.validate()
}
