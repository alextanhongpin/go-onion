package main

import (
	"log"
	"net/http"

	"github.com/alextanhongpin/go-onion/car-service"
	"github.com/alextanhongpin/go-onion/internal/database"
	"github.com/alextanhongpin/go-onion/internal/schema"

	"github.com/julienschmidt/httprouter"
)

var schemas map[string]string

func init() {
	schemas = make(map[string]string)
	// TODO: Serve static json schemas and load them from the url
	schemas["car:one_request"] = "file:///Users/alextanhongpin/Documents/golang/src/github.com/alextanhongpin/go-onion/schema/one.json"
}

const port = ":8080"

func main() {

	router := httprouter.New()
	db := database.New()
	s := schema.New(schemas)

	// TODO: car service should return a router with endpoints
	router = car.New(db, router, s)

	log.Printf("listening to port *%s. press ctrl + c to cancel.\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
