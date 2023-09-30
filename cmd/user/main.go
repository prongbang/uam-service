package main

import (
	"github.com/prongbang/user-service/internal/service"
	"github.com/prongbang/user-service/internal/service/database"
)

func main() {
	// Database
	dbDriver := database.NewDatabaseDriver()

	// Service
	svc := service.New(dbDriver)

	// gRPC
	grpcs := svc.NewGRPC()
	grpcs.Register()

	// APIs
	apis := svc.NewAPI()
	apis.Register()
}
