package collector

import (
	"os"
	"strconv"
	"strings"
)

func GetMemoryUsage() float64 {
	data, _ := os.ReadFile("/proc/meminfo")
	lines := strings.Split(string(data), "\n")

	var total, available uint64
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:" :
			total, _ = strconv.ParseUint(fields[1],10, 64)
		case "MemAvailable:" :
			available, _ = strconv.ParseUint(fields[1], 10, 64)
		}

	}

		used := total - available
		return float64(used) / float64(total) * 100
	
}