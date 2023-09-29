package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/user"
	"github.com/prongbang/user-service/pkg/core"
)

type Routers interface {
	core.Routers
}

type routers struct {
	UserRouter user.Router
}

func (r *routers) Initials(app *fiber.App) {
	r.UserRouter.Initial(app)
}

func NewRouters(userRouter user.Router) Routers {
	return &routers{
		UserRouter: userRouter,
	}
}
