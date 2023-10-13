package uam

import (
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50052"

type GRPCListener interface {
	Serve()
}

type gRPCListener struct {
	UamServer UamServer
}

// Serve implements GRPCListener.
func (l *gRPCListener) Serve() {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		RegisterUamServer(s, l.UamServer)

		// Register reflection service on gRPC server.
		reflection.Register(s)
		log.Printf("Server listening at %v", lis.Addr())
		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func NewListener(uamServer UamServer) GRPCListener {
	return &gRPCListener{
		UamServer: uamServer,
	}
}
