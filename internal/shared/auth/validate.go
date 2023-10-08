package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/localizations"
	"github.com/prongbang/user-service/pkg/common"
	"github.com/prongbang/user-service/pkg/core"
)

type Validate interface {
	Login(c *fiber.Ctx) error
}

type validate struct {
}

func (v *validate) Login(c *fiber.Ctx) error {
	b := Login{}
	err := c.BodyParser(&b)
	if err != nil || b.Password == "" || (b.Email == "" && b.Username == "") {
		return core.Unauthorized(c, localizations.CommonInvalidData)
	}

	if b.Email != "" && b.Username != "" {
		if !common.IsEmail(b.Email) {
			return core.Unauthorized(c, localizations.CommonInvalidData)
		}
	}

	return c.Next()
}

func NewValidate() Validate {
	return &validate{}
}
