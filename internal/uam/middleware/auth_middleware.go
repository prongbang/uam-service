package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/pkg/casbinx"
)

type AuthMiddleware interface {
	New() fiber.Handler
}

type authMiddleware struct {
	CasbinXs casbinx.CasbinXs
}

func (a *authMiddleware) New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

func NewAuthMiddleware(casbinXs casbinx.CasbinXs) AuthMiddleware {
	return &authMiddleware{
		CasbinXs: casbinXs,
	}
}
