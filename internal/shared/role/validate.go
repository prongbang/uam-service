package role

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/core"
)

type Validate interface {
	core.Handler
}

type validate struct {
}

func (v *validate) GetById(c *fiber.Ctx) error {
	b := GetByIdRequest{}
	err := c.BodyParser(&b)
	if err != nil || b.ID != "" {
		return c.Next()
	}
	return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
}

func (v *validate) GetList(c *fiber.Ctx) error {
	return c.Next()
}

func (v *validate) Create(c *fiber.Ctx) error {
	b := CreateRole{}
	err := c.BodyParser(&b)
	if err != nil {
		return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
	}

	vx := validator.New()
	if err = vx.Struct(b); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}
	return c.Next()
}

func (v *validate) Update(c *fiber.Ctx) error {
	b := UpdateRole{}
	err := c.BodyParser(&b)
	if err != nil || b.ID == "" {
		return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
	}

	vx := validator.New()
	if err = vx.Struct(b); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}
	return c.Next()
}

func (v *validate) Delete(c *fiber.Ctx) error {
	b := DelByIdRequest{}
	err := c.BodyParser(&b)
	if err != nil || b.ID != "" {
		return c.Next()
	}
	return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
}

func NewValidate() Validate {
	return &validate{}
}
