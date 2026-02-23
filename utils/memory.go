package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	Total       float64
	Used        float64
	Free        float64
	UsedPercent float64
	SwapTotal   float64
	SwapUsed    float64
}

func GetMemoryInfo() (MemoryInfo, error) {

	file, _ := os.Open("/proc/meminfo")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var memTotal, memUsed, memFree, swapTotal, swapUsed float64

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			value, _ := strconv.ParseFloat(fields[1], 64)
			memTotal = value / 1024 / 1024
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			value, _ := strconv.ParseFloat(fields[1], 64)
			memFree = value / 1024 / 1024
		}

		if strings.HasPrefix(line, "SwapTotal:") {
			fields := strings.Fields(line)
			value, _ := strconv.ParseFloat(fields[1], 64)
			swapTotal = value / 1024 / 1024

			if strings.HasPrefix(line, "SwapFree:") {
				fields := strings.Fields(line)
				value, _ := strconv.ParseFloat(fields[1], 64)
				swapFree := value / 1024 / 1024
				swapUsed = swapTotal - swapFree

			}
		}
	}

	memUsed = memTotal - memFree
	memUsagePercent := (memUsed / memTotal) * 100

	return MemoryInfo{
		Total:       memTotal,
		Used:        memUsed,
		Free:        memFree,
		UsedPercent: memUsagePercent,
		SwapTotal:   swapTotal,
		SwapUsed:    swapUsed,
	}, nil
}
