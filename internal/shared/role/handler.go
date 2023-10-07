package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/user-service/internal/localizations"
	"github.com/prongbang/user-service/internal/shared/user"
	"github.com/prongbang/user-service/pkg/core"
)

type Handler interface {
	core.Handler
}

type handler struct {
	RoleUc UseCase
	UserUc user.UseCase
}

func (h *handler) GetById(c *fiber.Ctx) error {
	b := GetByIdRequest{}
	_ = c.BodyParser(&b)

	data := h.RoleUc.GetById(b.ID)
	if data.ID == "" {
		return core.NotFound(c, localizations.CommonNotFoundData)
	}

	return core.Ok(c, data)
}

func (h *handler) GetList(c *fiber.Ctx) error {
	payload := core.Payload(c)

	res := h.RoleUc.GetListByUnderRoles(payload.Roles)

	return core.Ok(c, res)
}

func (h *handler) Create(c *fiber.Ctx) error {
	b := CreateRole{}
	_ = c.BodyParser(&b)

	if err := h.RoleUc.Add(&b); err != nil {
		return core.BadRequest(c, err.Error())
	}

	return core.Created(c, b)
}

func (h *handler) Update(c *fiber.Ctx) error {
	b := UpdateRole{}
	_ = c.BodyParser(&b)

	if err := h.RoleUc.Update(&b); err != nil {
		return core.BadRequest(c, err.Error())
	}

	return core.Ok(c, b)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	b := GetByIdRequest{}
	_ = c.BodyParser(&b)

	if err := h.RoleUc.Delete(b.ID); err != nil {
		return core.BadRequest(c, localizations.CommonCannotDeletePleaseTryAgain)
	}

	return core.Ok(c, core.SuccessData(c, localizations.CommonDeleteSuccess))
}

func NewHandler(userUc user.UseCase, roleUc UseCase) Handler {
	return &handler{
		UserUc: userUc,
		RoleUc: roleUc,
	}
}