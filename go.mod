module github.com/whatvn/discovery

go 1.13

require (
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/whatvn/denny v1.0.16
	go.etcd.io/etcd v3.3.22+incompatible
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
