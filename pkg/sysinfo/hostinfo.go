package sysinfo

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/host"
	log "github.com/sirupsen/logrus"
)

type HostInfo struct {
	OperativeSystem string
	PlatformVersion string
}

func (h *HostInfo) LoadsInfo() error {
	info, err := host.Info()

	if err != nil {
		log.Error(err)
		return err
	}

	h.OperativeSystem = info.OS
	h.PlatformVersion = info.PlatformVersion
	return nil
}

func (h *HostInfo) ReadTemperatures() (int64, int64, error) {
	var (
		cpuAvgTemp int64
		gpuAvgTemp int64
	)

	results, err := host.SensorsTemperatures()
	if err != nil {
		return 0, 0, err
	}

	if runtime.GOOS != "windows" {
		log.Warning("%s not supported Yet", runtime.GOOS)
	}

	for _, r := range results {
		cpuAvgTemp += int64(r.Temperature)
	}

	cpuAvgTemp = cpuAvgTemp / int64(len(results))
	return cpuAvgTemp, gpuAvgTemp, errors.New("cannot retrieve temperatures")
}

func (h *HostInfo) String() string {
	return fmt.Sprintf("Operating System: %s %s", h.OperativeSystem, h.PlatformVersion)
}
