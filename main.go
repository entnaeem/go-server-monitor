package main

import (
	"fmt"
	"go-server-monitor/utils"
	"os"
	"runtime"
)

func main() {

	OS := runtime.GOOS
	if OS == "windows" {
		fmt.Println("windows is not supported yet")
		os.Exit(1)
	}
	host, err := utils.GetHostinfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host.Hostname)
	fmt.Println(host.OSName)
	fmt.Println(host.OSArch)
	fmt.Println(host.KernelVersion)

}
