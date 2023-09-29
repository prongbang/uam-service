package main

import (
	"github.com/prongbang/user-service/internal/api"
	"github.com/prongbang/user-service/internal/database"
	"github.com/prongbang/user-service/internal/grpc"
)

func main() {
	// Database
	dbDriver := database.NewDatabaseDriver()

	// gRPC
	grpcs := grpc.CreateGRPC(dbDriver)
	grpcs.Register()

	// APIs
	apis := api.CreateAPI(dbDriver)
	apis.Register()
}
