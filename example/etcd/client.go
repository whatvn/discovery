package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/whatvn/discovery"
	"github.com/whatvn/discovery/etcd"
	pb "github.com/whatvn/discovery/example/protobuf"
	"google.golang.org/grpc"
)

func main() {
	registry := etcd.NewResolver("10.109.3.93:7379", "hello.svc")
	conn, err := grpc.Dial(registry.SvcName(), discovery.DefaultBalancePolicy(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)
	response, err := client.SayHelloAnonymous(context.Background(), &empty.Empty{})
	fmt.Println(response, err)
}
