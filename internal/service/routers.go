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
	UamRouter uam.APIRouter
}

func (r *routers) Initials(app *fiber.App) {
	r.UamRouter.Initial(app)
}

func NewRouters(userRouter uam.APIRouter) Routers {
	return &routers{
		UamRouter: userRouter,
	}
}
