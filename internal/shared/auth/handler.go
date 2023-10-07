package auth

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Login(c *fiber.Ctx) error
}

type handler struct {
	Uc UseCase
}

func (h *handler) Login(c *fiber.Ctx) error {

	return nil
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}
