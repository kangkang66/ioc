package pkg

import (
	"github.com/google/wire"
	"gomod/pkg/xmysql"
	"gomod/pkg/xredis"
)

var Provider = wire.NewSet(
	xredis.NewRdsClient,
	xmysql.NewDB,
)
