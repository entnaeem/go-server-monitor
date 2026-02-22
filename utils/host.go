package utils

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

type HostInfo struct {
	Hostname      string
	OSName        string
	OSArch        string
	KernelVersion string
}

func GetHostinfo() (HostInfo, error) {
	host, err := os.Hostname()
	if err != nil {
		return HostInfo{}, err
	}

	release, err := os.Open("/etc/os-release")
	if err != nil {
		return HostInfo{}, err
	}
	defer release.Close()

	var name, version string
	scanner := bufio.NewScanner(release)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "NAME=") {
			name = strings.Trim(line[5:], "\"")
		}
		if strings.HasPrefix(line, "VERSION=") {
			version = strings.Trim(line[8:], "\"")
		}
	}

	if err := scanner.Err(); err != nil {
		return HostInfo{}, err
	}
	osname := name + " " + version

	osarch := runtime.GOARCH

	kernelVersion, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return HostInfo{}, err
	}

	return HostInfo{
		Hostname:      host,
		OSName:        osname,
		OSArch:        osarch,
		KernelVersion: strings.TrimSpace(string(kernelVersion)),
	}, nil
}
