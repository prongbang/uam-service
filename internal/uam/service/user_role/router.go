package user_role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	UserRoleHandle   Handler
	UserRoleValidate Validate
}

// Initial implements Router.
func (r *router) Initial(app *fiber.App) {
}

func NewRouter(userRoleHandle Handler, userRoleValidate Validate) Router {
	return &router{
		UserRoleHandle:   userRoleHandle,
		UserRoleValidate: userRoleValidate,
	}
}
