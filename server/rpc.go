package server

import (
	"context"
	"log"

	"github.com/chheller/go-rpc-todo/modules/helloworld"
	"google.golang.org/grpc"
)


type RPCServer struct {
 Server *grpc.Server
}

type RPCServices struct {
}

type HelloWorldService struct {
	helloworld.UnimplementedGreeterServer
}
func (s *HelloWorldService) SayHello(_ context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (rpc RPCServer) Init() {
	helloworld.RegisterGreeterServer(rpc.Server, &HelloWorldService{})
}