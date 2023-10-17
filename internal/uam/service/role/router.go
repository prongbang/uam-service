package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	RoleHandle   Handler
	RoleValidate Validate
}

// Initial implements Router.
func (r *router) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Post("/role/read/all", r.RoleValidate.GetList, r.RoleHandle.GetList)
		v1.Post("/role/read", r.RoleValidate.GetById, r.RoleHandle.GetById)
		v1.Post("/role/create", r.RoleValidate.Create, r.RoleHandle.Create)
		v1.Post("/role/update", r.RoleValidate.Update, r.RoleHandle.Update)
		v1.Post("/role/delete", r.RoleValidate.Delete, r.RoleHandle.Delete)
	}
}

func NewRouter(roleHandle Handler, roleValidate Validate) Router {
	return &router{
		RoleHandle:   roleHandle,
		RoleValidate: roleValidate,
	}
}
