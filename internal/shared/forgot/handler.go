package forgot

type Handler interface {
}

type handler struct {
	UserUc UseCase
}

func NewHandler(userUc UseCase) Handler {
	return &handler{
		UserUc: userUc,
	}
}
