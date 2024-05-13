package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LiveResultModel = (*customLiveResultModel)(nil)

type (
	// LiveResultModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLiveResultModel.
	LiveResultModel interface {
		liveResultModel
	}

	customLiveResultModel struct {
		*defaultLiveResultModel
	}
)

// NewLiveResultModel returns a model for the database table.
func NewLiveResultModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LiveResultModel {
	return &customLiveResultModel{
		defaultLiveResultModel: newLiveResultModel(conn, c, opts...),
	}
}
