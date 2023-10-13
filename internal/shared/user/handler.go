package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/core"
)

type Handler interface {
	core.Handler
	GetByMe(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	UpdatePasswordMe(c *fiber.Ctx) error
}

type handler struct {
	UserUc UseCase
}

func (h *handler) UpdatePassword(c *fiber.Ctx) error {
	return c.Next()
}

func (h *handler) UpdatePasswordMe(c *fiber.Ctx) error {
	return c.Next()
}

func (h *handler) GetByMe(c *fiber.Ctx) error {
	payload := core.Payload(c)

	data := h.UserUc.GetById(payload.Sub)
	if *data.ID == "" {
		return core.NotFound(c, core.MessageText(c, localizations.CommonNotFoundData))
	}

	// Reset sensitive data
	data.Password = ""

	return core.Ok(c, data)
}

func (h *handler) GetById(c *fiber.Ctx) error {
	body := GetByIdRequest{}
	_ = c.BodyParser(&body)

	data := h.UserUc.GetById(body.ID)
	if *data.ID == "" {
		return core.NotFound(c, core.MessageText(c, localizations.CommonNotFoundData))
	}

	// Reset sensitive data
	data.Password = ""

	return core.Ok(c, data)
}

func (h *handler) GetList(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) Create(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) Update(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) Delete(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewHandler(userUc UseCase) Handler {
	return &handler{
		UserUc: userUc,
	}
}
