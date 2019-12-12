package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
)

type CPUInfo struct {
	Info
}

func (c *CPUInfo) LoadsInfo() error {
	cpu, err := cpu.Info()

	if err != nil {
		log.Error(err)
		return err
	}

	c.Name = "CPU"
	c.ModelName = cpu[0].ModelName
	return nil
}

func (c *CPUInfo) String() string {
	return fmt.Sprintf("%s: %s", c.Name, c.ModelName)
}
