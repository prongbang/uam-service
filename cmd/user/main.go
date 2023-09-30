package main

import (
	"github.com/prongbang/user-service/internal/service"
	"github.com/prongbang/user-service/internal/service/database"
	"github.com/prongbang/user-service/internal/service/schema"
)

func main() {
	// Database
	dbDriver := database.NewDatabaseDriver()

	// Schema
	schema.New(dbDriver).Initial()

	// Service
	svc := service.New(dbDriver)

	// gRPC
	svc.NewGRPC().Register()

	// APIs
	svc.NewAPI().Register()
}
