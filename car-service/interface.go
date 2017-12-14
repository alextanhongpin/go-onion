package car

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Store defines the contracts with the data source such as database
// of external http requests. This layer is mockable.
type Store interface {
	One(context.Context, OneRequest) (*OneResponse, error)
}

// Model contains business logic or composition of business logic
// and does not access the database directly. Validation should be
// carried out here
type Model interface {
	One(context.Context, OneRequest) (*OneResponse, error)
}

// Route is the endpoint for our services where the model is called
type Route interface {
	GetOne(http.ResponseWriter, *http.Request, httprouter.Params)
}

// OneRequest is a sample request to the service
type OneRequest struct {
	Name string `json:"name"`
}

// OneResponse is a sample response from the one service
type OneResponse struct {
	Name string `json:"name"`
}
