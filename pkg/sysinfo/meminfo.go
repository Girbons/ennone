package sysinfo

import (
	"fmt"

	"github.com/Girbons/ennone/internal/util"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

type MemInfo struct {
	Info
	Total float64
}

func (m *MemInfo) LoadsInfo() error {
	mem, err := mem.VirtualMemory()

	if err != nil {
		log.Error(err)
		return err
	}

	m.Name = "RAM"
	m.Total = util.ByteToGigabyte(mem.Total)
	return nil
}

func (m *MemInfo) String() string {
	return fmt.Sprintf("%s: %.2f GB", m.Name, m.Total)
}
