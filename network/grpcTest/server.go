package grpcTest

import (
	"context"
	pb "github.com/shibas11/go-hello-world/network/types"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	Port int
}

func NewServer(port int) *Server {
	s := new(Server)
	s.Port = port

	return s
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, s) // GRPC Server 구현체와 연결

	return grpcServer.Serve(listener)
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
