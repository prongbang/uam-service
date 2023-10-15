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
	body := Password{}
	_ = c.BodyParser(&body)

	payload := core.Payload(c)
	body.UserID = payload.Sub

	if err := h.UserUc.UpdatePassword(&body); err != nil {
		return core.BadRequest(c, core.Translate(c, err.Error()))
	}
	return core.Ok(c, nil)
}

func (h *handler) GetByMe(c *fiber.Ctx) error {
	payload := core.Payload(c)

	data := h.UserUc.GetById(payload.Sub)
	if !core.IsUuid(data.ID) {
		return core.NotFound(c)
	}

	// Reset sensitive data
	data.Password = ""

	return core.Ok(c, data)
}

func (h *handler) GetById(c *fiber.Ctx) error {
	body := GetByIdRequest{}
	_ = c.BodyParser(&body)

	data := h.UserUc.GetById(body.ID)
	if !core.IsUuid(data.ID) {
		return core.NotFound(c)
	}

	// Reset sensitive data
	data.Password = ""

	return core.Ok(c, data)
}

func (h *handler) GetList(c *fiber.Ctx) error {
	paging := core.PagingBody(c)

	params := Params{}

	getCount := func() int64 { return h.UserUc.Count(params) }

	getData := func(limit int, offset int) any {
		params.LimitNo = paging.Limit
		params.OffsetNo = offset
		return h.UserUc.GetList(params)
	}

	return core.Ok(c, core.Pagination(paging.Page, paging.Limit, getCount, getData))
}

func (h *handler) Create(c *fiber.Ctx) error {
	b := CreateUser{}
	_ = c.BodyParser(&b)

	if err := h.UserUc.Add(&b); err != nil {
		return core.BadRequest(c, *err)
	}

	usr := h.UserUc.GetById(*b.ID)

	// Reset sensitive data
	usr.Password = ""

	return core.Created(c, usr)
}

func (h *handler) Update(c *fiber.Ctx) error {
	body := UpdateUser{}
	_ = c.BodyParser(&body)

	if usr := h.UserUc.GetById(body.ID); core.IsUuid(usr.ID) {
		if err := h.UserUc.Update(&body); err != nil {
			return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
		}
		usr = h.UserUc.GetById(body.ID)

		// Reset sensitive data
		usr.Password = ""

		return core.Ok(c, usr)
	}

	return core.NotFound(c)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	body := DeleteByIdRequest{}
	_ = c.BodyParser(&body)

	if usr := h.UserUc.GetById(body.ID); usr.ID != nil && *usr.ID == body.ID {
		if err := h.UserUc.Delete(body.ID); err != nil {
			return core.BadRequest(c, core.Translate(c, localizations.CommonInvalidData))
		}
		return core.Ok(c, nil)
	}
	return core.NotFound(c)
}

func NewHandler(userUc UseCase) Handler {
	return &handler{
		UserUc: userUc,
	}
}
