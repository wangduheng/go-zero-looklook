package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DouyinLiveModel = (*customDouyinLiveModel)(nil)

type (
	// DouyinLiveModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDouyinLiveModel.
	DouyinLiveModel interface {
		douyinLiveModel
	}

	customDouyinLiveModel struct {
		*defaultDouyinLiveModel
	}
)

// NewDouyinLiveModel returns a model for the database table.
func NewDouyinLiveModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DouyinLiveModel {
	return &customDouyinLiveModel{
		defaultDouyinLiveModel: newDouyinLiveModel(conn, c, opts...),
	}
}
