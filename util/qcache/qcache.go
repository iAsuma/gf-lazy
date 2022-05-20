package qcache

import "github.com/gogf/gf/v2/util/grand"

// SnowsSlideRand 随机数，主要用于防止缓存雪崩，默认10-15分钟
func SnowsSlideRand(n ...int) int {
	if len(n) != 2 {
		return grand.N(600, 900)
	}

	return grand.N(n[0], n[1])
}
