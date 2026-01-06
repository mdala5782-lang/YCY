package collector

import "syscall"

func GetDiskUsage(path string) float64 {
	var stat syscall.Statfs_t
	_ = syscall.Statfs(path, &stat)

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free

	return float64(used) / float64(total) * 100
}
