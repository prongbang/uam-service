package interceptor

type Interceptors struct {
	Auth AuthInterceptor
}

func New(
	authInterceptor AuthInterceptor,
) Interceptors {
	return Interceptors{
		Auth: authInterceptor,
	}
}
