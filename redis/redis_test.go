package redis

import (
	"github.com/whatvn/discovery"
	"testing"
)

func TestNew(t *testing.T) {
	var (
		name = "redis.disc.svc"
		svcName = discovery.Prefix + ":///" + name
	)
	type args struct {
		redisAddr     string
		redisPassword string
		serviceName   string
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
			if got := New(tt.args.redisAddr, tt.args.redisPassword, tt.args.serviceName); got.SvcName() != svcName {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
