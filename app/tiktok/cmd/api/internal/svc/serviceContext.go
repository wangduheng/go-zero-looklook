package svc

import (
	"looklook/app/tiktok/cmd/api/internal/config"
	"looklook/app/tiktok/cmd/rpc/tiktok"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	TiktokRpc tiktok.Tiktok
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TiktokRpc: tiktok.NewTiktok(zrpc.MustNewClient(c.TiktokRpcConf)),
	}
}
