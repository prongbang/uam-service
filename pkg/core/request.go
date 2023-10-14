package core

import (
	"github.com/gofiber/fiber/v2"
)

type PagingRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p PagingRequest) Invalid() bool {
	return p.Page <= 0 || (p.Limit > 0 && p.Limit > 100)
}

func PagingBody(c *fiber.Ctx, limit ...int) PagingRequest {
	body := PagingRequest{}
	if err := c.BodyParser(&body); err != nil {
		body.Page = 1
		body.Limit = PagingLimitDefault
	}

	if body.Page == 0 {
		body.Page = 1
	}

	if len(limit) == 0 {
		body.Limit = PagingLimitDefault
	}

	return body
}
