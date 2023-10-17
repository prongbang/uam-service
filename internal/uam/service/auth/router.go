package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	AuthHandle   Handler
	AuthValidate Validate
}

// Initial implements Router.
func (r *router) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Post("/auth/login", r.AuthValidate.Login, r.AuthHandle.Login)
	}
}

func NewRouter(authHandle Handler, authValidate Validate) Router {
	return &router{
		AuthHandle:   authHandle,
		AuthValidate: authValidate,
	}
}
