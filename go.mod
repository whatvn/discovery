module github.com/whatvn/discovery

go 1.13

require (
	github.com/coreos/etcd v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/whatvn/denny v1.0.3
	go.etcd.io/etcd v3.3.22+incompatible
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
)

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
)
