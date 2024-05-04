package svc

import (
	"looklook/app/tiktok/cmd/rpc/internal/config"
	"looklook/app/tiktok/model"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	AsynqClient *asynq.Client

	TiktokModel model.DouyinLiveModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		AsynqClient: newAsynqClient(c),

		TiktokModel: model.NewDouyinLiveModel(sqlConn, c.Cache),
	}
}
