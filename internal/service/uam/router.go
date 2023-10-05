package uam

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/shared/role"
	"github.com/prongbang/user-service/pkg/core"
)

type APIRouter interface {
	core.Router
}

type apiRouter struct {
	RoleHandle   role.Handler
	RoleValidate role.Validate
}

// Initial implements APIRouter.
func (r *apiRouter) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Post("/role", r.RoleValidate.GetList, r.RoleHandle.GetList)
		v1.Post("/role/id", r.RoleValidate.GetById, r.RoleHandle.GetById)
		v1.Post("/role/create", r.RoleValidate.Create, r.RoleHandle.Create)
		v1.Post("/role/update", r.RoleValidate.Update, r.RoleHandle.Update)
		v1.Post("/role/delete", r.RoleValidate.Delete, r.RoleHandle.Delete)
	}
}

func NewRouter(handle role.Handler, roleValidate role.Validate) APIRouter {
	return &apiRouter{
		RoleHandle:   handle,
		RoleValidate: roleValidate,
	}
}
