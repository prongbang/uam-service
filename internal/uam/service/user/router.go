package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	UserHandle   Handler
	UserValidate Validate
}

// Initial implements Router.
func (r *router) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
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

func NewRouter(userHandle Handler, userValidate Validate) Router {
	return &router{
		UserHandle:   userHandle,
		UserValidate: userValidate,
	}
}
