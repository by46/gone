package main

import (
	"log"
	"net"

	"github.com/by46/gone/im"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *im.HelloRequest) (*im.HelloReply, error) {
	return &im.HelloReply{Message: in.Name,}, nil
}
func (s *server) SayHelloAgain(ctx context.Context, in *im.HelloRequest) (*im.HelloReply, error) {
	return &im.HelloReply{Message: "Hello again " + in.Name}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	im.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	s.Serve(l)
}
