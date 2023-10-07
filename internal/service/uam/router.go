package uam

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/shared/auth"
	"github.com/prongbang/user-service/internal/shared/role"
	"github.com/prongbang/user-service/pkg/core"
)

type APIRouter interface {
	core.Router
}

type apiRouter struct {
	RoleHandle   role.Handler
	AuthHandle   auth.Handler
	RoleValidate role.Validate
	AuthValidate auth.Validate
}

// Initial implements APIRouter.
func (r *apiRouter) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Post("/role/get/all", r.RoleValidate.GetList, r.RoleHandle.GetList)
		v1.Post("/role/get", r.RoleValidate.GetById, r.RoleHandle.GetById)
		v1.Post("/role/create", r.RoleValidate.Create, r.RoleHandle.Create)
		v1.Post("/role/update", r.RoleValidate.Update, r.RoleHandle.Update)
		v1.Post("/role/delete", r.RoleValidate.Delete, r.RoleHandle.Delete)
		v1.Post("/auth/login", r.AuthValidate.Login, r.AuthHandle.Login)
	}
}

func NewRouter(
	roleHandle role.Handler,
	authHandle auth.Handler,
	roleValidate role.Validate,
	authValidate auth.Validate,
) APIRouter {
	return &apiRouter{
		RoleHandle:   roleHandle,
		AuthHandle:   authHandle,
		RoleValidate: roleValidate,
		AuthValidate: authValidate,
	}
}
