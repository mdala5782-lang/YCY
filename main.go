// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	for {
// 		cpu := GetCPUUsage()
// 		mem := GetMemoryUsage()
// 		disk := GetDiskUsage("/")

// 		fmt.Printf(
// 			"CPU: %.2f%% | Memory: %.2f%%  | Disk: %.2f%%\n",
// 			cpu,
// 			mem,
// 			disk,
// 		)

// 		time.Sleep(1 * time.Second)
// 	}
// }
package main

func main() {
	go CollectMetrics()
	go StartHTTP()

	select {} // 阻塞主 goroutine
}
