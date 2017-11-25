package car

import (
	"context"
)

type model struct {
	store Store
}

func (m model) One(ctx context.Context, req OneRequest) (*OneResponse, error) {
	// Perform validation and composition of business logic here
	res, err := m.store.One(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewModel returns a new model that wraps the store
func NewModel(store Store) Model {
	return &model{
		store: store,
	}
}
