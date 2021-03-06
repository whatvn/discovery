package redis

import (
	"github.com/whatvn/discovery"
	"testing"
)

var (
	svcName = "redis.disc.svc"
	addr    = "127.0.0.1:5389"
)

func Test_redis_Register(t *testing.T) {
	type args struct {
		addr string
		ttl  int
	}
	tests := []struct {
		name    string
		r       discovery.Registry
		args    args
		wantErr error
	}{
		{
			"Test register",
			New("127.0.0.1:6379", "", svcName),
			args{
				addr,
				5,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Register(tt.args.addr, tt.args.ttl); err != tt.wantErr {
				t.Errorf("redis.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redis_UnRegister(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		r       discovery.Registry
		args    args
		wantErr error
	}{
		{
			"Test unregister",
			New("127.0.0.1:6379", "", svcName),
			args{
				addr,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Register(addr, 5); err != nil {
				t.Errorf("redis.Register() error = %v, wantErr %v", err, nil)
			}
			if err := tt.r.UnRegister(tt.args.addr); err != nil {
				t.Errorf("redis.UnRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
