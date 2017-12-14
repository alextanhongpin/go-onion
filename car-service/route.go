package car

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type route struct {
	model Model
}

func (rt route) GetOne(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// var req OneRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	req := OneRequest{
		Name: r.URL.Query().Get("name"),
	}

	res, err := rt.model.One(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// NewRoute returns a pointer to the route struct
func NewRoute(model Model) Route {
	return &route{
		model: model,
	}
}
