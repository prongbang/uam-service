package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/common"
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
	body := GetByIdRequest{}
	err := c.BodyParser(&body)
	if err == nil && core.IsUuid(&body.ID) {
		return c.Next()
	}
	return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
}

func (v *validate) GetList(c *fiber.Ctx) error {
	paging := core.PagingBody(c)

	if paging.Invalid() {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonPagingInvalid))
	}

	return c.Next()
}

func (v *validate) Create(c *fiber.Ctx) error {
	b := CreateUser{}
	if err := c.BodyParser(&b); err != nil {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
	}

	if b.Username == "" && b.Email == "" {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
	}

	vld := validator.New()
	if err := vld.Struct(b); err != nil {
		return core.BadRequest(c, err)
	}
	if (b.Username != "" && len(b.Username) < 4) || (b.Email != "" && !common.IsEmail(b.Email)) || len(b.Password) < 8 {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
	}

	return c.Next()
}

func (v *validate) Update(c *fiber.Ctx) error {
	body := UpdateUser{}
	err := c.BodyParser(&body)
	if err != nil || !core.IsUuid(&body.ID) {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
	}

	if body.Username == "" && body.Email == "" {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
	}

	if (body.Username != "" && len(body.Username) < 4) || (body.Email != "" && !common.IsEmail(body.Email)) {
		return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
	}

	return c.Next()
}

func (v *validate) Delete(c *fiber.Ctx) error {
	body := DeleteByIdRequest{}
	err := c.BodyParser(&body)
	if err == nil && core.IsUuid(&body.ID) {
		return c.Next()
	}
	return core.BadRequest(c, core.MessageText(c, localizations.CommonInvalidData))
}

func NewValidate() Validate {
	return &validate{}
}
