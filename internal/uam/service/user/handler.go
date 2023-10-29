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
	body := Password{}
	_ = c.BodyParser(&body)

	if err := h.UserUc.UpdatePassword(&body); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}
	return core.Ok(c, nil)
}

func (h *handler) UpdatePasswordMe(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)
	body := Password{}
	_ = c.BodyParser(&body)

	body.UserID = payload.UserID

	if err := h.UserUc.UpdatePassword(&body); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}
	return core.Ok(c, nil)
}

func (h *handler) GetByMe(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)

	data := h.UserUc.GetById(ParamsGetById{ID: payload.UserID, Payload: payload})
	if !core.IsUuid(data.ID) {
		return core.NotFound(c)
	}

	// Reset sensitive data
	data.Password = ""

	return core.Ok(c, data)
}

func (h *handler) GetById(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)
	body := GetByIdRequest{}
	_ = c.BodyParser(&body)

	data := h.UserUc.GetById(ParamsGetById{ID: body.ID, Payload: payload})
	if !core.IsUuid(data.ID) {
		return core.NotFound(c)
	}

	// Reset sensitive data
	data.Password = ""

	return core.Ok(c, data)
}

func (h *handler) GetList(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)
	paging := core.PagingBody(c)

	params := Params{
		Payload: payload,
	}

	getCount := func() int64 { return h.UserUc.Count(params) }

	getData := func(limit int, offset int) []User {
		params.LimitNo = paging.Limit
		params.OffsetNo = offset
		return h.UserUc.GetList(params)
	}

	return core.Ok(c, core.Pagination[User](paging.Page, paging.Limit, getCount, getData))
}

func (h *handler) Create(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)
	body := CreateUser{}
	_ = c.BodyParser(&body)

	body.CreatedBy = payload.UserID

	usr, err := h.UserUc.Add(&body)
	if err != nil {
		return core.BadRequest(c, *err)
	}

	// Reset sensitive data
	usr.Password = ""

	return core.Created(c, usr)
}

func (h *handler) Update(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)
	body := UpdateUser{}
	_ = c.BodyParser(&body)

	if usr := h.UserUc.GetById(ParamsGetById{ID: body.ID, Payload: payload}); core.IsUuid(usr.ID) {
		us, err := h.UserUc.Update(&body)
		if err != nil {
			return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
		}

		// Reset sensitive data
		us.Password = ""

		return core.Ok(c, us)
	}

	return core.NotFound(c)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	payload := core.HttpPayload(c)
	body := DeleteByIdRequest{}
	_ = c.BodyParser(&body)

	data := DeleteUser{ID: body.ID, Payload: payload}
	if err := h.UserUc.Delete(data); err != nil {
		return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
	}
	return core.Ok(c, nil)
}

func NewHandler(userUc UseCase) Handler {
	return &handler{
		UserUc: userUc,
	}
}
