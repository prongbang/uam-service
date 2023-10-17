package main

import (
	_ "github.com/prongbang/goenv"
	"github.com/prongbang/uam-service/internal/pkg/casbinx"
	"github.com/prongbang/uam-service/internal/uam"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/prongbang/uam-service/internal/uam/schema"
)

func main() {
	// Database
	dbDriver := database.NewDatabaseDriver()

	// Schema
	schema.New(dbDriver).Initial()

	// Casbin
	enforce := casbinx.New()

	// Services
	svc := uam.New(dbDriver, enforce)

	// gRPC
	svc.NewGRPC().Register()

	// APIs
	svc.NewAPI().Register()
}
