package uam

import (
	"github.com/prongbang/uam-service/internal/uam/interceptor"
	"github.com/prongbang/uam-service/internal/uam/service/auth"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50052"

type Listeners interface {
	Serve()
}

type listeners struct {
	Interceptors interceptor.Interceptors
	AuthServer   auth.AuthServer
	RoleServer   role.RoleServer
	UserServer   user.UserServer
}

// Serve implements Listeners.
func (l *listeners) Serve() {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(l.Interceptors.Auth.Intercept),
		)
		auth.RegisterAuthServer(s, l.AuthServer)
		role.RegisterRoleServer(s, l.RoleServer)
		user.RegisterUserServer(s, l.UserServer)

		// Register reflection uam on gRPC server.
		reflection.Register(s)
		log.Printf("Server listening at %v", lis.Addr())
		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func NewListeners(
	interceptors interceptor.Interceptors,
	authServer auth.AuthServer,
	roleServer role.RoleServer,
	userServer user.UserServer,
) Listeners {
	return &listeners{
		Interceptors: interceptors,
		AuthServer:   authServer,
		RoleServer:   roleServer,
		UserServer:   userServer,
	}
}
