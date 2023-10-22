package interceptor

type Interceptors struct {
	JWEInterceptor JWEInterceptor
}

func New(
	jweInterceptor JWEInterceptor,
) Interceptors {
	return Interceptors{
		JWEInterceptor: jweInterceptor,
	}
}
