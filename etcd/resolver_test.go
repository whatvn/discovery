package etcd

import (
	"github.com/whatvn/discovery"
	"reflect"
	"testing"
)

func TestNewResolver(t *testing.T) {
	var (
		name = "etcd.disc.svc"
		svcName = discovery.Prefix + ":///" + name
	)
	type args struct {
		etcdAddrs   string
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Init",
			args{"127.0.0.1:2379", name},
			svcName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResolver(tt.args.etcdAddrs, tt.args.serviceName); !reflect.DeepEqual(got.SvcName(), tt.want) {
				t.Errorf("NewResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

