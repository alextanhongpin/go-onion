package car

import (
	"context"

	"github.com/alextanhongpin/go-onion/internal/database"
)

type store struct {
	db database.Client
}

func (s store) One(ctx context.Context, req OneRequest) (*OneResponse, error) {
	res := s.db.Query("SELECT * FROM *")
	return &OneResponse{
		Name: res + ": " + req.Name,
	}, nil
}

// NewStore returns an instance of new store which is basically a repository pattern making connections to the database
func NewStore(db database.Client) Store {
	return &store{
		db: db,
	}
}
