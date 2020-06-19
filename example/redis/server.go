package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/whatvn/discovery/example/protobuf"
	"github.com/whatvn/discovery/redis"
	"google.golang.org/grpc"
	"net"
)

type Hello struct{}

func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	response := &pb.HelloResponse{
		Reply: "hi",
	}
	return response, nil
}


func (s *Hello) SayHelloAnonymous(ctx context.Context, in *empty.Empty) (*pb.HelloResponse, error) {

	response := &pb.HelloResponse{
		Reply: "hoho",
	}
	return response, nil
}
func main()  {
	addr := "127.0.0.1:8080"
	registry := redis.New("127.0.0.1:6379", "","hello.svc")
	if err := registry.Register(addr, 5); err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, new(Hello))
	server.Serve(listener)
}
