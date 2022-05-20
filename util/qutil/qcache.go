package qutil

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

// SnowsSlideRand 随机数，主要用于防止缓存雪崩，默认10-15分钟
func SnowsSlideRand(n ...int) int {
	if len(n) != 2 {
		return grand.N(600, 900)
	}

	return grand.N(n[0], n[1])
}

// GetFromCache 从缓存中读取数据
func GetFromCache(ctx context.Context, key string, callable func() (interface{}, error), forceUpdate ...bool) (data *gvar.Var, err error) {
	redisCli := g.Redis()
	data, err = redisCli.Do(ctx, "GET", key)

	if g.IsEmpty(data) || (len(forceUpdate) > 0 && forceUpdate[0] == true) {
		callData, err := callable()
		if err != nil {
			return data, err
		}

		_, err = redisCli.Do(ctx, "SETEX", key, SnowsSlideRand(), callData)
		data = gvar.New(callData)
	}

	return data, err
}
