package qcache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

// GetData 从缓存中读取数据
func GetData(ctx context.Context, key string, callable func() (interface{}, error), timeInterval ...int) (data *gvar.Var, err error) {
	data, err = g.Redis().Do(ctx, "GET", key)

	if g.IsEmpty(data) {
		callData, err := callable()
		if err != nil {
			return data, err
		}

		var timeItv int
		if len(timeInterval) > 0 {
			timeItv = timeInterval[0]
		} else {
			timeItv = SnowsSlideRand()
		}

		_, err = g.Redis().Do(ctx, "SETEX", key, timeItv, callData)
		data = gvar.New(callData)
	}

	return data, err
}

// GetDataForceUpdate 读取数据并强制更新缓存
func GetDataForceUpdate(ctx context.Context, key string, callable func() (interface{}, error), timeInterval ...int) (data *gvar.Var, err error) {
	callData, err := callable()
	if err != nil {
		return data, err
	}

	var timeItv int
	if len(timeInterval) > 0 {
		timeItv = timeInterval[0]
	} else {
		timeItv = SnowsSlideRand()
	}

	_, err = g.Redis().Do(ctx, "SETEX", key, timeItv, callData)
	data = gvar.New(callData)

	return data, err
}
