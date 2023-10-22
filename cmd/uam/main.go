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
	enforceRbac := casbinx.NewCasbinX(casbinx.ModelRbacPolicy)
	enforceRest := casbinx.NewCasbinX(casbinx.ModelRestPolicy)
	enforceGrpc := casbinx.NewCasbinX(casbinx.ModelGrpcPolicy)
	casbinXs := casbinx.New(enforceRbac, enforceRest, enforceGrpc)

	// Services
	svc := uam.New(dbDriver, casbinXs)

	// gRPC
	svc.NewGRPC().Register()

	// APIs
	svc.NewAPI().Register()
}
