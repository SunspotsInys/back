package utils

import (
	"sync"
	"time"
)

const (
	epoch          = uint64(1626480000000)             // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	bitTimestamp   = uint(56)                          // 时间戳占用位数
	bitSequence    = uint(8)                           // 序列所占的位数
	maxTimestamp   = uint64(-1 ^ (-1 << bitTimestamp)) // 时间戳最大值
	maxSequence    = uint64(-1 ^ (-1 << bitSequence))  // 支持的最大序列id数量
	shiftWorkerid  = uint64(bitSequence)               // 机器id左移位数
	shiftTimestamp = uint64(bitSequence)               // 时间戳左移位数
)

type snowFlake struct {
	mutex     sync.Mutex
	year      int
	timestamp uint64
	sequence  uint64
}

func (s *snowFlake) GetVal() uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now()
	if now.Year() != s.year || s.sequence >= maxSequence {
		s.year = now.Year()
		s.timestamp = uint64(now.UnixNano()) / 1000000
		s.sequence = 0
	}
	s.sequence++
	return (((s.timestamp - epoch) << uint64(bitTimestamp)) + s.sequence)
}

var sf *snowFlake = nil
var once sync.Once

func GetSnowflakeInstance() *snowFlake {
	once.Do(func() {
		sf = &snowFlake{}
	})
	return sf
}
