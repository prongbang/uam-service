package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/pkg/core"
)

type Validate interface {
	core.Handler
	GetByMe(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	UpdatePasswordMe(c *fiber.Ctx) error
}

type validate struct {
}

func (v *validate) UpdatePassword(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) UpdatePasswordMe(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) GetByMe(c *fiber.Ctx) error {
	return c.Next()
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
