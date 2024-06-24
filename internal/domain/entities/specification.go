package entities

import (
	"errors"
	"time"
)

var (
	ErrInvalidMaxKmsAway = errors.New("invalid max kms away")
	ErrInvalidMinPrice   = errors.New("invalid min price")
	ErrInvalidMaxPrice   = errors.New("invalid max price")
	ErrInvalidPriceRange = errors.New("invalid price range")
)

type Specification struct {
	coordinates    Coordinates
	maxKmsAway     int
	availableFrom  time.Time
	availableUntil time.Time
	minPrice       int
	maxPrice       int
	features       Features
}

type SpecificationOpt func(*Specification)

func (q Specification) validate() error {
	if q.coordinates == (Coordinates{}) {
		return ErrMissingCoordinates
	}
	if q.maxKmsAway < 0 || q.maxKmsAway > 100 {
		return ErrInvalidMaxKmsAway
	}
	if !q.availableFrom.IsZero() && !q.availableUntil.IsZero() && q.availableFrom.After(q.availableUntil) {
		return ErrInvalidTimeRange
	}
	if q.minPrice < 0 {
		return ErrInvalidMinPrice
	}
	if q.maxPrice < 0 {
		return ErrInvalidMaxPrice
	}
	if q.minPrice > q.maxPrice {
		return ErrInvalidPriceRange
	}
	return nil
}

func NewSpecification(coordinates Coordinates, maxKmsAway int, opts ...SpecificationOpt) (Specification, error) {
	query := &Specification{
		coordinates: coordinates,
		maxKmsAway:  maxKmsAway,
	}
	for _, opt := range opts {
		opt(query)
	}

	return *query, query.validate()
}

func WithMinPrice(val int) SpecificationOpt {
	return func(q *Specification) {
		q.minPrice = val
	}
}

func WithMaxPrice(val int) SpecificationOpt {
	return func(q *Specification) {
		q.maxPrice = val
	}
}

func WithAvailableFrom(val time.Time) SpecificationOpt {
	return func(q *Specification) {
		q.availableFrom = val
	}
}

func WithAvailableUntil(val time.Time) SpecificationOpt {
	return func(q *Specification) {
		q.availableUntil = val
	}
}

func WithFeatures(val Features) SpecificationOpt {
	return func(q *Specification) {
		q.features = val
	}
}
