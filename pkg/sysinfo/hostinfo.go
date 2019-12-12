package sysinfo

import (
	"fmt"

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

func (h *HostInfo) ReadTemperatures() error {
	temperatures, err := host.SensorsTemperatures()

	if err != nil {
		return err
	}

	fmt.Println(temperatures)

	//for _, temp := range temperatures {
	//}

	return nil
}

func (h *HostInfo) String() string {
	return fmt.Sprintf("Operating System: %s %s", h.OperativeSystem, h.PlatformVersion)
}
