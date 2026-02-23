package utils

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type FullDiskUsage struct {
	Total   uint64
	Free    uint64
	Used    uint64
	UsedPct float64
}

func GetDiskUsage() (*FullDiskUsage, error) {

	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var total, free, used uint64

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		total += usage.Total
		free += usage.Free
		used += usage.Used
	}

	var usedPct float64
	if total > 0 {
		usedPct = (float64(used) / float64(total)) * 100
	}

	return &FullDiskUsage{
		Total:   total,
		Free:    free,
		Used:    used,
		UsedPct: usedPct,
	}, nil
}
