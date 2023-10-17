package forgot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	ForgotHandle   Handler
	ForgotValidate Validate
}

// Initial implements Router.
func (r *router) Initial(app *fiber.App) {
}

func NewRouter(forgotHandle Handler, forgotValidate Validate) Router {
	return &router{
		ForgotHandle:   forgotHandle,
		ForgotValidate: forgotValidate,
	}
}
