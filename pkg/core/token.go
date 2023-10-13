package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/pkg/token"
)

type AccessToken struct {
	Token string `json:"token"`
}

func Payload(c *fiber.Ctx) token.Claims {
	a := new(AccessToken)
	err := c.BodyParser(a)
	if err != nil {
		return token.Claims{}
	}
	payload, err := token.Payload(a.Token)
	if err != nil {
		return token.Claims{}
	}
	return *payload
}
