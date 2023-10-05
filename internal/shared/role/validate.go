package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/pkg/core"
)

type Validate interface {
	core.Handler
}

type validate struct {
}

func (v *validate) GetById(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) GetList(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Create(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Update(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Delete(c *fiber.Ctx) error {
	return c.Next()
}

func NewValidate() Validate {
	return &validate{}
}
