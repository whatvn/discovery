package redis

import (
	"reflect"
	"testing"

	"github.com/whatvn/discovery"
)

func TestNewResolver(t *testing.T) {
	var (
		name = "redis.disc.svc"
		svcName = discovery.Prefix + ":///" + name
	)
	type args struct {
		redisAddr   string
		redisPwd    string
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Init",
			args{"127.0.0.1:6379", "", "redis.disc.svc"},
			svcName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResolver(tt.args.redisAddr, tt.args.redisPwd, tt.args.serviceName); !reflect.DeepEqual(got.SvcName(), tt.want) {
				t.Errorf("NewResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}


