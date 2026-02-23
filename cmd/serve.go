package cmd

import (
	"fmt"
	"go-server-monitor/utils"
)

func Serve() {

	host, err := utils.GetHostinfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("--------------Host Information:-----------------")
	fmt.Println("Hostname:", host.Hostname)
	fmt.Println("OS Name:", host.OSName)
	fmt.Println("OS Arch:", host.OSArch)
	fmt.Println("Kernel Version:", host.KernelVersion)

	cpuinfo, err := utils.GetCPUUsage()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n")
	fmt.Println("--------------CPU Information:-----------------")
	fmt.Println("CPU Name:", cpuinfo.ModelName)
	fmt.Println("CPU Cores:", cpuinfo.Cores)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuinfo.UsagePct)

	fmt.Println("\n")

	meminfo, err := utils.GetMemoryInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("--------------Memory Information:-----------------")
	fmt.Printf("Total Memory: %.2f GB\n", meminfo.Total)
	fmt.Printf("Used Memory: %.2f GB\n", meminfo.Used)
	fmt.Printf("Free Memory: %.2f GB\n", meminfo.Free)
	fmt.Printf("Memory Usage: %.2f%%\n", meminfo.UsedPercent)
	fmt.Printf("Swap Total: %.2f GB\n", meminfo.SwapTotal)
	fmt.Printf("Swap Used: %.2f GB\n", meminfo.SwapUsed)

	fmt.Println("\n")
	diskinfo, err := utils.GetDiskUsage()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("--------------Disk Information:-----------------")
	fmt.Printf("Total Disk Space: %.2f GB\n", float64(diskinfo.Total)/1024/1024/1024)
	fmt.Printf("Used Disk Space: %.2f GB\n", float64(diskinfo.Used)/1024/1024/1024)
	fmt.Printf("Free Disk Space: %.2f GB\n", float64(diskinfo.Free)/1024/1024/1024)
	fmt.Printf("Disk Usage: %.2f%%\n", diskinfo.UsedPct)
}
