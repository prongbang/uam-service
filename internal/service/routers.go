package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/service/user"
	"github.com/prongbang/user-service/pkg/core"
)

type Routers interface {
	core.Routers
}

type routers struct {
	UserRouter user.APIRouter
}

func (r *routers) Initials(app *fiber.App) {
	r.UserRouter.Initial(app)
}

func NewRouters(userRouter user.APIRouter) Routers {
	return &routers{
		UserRouter: userRouter,
	}
}
