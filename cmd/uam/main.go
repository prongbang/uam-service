package main

import (
	_ "github.com/prongbang/goenv"
	"github.com/prongbang/uam-service/internal/pkg/casbinx"
	"github.com/prongbang/uam-service/internal/service"
	"github.com/prongbang/uam-service/internal/service/database"
	"github.com/prongbang/uam-service/internal/service/schema"
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
