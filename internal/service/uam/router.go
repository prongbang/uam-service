package uam

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/shared/auth"
	"github.com/prongbang/uam-service/internal/shared/forgot"
	"github.com/prongbang/uam-service/internal/shared/role"
	"github.com/prongbang/uam-service/internal/shared/user"
	"github.com/prongbang/uam-service/internal/shared/user_role"
	"github.com/prongbang/uam-service/pkg/core"
)

type APIRouter interface {
	core.Router
}

type apiRouter struct {
	RoleHandle       role.Handler
	AuthHandle       auth.Handler
	UserHandle       user.Handler
	ForgotHandle     forgot.Handler
	UserRoleHandle   user_role.Handler
	RoleValidate     role.Validate
	AuthValidate     auth.Validate
	UserValidate     user.Validate
	ForgotValidate   forgot.Validate
	UserRoleValidate user_role.Validate
}

// Initial implements APIRouter.
func (r *apiRouter) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Post("/role/read/all", r.RoleValidate.GetList, r.RoleHandle.GetList)
		v1.Post("/role/read", r.RoleValidate.GetById, r.RoleHandle.GetById)
		v1.Post("/role/create", r.RoleValidate.Create, r.RoleHandle.Create)
		v1.Post("/role/update", r.RoleValidate.Update, r.RoleHandle.Update)
		v1.Post("/role/delete", r.RoleValidate.Delete, r.RoleHandle.Delete)

		v1.Post("/auth/login", r.AuthValidate.Login, r.AuthHandle.Login)

		v1.Post("/user/me", r.UserValidate.GetByMe, r.UserHandle.GetByMe)
		v1.Post("/user/read", r.UserValidate.GetById, r.UserHandle.GetById)
		v1.Post("/user/read/list", r.UserValidate.GetList, r.UserHandle.GetList)
		v1.Post("/user/create", r.UserValidate.Create, r.UserHandle.Create)
		v1.Post("/user/update", r.UserValidate.Update, r.UserHandle.Update)
		v1.Post("/user/update/pwd", r.UserValidate.UpdatePassword, r.UserHandle.UpdatePassword)
		v1.Post("/user/update/pwd/me", r.UserValidate.UpdatePasswordMe, r.UserHandle.UpdatePasswordMe)
		v1.Post("/user/delete", r.UserValidate.Delete, r.UserHandle.Delete)
	}
}

func NewRouter(
	roleHandle role.Handler,
	authHandle auth.Handler,
	userHandle user.Handler,
	forgotHandle forgot.Handler,
	userRoleHandle user_role.Handler,
	roleValidate role.Validate,
	authValidate auth.Validate,
	userValidate user.Validate,
	forgotValidate forgot.Validate,
	userRoleValidate user_role.Validate,
) APIRouter {
	return &apiRouter{
		RoleHandle:       roleHandle,
		AuthHandle:       authHandle,
		UserHandle:       userHandle,
		ForgotHandle:     forgotHandle,
		UserRoleHandle:   userRoleHandle,
		RoleValidate:     roleValidate,
		AuthValidate:     authValidate,
		UserValidate:     userValidate,
		ForgotValidate:   forgotValidate,
		UserRoleValidate: userRoleValidate,
	}
}
