package application

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/gdelafuente/accommodation-marketplace/internal/domain/entities"
)

var (
	viewAccommodations = []entities.Accommodation{
		entities.Accommodation{},
		entities.Accommodation{},
	}
)

func buildAccommodationSpecification(t *testing.T) entities.Specification {
	coordinates := entities.Coordinates{
		Latitude:  1.0,
		Longitude: 2.0,
	}

	spec, err := entities.NewSpecification(coordinates, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	return spec
}

func buildMockAccommodationsView() *entities.AccommodationsViewMock {
	return &entities.AccommodationsViewMock{
		FetchAccommodationsForSpecificationFunc: func(ctx context.Context, spec entities.Specification) ([]entities.Accommodation, error) {
			return viewAccommodations, nil
		},
	}
}

func Test_Search_OK(t *testing.T) {
	view := buildMockAccommodationsView()
	searcher := NewSearcher(view)

	spec := buildAccommodationSpecification(t)
	result, err := searcher.Search(context.Background(), spec)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	viewCalls := view.FetchAccommodationsForSpecificationCalls()
	if len(viewCalls) != 1 {
		t.Fatalf("Unexpected read file calls count: %d", len(viewCalls))
	}
	if viewCalls[0].Spec != spec {
		t.Fatalf("Unxpected spec: %v", viewCalls[0].Spec)
	}

	if !reflect.DeepEqual(result, viewAccommodations) {
		t.Fatalf("Unxpected result: %v", result)
	}

}

func Test_Search_ViewError(t *testing.T) {
	theError := errors.New("kaboom")
	view := &entities.AccommodationsViewMock{
		FetchAccommodationsForSpecificationFunc: func(ctx context.Context, spec entities.Specification) ([]entities.Accommodation, error) {
			return nil, theError
		},
	}
	searcher := NewSearcher(view)

	_, err := searcher.Search(context.Background(), buildAccommodationSpecification(t))

	if !errors.Is(err, ErrAccommodationsViewFailed) || !errors.Is(err, theError) {
		t.Fatalf("Expected error wrapped with ErrAccommodationsViewFailed, got %v", err)
	}

}
