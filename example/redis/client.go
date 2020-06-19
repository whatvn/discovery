package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/whatvn/discovery"
	pb "github.com/whatvn/discovery/example/protobuf"
	"github.com/whatvn/discovery/redis"
	"google.golang.org/grpc"
)

func main() {
	registry := redis.NewResolver("127.0.0.1:6379", "", "hello.svc")
	conn, err := grpc.Dial(registry.SvcName(), discovery.DefaultBalancePolicy(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)
	response, err := client.SayHelloAnonymous(context.Background(), &empty.Empty{})
	fmt.Println(response, err)
}
