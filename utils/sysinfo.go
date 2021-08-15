package utils

import (
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type sysInfo struct {
	Time string  `json:"time"`
	Num  int     `json:"num"`
	CPU  float64 `json:"cpu"`
	Mem  float64 `json:"mem"`
}

type sysInfos struct {
	SI       []sysInfo
	maxLen   int
	ticker   *time.Ticker
	interval time.Duration
	rw       sync.RWMutex
}

func (si *sysInfos) Start() {
	si.ticker = time.NewTicker(si.interval)
	go func() {
		for {
			<-si.ticker.C
			si.GetSysInfo()
		}
	}()
}

func (si *sysInfos) GetSysInfo() {
	si.rw.Lock()
	defer si.rw.Unlock()
	cpus, err := cpu.Percent(time.Duration(time.Second), false)
	if err != nil || len(cpus) == 0 {
		logger.Error().Msgf("Failed to get CPU usage, err = %v", err)
		return
	}
	mem, err := mem.VirtualMemory()
	if err != nil {
		logger.Error().Msgf("Failed to get memory usage, err = %v", err)
		return
	}
	si.SI = append(si.SI, sysInfo{
		Time: time.Now().Format("15:04:05"),
		Num:  runtime.NumGoroutine(),
		CPU:  cpus[0],
		Mem:  float64(mem.Used) / float64(mem.Total),
	})
	if len(si.SI) > si.maxLen {
		si.SI = si.SI[1:]
	}
}

var si *sysInfos = &sysInfos{
	SI:       make([]sysInfo, 0, 1444),
	maxLen:   1440,
	interval: time.Minute,
}

func GetSysInfos() *[]sysInfo {
	return &(si.SI)
}

func GetNewestSysInfo() *sysInfo {
	if len(si.SI) == 0 {
		return nil
	}
	return &(si.SI[len(si.SI)-1])
}

func init() {
	si.Start()
}
