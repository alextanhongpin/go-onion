package car

import (
	"github.com/alextanhongpin/go-onion/internal/database"
	"github.com/alextanhongpin/go-onion/internal/schema"

	"github.com/julienschmidt/httprouter"
)

// New returns a new car service
func New(db database.Client, router *httprouter.Router, schema *schema.Loader) *httprouter.Router {
	route := NewRoute(NewModel(NewStore(db), schema))

	// TODO: Correct the naming, as it will only return one car for now
	router.GET("/cars", route.GetOne)
	return router
}
