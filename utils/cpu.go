package utils

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUStatus struct {
	ModelName string
	Cores     int
	UsagePct  float64
}

func GetCPUUsage() (*CPUStatus, error) {
	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	model := "Unknown"
	cores := 0
	if len(info) > 0 {
		model = info[0].ModelName
		cores = len(info)
	}

	percentages, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		return nil, err
	}

	var usage float64
	if len(percentages) > 0 {
		usage = percentages[0]
	}

	return &CPUStatus{
		ModelName: model,
		Cores:     cores,
		UsagePct:  usage,
	}, nil
}
