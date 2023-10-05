package role

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	id := c.Params("id")
	if id != "" {
		return core.Ok(c, h.RoleUc.GetById(id))
	}
	return core.BadRequest(c, "Required id")
}

func (h *handler) GetList(c *fiber.Ctx) error {
	payload := core.Payload(c)

	res := h.RoleUc.GetListByUnderRoles(payload.Roles)

	return core.Ok(c, res)
}

func (h *handler) Create(c *fiber.Ctx) error {
	b := Role{}
	if err := c.BodyParser(&b); err != nil {
		return core.BadRequest(c, "Bad Request")
	}

	v := validator.New()
	if err := v.Struct(b); err != nil {
		return core.BadRequest(c, err.Error())
	}

	if err := h.RoleUc.Add(&b); err != nil {
		return core.BadRequest(c, err.Error())
	}
	return core.Created(c, b)
}

func (h *handler) Update(c *fiber.Ctx) error {
	b := Role{}
	id := c.Params("id")
	bErr := c.BodyParser(&b)
	if bErr != nil || id == "" {
		return core.BadRequest(c, "Required id")
	}

	b.ID = id
	v := validator.New()
	if err := v.Struct(b); err != nil {
		return core.BadRequest(c, err.Error())
	}

	if err := h.RoleUc.Update(&b); err != nil {
		return core.BadRequest(c, err.Error())
	}
	return core.Ok(c, b)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return core.BadRequest(c, "Required id")
	}

	if err := h.RoleUc.Delete(id); err != nil {
		return core.BadRequest(c, "Can't delete")
	}
	return core.Ok(c, fiber.Map{"message": "Delete success"})
}

func NewHandler(userUc user.UseCase, roleUc UseCase) Handler {
	return &handler{
		UserUc: userUc,
		RoleUc: roleUc,
	}
}
