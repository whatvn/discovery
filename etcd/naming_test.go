package etcd

import (
	"github.com/whatvn/discovery"
	"testing"
)

var (
	name    = "etcd.disc.svc"
	svcName = discovery.Prefix + ":///" + name
)

func Test_etcd_Register(t *testing.T) {

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
			"Test Register",
			New("127.0.0.1:6379", svcName),
			args{
				"127.0.0.1:2367",
				5,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Register(tt.args.addr, tt.args.ttl); err != tt.wantErr {
				t.Errorf("etcd.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_etcd_UnRegister(t *testing.T) {
	var (
		addr = "127.0.0.1:2367"
	)
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
			"Test Un Register",
			New("127.0.0.1:2379", svcName),
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
			if err := tt.r.UnRegister(tt.args.addr); err != tt.wantErr {
				t.Errorf("etcd.UnRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
