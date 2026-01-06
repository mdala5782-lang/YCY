package collector

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type CpuStat struct {
	idle  uint64
	total uint64
}

func ReadCpuStat() CpuStat {
	data, _ := os.ReadFile("/proc/stat") //读取cpu信息到data变量中
	// 按行分割并取第一行，然后按空格分割成字段
	fields := strings.Fields(strings.Split(string(data), "\n")[0])
	// 计算总的CPU时间
	var total uint64
	for i := 1; i < len(fields); i++ {
		val, _ := strconv.ParseUint(fields[i], 10, 64)
		total += val
	}

	idle, _ := strconv.ParseUint(fields[4], 10, 64)
	return CpuStat{idle: idle, total: total}
}

func GetCPUUsage() float64 {
	stat1 := ReadCpuStat()
	time.Sleep(500 * time.Millisecond)
	stat2 := ReadCpuStat()

	idle := stat2.idle - stat1.idle
	total := stat2.total - stat1.total

	return 100 * (1 - float64(idle)/float64(total))
}
