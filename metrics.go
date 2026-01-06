package main

import (
	"math"
	"sync/atomic"
)

var cpuUsage uint64
var memUsage uint64
var diskUsage uint64

// ===== float64 <-> uint64 转换 =====

func floatToUint64(v float64) uint64 {
	return math.Float64bits(v)
}

func uint64ToFloat(v uint64) float64 {
	return math.Float64frombits(v)
}

// ===== Set =====

func SetCPU(v float64) {
	atomic.StoreUint64(&cpuUsage, floatToUint64(v))
}

func SetMem(v float64) {
	atomic.StoreUint64(&memUsage, floatToUint64(v))
}

func SetDisk(v float64) {
	atomic.StoreUint64(&diskUsage, floatToUint64(v))
}

// ===== Get =====

func GetCPU() float64 {
	return uint64ToFloat(atomic.LoadUint64(&cpuUsage))
}

func GetMem() float64 {
	return uint64ToFloat(atomic.LoadUint64(&memUsage))
}

func GetDisk() float64 {
	return uint64ToFloat(atomic.LoadUint64(&diskUsage))
}
