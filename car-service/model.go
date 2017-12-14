package car

import (
	"context"

	"github.com/alextanhongpin/go-onion/internal/schema"
)

type model struct {
	store  Store
	schema *schema.Loader
}

func (m model) One(ctx context.Context, req OneRequest) (*OneResponse, error) {
	// Perform validation and composition of business logic here
	if errors := m.schema.Validate("car:one_request", req); len(errors) > 0 {
		return nil, errors[0]
	}
	return m.store.One(ctx, req)
}

// NewModel returns a new model that wraps the store
func NewModel(store Store, schema *schema.Loader) Model {
	return &model{
		store:  store,
		schema: schema,
	}
}
