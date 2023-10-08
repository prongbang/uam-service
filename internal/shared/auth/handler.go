package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/pkg/core"
)

type Handler interface {
	Login(c *fiber.Ctx) error
}

type handler struct {
	Uc UseCase
}

func (h *handler) Login(c *fiber.Ctx) error {
	b := Login{}
	_ = c.BodyParser(&b)

	data, err := h.Uc.Login(b)
	if err != nil {
		return core.Unauthorized(c, err.Error())
	}

	return core.Ok(c, data)
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}
