package middleware

type Middlewares struct {
	Auth AuthMiddleware
}

func New(
	authMiddleware AuthMiddleware,
) Middlewares {
	return Middlewares{
		Auth: authMiddleware,
	}
}
