package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/pkg/token"
)

func HttpPayload(c *fiber.Ctx) token.Claims {
	a := new(token.AccessToken)
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

func GrpcPayload(jwe string) token.Claims {
	payload, err := token.Payload(jwe)
	if err != nil {
		return token.Claims{}
	}
	return *payload
}
