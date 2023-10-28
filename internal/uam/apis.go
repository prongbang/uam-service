package uam

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/prongbang/uam-service/internal/uam/middleware"
)

type API interface {
	Register()
}

type api struct {
	Routers     Routers
	Middlewares middleware.Middlewares
}

func (a *api) Register() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "X-Platform, X-Api-Key, Authorization, Access-Control-Allow-Credentials, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
	}))
	app.Use(a.Middlewares.Auth.New())

	// Routers
	a.Routers.Initials(app)

	// Serve
	_ = app.Listen(":9001")
}

func NewAPI(router Routers, middlewares middleware.Middlewares) API {
	return &api{
		Routers:     router,
		Middlewares: middlewares,
	}
}
