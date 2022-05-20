package qcalc

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"time"
)

func MakeTimeSequence(ctx context.Context, workId ...int8) string {
	s, _ := insTimeSequence.Make(ctx, workId...)
	return s
}

// MaxTimeInterval 最大时间间隔
const MaxTimeInterval = 86400000

// MaxThresholdNum 允许最大的阈值
const MaxThresholdNum = 99999999

// MaxWorkId 最大工作位ID
const MaxWorkId = 9

// MinRandNum 最小随机数
const MinRandNum = 0

// MaxRandNum 最大随机数
const MaxRandNum = 999

type timeSequence struct {
	ctx            context.Context // sequence上下文
	SequenceCache  *gcache.Cache   // 序列缓存对象
	startTimestamp int64           // 比较的开始时间
	lastTimestamp  int64           // 上次时间戳
	workId         int8            // 工作位/入口 ID，仅支持0-9的数
	loop           int8            //序列循环标记
}

var insTimeSequence = timeSequence{
	SequenceCache: gcache.New(),
}

func TimeSequence() *timeSequence {
	return &insTimeSequence
}

func (t *timeSequence) Make(ctx context.Context, workId ...int8) (sequence string, err error) {
	now := gtime.Now()
	nowTimestamp := now.TimestampMilli() //当前的毫秒时间戳
	preSequence := now.Format("ymd")     //序列前缀

	if nowTimestamp < t.lastTimestamp {
		return "", errors.New("服务器时间异常")
	}

	t.ctx = ctx
	t.setWorkId(workId...)

	if t.startTimestamp == 0 {
		t.setStartTimestamp(now)
	}

	timeInterval := nowTimestamp - t.startTimestamp
	sequence = t.sprintfSequence(preSequence, timeInterval)

	if t.inCurrentSequenceCache(sequence) {
		if t.loop > 6 {
			t.loop = 0 // 序列重复标记释放为0
			sequence = ""
		} else if t.loop > 5 {
			time.Sleep(1 * time.Millisecond) //阻塞到下一毫秒执行
			t.loop++
			sequence, err = t.Make(t.ctx)
		} else {
			timeInterval = t.getRandNum(MaxTimeInterval, MaxThresholdNum) //随机数字代替时间间隔
			sequence = t.sprintfSequence(preSequence, timeInterval)
			t.loop++
		}
	} else {
		t.loop = 0 // 序列重复标记释放为0
	}

	t.setLastSequenceCache(sequence, timeInterval, nowTimestamp)
	return sequence, err
}

// 18位字符串生成规则 （时间6位 + 工作位ID1位 + 唯一数8位 + 随机数3位）
func (t *timeSequence) sprintfSequence(prefix interface{}, unique interface{}) string {
	workId := t.workId
	rand := t.getRandNum(MinRandNum, MaxRandNum)
	return fmt.Sprintf("%06v%01v%08v%03v", prefix, workId, unique, rand)
}

func (t *timeSequence) setWorkId(workId ...int8) int8 {
	if len(workId) > 0 {
		t.workId = gconv.Int8(workId[0])
	} else {
		t.workId = 0
	}

	if t.workId > MaxWorkId {
		t.workId = 0 // 工作位ID不能大于9，受限于规则（18位订单号），暂时不能大于9。后期可修改。
	}

	return t.workId
}

// 获取当天的初始时间戳
func (t *timeSequence) setStartTimestamp(now *gtime.Time) int64 {
	start := now.StartOfDay()
	t.startTimestamp = start.TimestampMilli()
	return t.startTimestamp
}

// 获取随机数
func (t *timeSequence) getRandNum(min int, max int) int64 {
	return int64(grand.N(min, max))
}

// 设置当天序列缓存
func (t *timeSequence) setLastSequenceCache(key string, timeInterval int64, nowTimestamp int64) {
	duration := 2 * time.Second
	if timeInterval >= MaxTimeInterval {
		duration = time.Duration(MaxTimeInterval-nowTimestamp) * time.Millisecond
	}

	err := t.SequenceCache.Set(t.ctx, key, 1, duration)
	if err != nil {
		return
	}
}

// 判断是否在当天序列内
func (t *timeSequence) inCurrentSequenceCache(key string) bool {
	exist, err := t.SequenceCache.Contains(t.ctx, key)
	if err != nil {
		return false
	}

	return exist
}
