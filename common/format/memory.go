package format

import (
	"fmt"
	"runtime"
	"strconv"
)

type Memory struct {
	stats *runtime.MemStats
}

func (m Memory) String() string {
	return fmt.Sprintf("%s / %s", bytesToHumanReadable(m.stats.Alloc), bytesToHumanReadable(m.stats.TotalAlloc))
}

func bytesToHumanReadable(bytes uint64) string {
	suffix := "B"
	if bytes > 1024 {
		bytes /= 1024
		suffix = "KiB"
	}

	if bytes > 1024 {
		bytes /= 1024
		suffix = "MiB"
	}

	if bytes > 1024 {
		bytes /= 1024
		suffix = "GiB"
	}

	return strconv.FormatUint(bytes, 10) + " " + suffix
}



func GetUsedMemory() (runtime.MemStats, Memory) {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	return memstats, Memory { stats: &memstats }
}