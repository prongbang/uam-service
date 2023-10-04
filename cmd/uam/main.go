package main

import (
	_ "github.com/prongbang/goenv"
	"github.com/prongbang/user-service/internal/pkg/casbinx"
	"github.com/prongbang/user-service/internal/service"
	"github.com/prongbang/user-service/internal/service/database"
	"github.com/prongbang/user-service/internal/service/schema"
)

func main() {
	// Database
	dbDriver := database.NewDatabaseDriver()

	// Schema
	schema.New(dbDriver).Initial()

	// Casbin
	enforce := casbinx.New()

	// Service
	svc := service.New(dbDriver, enforce)

	// gRPC
	svc.NewGRPC().Register()

	// APIs
	svc.NewAPI().Register()
}
