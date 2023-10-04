package uam

import "github.com/prongbang/user-service/internal/shared/user"

type APIHandler interface {
}

type apiHandler struct {
	Uc user.UseCase
}

func NewHandler(uc user.UseCase) APIHandler {
	return &apiHandler{
		Uc: uc,
	}
}
