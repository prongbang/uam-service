package uam

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/uam/service/auth"
	"github.com/prongbang/uam-service/internal/uam/service/forgot"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"github.com/prongbang/uam-service/internal/uam/service/user_role"
	"github.com/prongbang/uam-service/pkg/core"
)

type Routers interface {
	core.Routers
}

type routers struct {
	AuthRouter     auth.Router
	ForgotRouter   forgot.Router
	RoleRouter     role.Router
	UserRouter     user.Router
	UserRoleRouter user_role.Router
}

func (r *routers) Initials(app *fiber.App) {
	r.AuthRouter.Initial(app)
	r.ForgotRouter.Initial(app)
	r.RoleRouter.Initial(app)
	r.UserRouter.Initial(app)
	r.UserRoleRouter.Initial(app)
}

func NewRouters(
	authRouter auth.Router,
	forgotRouter forgot.Router,
	roleRouter role.Router,
	userRouter user.Router,
	userRoleRouter user_role.Router,
) Routers {
	return &routers{
		AuthRouter:     authRouter,
		ForgotRouter:   forgotRouter,
		RoleRouter:     roleRouter,
		UserRouter:     userRouter,
		UserRoleRouter: userRoleRouter,
	}
}
