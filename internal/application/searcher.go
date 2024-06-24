package application

import (
	"context"
	"errors"

	"github.com/gdelafuente/accommodation-marketplace/internal/domain/entities"
)

var (
	ErrAccommodationsViewFailed = errors.New("failed to fetch accommodations")
)

type Searcher struct {
	view entities.AccommodationsView
}

func NewSearcher(view entities.AccommodationsView) *Searcher {
	return &Searcher{view: view}
}

// TODO: this should receive an application layer DTO instead of a domain entity
func (s *Searcher) Search(ctx context.Context, spec entities.Specification) ([]entities.Accommodation, error) {
	// TODO: caching

	accommodations, err := s.view.FetchAccommodationsForSpecification(ctx, spec)
	if err != nil {
		return nil, errors.Join(ErrAccommodationsViewFailed, err)
	}

	// TODO: ranking

	return accommodations, nil
}
