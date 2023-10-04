package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/service/uam"
	"github.com/prongbang/user-service/pkg/core"
)

type Routers interface {
	core.Routers
}

type routers struct {
	UserRouter uam.APIRouter
}

func (r *routers) Initials(app *fiber.App) {
	r.UserRouter.Initial(app)
}

func NewRouters(userRouter uam.APIRouter) Routers {
	return &routers{
		UserRouter: userRouter,
	}
}
