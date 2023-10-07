package core

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/localizations"
	"net/http"
)

var (
	StatusOK           = http.StatusText(http.StatusOK)
	StatusCreated      = http.StatusText(http.StatusCreated)
	StatusBadRequest   = http.StatusText(http.StatusBadRequest)
	StatusUnauthorized = http.StatusText(http.StatusUnauthorized)
	StatusForbidden    = http.StatusText(http.StatusForbidden)
	StatusNotFound     = http.StatusText(http.StatusNotFound)
)

type Error struct {
	Message string `json:"message"`
}

type Success struct {
	Message string `json:"message"`
}

type Response struct {
	Data    any `json:"data,omitempty"`
	Message any `json:"message"`
}

func SuccessData(c *fiber.Ctx, key string) Success {
	return Success{
		Message: MessageText(c, key),
	}
}

func MessageText(c *fiber.Ctx, key string) string {
	locale := c.Get(fiber.HeaderAcceptLanguage)
	if locale == "" {
		locale = localizations.En
	}
	return localizations.Translate(locale, key)
}

func Ok(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

func SendStream(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

func Created(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusCreated).JSON(data)
}

func BadRequest(c *fiber.Ctx, data ...any) error {
	if len(data) == 0 {
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Message: StatusBadRequest,
		})
	}

	// Validation Errors
	value := data[0]
	if _, okValidation := value.(validator.ValidationErrors); okValidation {
		msgMap := fiber.Map{}
		for _, e := range value.(validator.ValidationErrors) {
			msgMap[e.Field()] = fiber.Map{
				"required": fmt.Sprintf(MessageText(c, localizations.CommonFieldIsRequiredAndNotEmpty), e.Field()),
			}
		}
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Message: msgMap,
		})
	}

	// Other error
	return c.Status(http.StatusBadRequest).JSON(&Response{
		Message: MessageText(c, value.(string)),
	})
}

func NotFound(c *fiber.Ctx, data ...any) error {
	message := ""
	if len(data) == 0 {
		message = http.StatusText(http.StatusNotFound)
	} else {
		message = MessageText(c, data[0].(string))
	}
	return c.Status(http.StatusNotFound).JSON(&Response{
		Message: message,
	})
}

func NoContent(c *fiber.Ctx, data ...any) error {
	message := ""
	if len(data) == 0 {
		message = http.StatusText(http.StatusNoContent)
	} else {
		message = MessageText(c, data[0].(string))
	}
	return c.Status(http.StatusNoContent).JSON(&Response{
		Message: message,
	})
}

func Unauthorized(c *fiber.Ctx, data ...any) error {
	message := ""
	if len(data) == 0 {
		message = http.StatusText(http.StatusUnauthorized)
	} else {
		message = MessageText(c, data[0].(string))
	}
	return c.Status(http.StatusUnauthorized).JSON(&Response{
		Message: message,
	})
}

func Forbidden(c *fiber.Ctx, data ...any) error {
	message := ""
	if len(data) == 0 {
		message = http.StatusText(http.StatusForbidden)
	} else {
		message = MessageText(c, data[0].(string))
	}
	return c.Status(http.StatusForbidden).JSON(&Response{
		Message: message,
	})
}
