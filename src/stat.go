package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func getStat() (string, uint8) {
	cpuUsage, err := cpu.Percent(time.Second, false)
	checkErr(err)

	ramUsage, err := mem.VirtualMemory()
	checkErr(err)
	ramPercent := float32(ramUsage.Used) / float32(ramUsage.Total) * 100

	cpuImage := uint8(0)
	if cpuUsage[0] > 66.67 {
		cpuImage = 2
	} else if cpuUsage[0] > 33.33 {
		cpuImage = 1
	}

	return fmt.Sprintf("CPU: %.1f%% | RAM: %.1f%%", cpuUsage[0], ramPercent), cpuImage
}
