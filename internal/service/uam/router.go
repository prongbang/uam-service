package uam

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/pkg/core"
)

type APIRouter interface {
	core.Router
}

type apiRouter struct {
	Handle APIHandler
}

// Initial implements APIRouter.
func (r *apiRouter) Initial(app *fiber.App) {

}

func NewRouter(handle APIHandler) APIRouter {
	return &apiRouter{
		Handle: handle,
	}
}
