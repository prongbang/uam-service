package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	Handle Handler
}

// Initial implements Router.
func (r *router) Initial(app *fiber.App) {

}

func NewRouter(handle Handler) Router {
	return &router{
		Handle: handle,
	}
}
