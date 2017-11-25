package main

import (
	"fmt"

	"github.com/alextanhongpin/go-onion/car-service"
	"github.com/alextanhongpin/go-onion/internal/database"
)

func main() {

	db := database.New()

	// TODO: car service should return a router with endpoints
	car := car.New(db)
	fmt.Println("main:", car)
}
