package user

type APIHandler interface {
}

type apiHandler struct {
	Uc UseCase
}

func NewHandler(uc UseCase) APIHandler {
	return &apiHandler{
		Uc: uc,
	}
}
