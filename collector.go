package main

import (
	"time"
	"sys-monitor/collector"
)

func CollectMetrics() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		go func() {
			SetCPU(collector.GetCPUUsage())
		}()
		go func() {
			SetMem(collector.GetMemoryUsage())
		}()
		go func() {
			SetDisk(collector.GetDiskUsage("/"))
		}()
	}
}
