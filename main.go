package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func byteToGigabyte(value uint64) int64 {
	return int64(value / 1073741824)
}

func main() {
	var cpuModelName string

	a := app.New()

	infoStat, err := host.Info()
	if err != nil {
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
	}

	for _, c := range cpuInfo {
		cpuModelName = c.ModelName
	}

	vm, _ := mem.VirtualMemory()

	w := a.NewWindow("ennone")

	hostInfoLabel := widget.NewLabel(
		fmt.Sprintf("Operating System: %s %s", infoStat.OS, infoStat.PlatformVersion))

	cpuInfoLabel := widget.NewLabel(fmt.Sprintf("CPU: %s", cpuModelName))

	memoryInfoLabel := widget.NewLabel(fmt.Sprintf("RAM: %d GB", byteToGigabyte(vm.Total)))

	vbox := widget.NewVBox(
		hostInfoLabel,
		cpuInfoLabel,
		memoryInfoLabel,
	)

	w.SetContent(vbox)
	w.Resize(fyne.NewSize(400, 250))

	w.ShowAndRun()
}
