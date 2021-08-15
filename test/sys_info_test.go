package test

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func TestCPU(t *testing.T) {
	info, _ := cpu.Percent(time.Duration(time.Second), false)
	t.Log(info)
	t.Log(mem.VirtualMemory())
}
