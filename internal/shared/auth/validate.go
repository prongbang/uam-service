package auth

import "github.com/gofiber/fiber/v2"

type Validate interface {
	Login(c *fiber.Ctx) error
}

type validate struct {
}

func (v *validate) Login(c *fiber.Ctx) error {
	return c.Next()
}

func NewValidate() Validate {
	return &validate{}
}
