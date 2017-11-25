package car

import (
	"context"
	"log"

	"github.com/alextanhongpin/go-onion/internal/database"
)

// New returns a new car service
func New(db database.Client) string {
	model := NewModel(NewStore(db))
	ctx := context.Background()

	// Test the model, should move this to unit test
	res, err := model.One(ctx, OneRequest{Name: "hello world"})
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Create the route for the service and return it
	// route := NewRoute(model)
	// router.Get("/car", route.GetOne)
	return res.Name
}
