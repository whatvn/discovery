# discovery

simple naming and resolver service for grpc, support etcd and redis.

# Usage

##1. Redis

### Server

```go

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

```

### Client

```go
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

```

##2. Etcd 

### Server

```go
package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/whatvn/discovery/etcd"
	"github.com/whatvn/discovery/example/protobuf"
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
	registry := etcd.New("10.109.3.93:7379", "hello.svc")
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

```

### Client 

```go
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

```


if you like this project, you may also like [denny](https://github.com/whatvn/denny), this is extracted from `denny`