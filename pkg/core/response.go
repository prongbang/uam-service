package core

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/code"
	"net/http"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Message string `json:"message"`
}

type Response struct {
	Code    string `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message any    `json:"message"`
	Cause   any    `json:"cause,omitempty"`
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

func Ok(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(&Response{
		Code:    fmt.Sprintf("%d", http.StatusOK),
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func OkResponse(c *fiber.Ctx, data Response) error {
	return c.Status(http.StatusOK).JSON(data)
}

func SendStream(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(&Response{
		Code:    fmt.Sprintf("%d", http.StatusCreated),
		Message: http.StatusText(http.StatusCreated),
		Data:    data,
	})
}

func BadRequest(c *fiber.Ctx, data any) error {
	if data == nil {
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Code:    code.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	// Validation Errors
	if _, okValidation := data.(validator.ValidationErrors); okValidation {
		cause := fiber.Map{}
		for _, e := range data.(validator.ValidationErrors) {
			cause[e.Field()] = fiber.Map{
				"required": fmt.Sprintf(MessageText(c, localizations.CommonFieldIsRequiredAndNotEmpty), e.Field()),
			}
		}
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Code:    code.StatusInformationRequired,
			Message: http.StatusText(http.StatusBadRequest),
			Cause:   cause,
		})
	}

	// Error
	if err, okError := data.(Error); okError {
		return c.Status(http.StatusBadRequest).JSON(&Response{
			Code:    err.Code,
			Message: err.Message,
		})
	}

	// Other error
	return c.Status(http.StatusBadRequest).JSON(&Response{
		Code:    code.StatusBadRequest,
		Message: data,
	})
}

func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusNotFound).JSON(&Response{
		Code:    code.StatusNotFound,
		Message: message,
	})
}

func NoContent(c *fiber.Ctx) error {
	return c.Status(http.StatusNoContent).JSON(&Response{
		Code:    code.StatusNoContent,
		Message: http.StatusText(http.StatusNoContent),
	})
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(&Response{
		Code:    code.StatusUnAuthorization,
		Message: message,
	})
}

func Forbidden(c *fiber.Ctx) error {
	return c.Status(http.StatusForbidden).JSON(&Response{
		Code:    code.StatusForbidden,
		Message: http.StatusText(http.StatusForbidden),
	})
}

func AttachmentCsv(c *fiber.Ctx, filename string) {
	AttachmentHeader(c, "text/csv", filename)
}

func AttachmentHeader(c *fiber.Ctx, contentType string, filename string) {
	c.Context().Response.Header.SetContentType(contentType)
	c.Context().Response.Header.SetCanonical([]byte(fiber.HeaderContentDisposition), []byte(fmt.Sprintf(`attachment; filename="%s"`, filename)))
}
