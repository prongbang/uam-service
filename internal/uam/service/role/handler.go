package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/core"
)

type Handler interface {
	core.Handler
}

type handler struct {
	RoleUc UseCase
}

func (h *handler) GetById(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)

	b := GetByIdRequest{}
	_ = c.BodyParser(&b)

	params := ParamsGetById{ID: b.ID, UserID: payload.Sub}
	data := h.RoleUc.GetById(params)
	if data.ID == "" {
		return core.NotFound(c)
	}

	return core.Ok(c, data)
}

func (h *handler) GetList(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)

	params := Params{
		UserID: payload.Sub,
	}
	res := h.RoleUc.GetList(params)

	return core.Ok(c, res)
}

func (h *handler) Create(c *fiber.Ctx) error {
	b := CreateRole{}
	_ = c.BodyParser(&b)

	if err := h.RoleUc.Add(&b); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}

	return core.Created(c, b)
}

func (h *handler) Update(c *fiber.Ctx) error {
	b := UpdateRole{}
	_ = c.BodyParser(&b)

	if err := h.RoleUc.Update(&b); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}

	return core.Ok(c, b)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)

	b := GetByIdRequest{}
	_ = c.BodyParser(&b)

	if err := h.RoleUc.DeleteByRole(payload.Roles, b.ID); err != nil {
		return core.BadRequest(c, core.Translate(c, localizations.CommonCannotDeletePleaseTryAgain))
	}

	return core.Ok(c, core.SuccessData(c, core.Translate(c, localizations.CommonDeleteSuccess)))
}

func NewHandler(roleUc UseCase) Handler {
	return &handler{
		RoleUc: roleUc,
	}
}
