package main

import (
	"context"
	"golangapp/client/Protos/greet"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:7075")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greet.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	greet.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, req *greet.HelloRequest) (*greet.HelloReply, error) {
	return &greet.HelloReply{Message: "Hello " + req.GetName()}, nil
}
