package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/Girbons/ennone/pkg/sysinfo"
)

func monitTemperatures(cpuTempLabel, gpuTempLabel *widget.Label, h *sysinfo.HostInfo) {
	for {
		cpuTemp, gpuTemp, _ := h.ReadTemperatures()

		cpuTempLabel.SetText(fmt.Sprintf("CPU: %d", cpuTemp))
		gpuTempLabel.SetText(fmt.Sprintf("GPU: %d", gpuTemp))

		time.Sleep(time.Duration(10) * time.Second)
	}
}

func main() {
	cpuInfo := &sysinfo.CPUInfo{}
	cpuInfo.LoadsInfo()

	memInfo := &sysinfo.MemInfo{}
	memInfo.LoadsInfo()

	hostInfo := &sysinfo.HostInfo{}
	hostInfo.LoadsInfo()

	a := app.New()

	hostInfoLabel := widget.NewLabel(hostInfo.String())
	hostInfoLabel.Alignment = fyne.TextAlignCenter

	cpuInfoLabel := widget.NewLabel(cpuInfo.String())
	cpuInfoLabel.Alignment = fyne.TextAlignCenter

	memoryInfoLabel := widget.NewLabel(memInfo.String())
	memoryInfoLabel.Alignment = fyne.TextAlignCenter

	iGroup := widget.NewGroup("System Info")
	iGroup.Append(hostInfoLabel)
	iGroup.Append(cpuInfoLabel)
	iGroup.Append(memoryInfoLabel)

	tGroup := widget.NewGroup("Temperatures")
	cpuTemperatureLabel := widget.NewLabel(fmt.Sprintf("CPU: %d C", 30))
	cpuTemperatureLabel.Alignment = fyne.TextAlignCenter
	gpuTemperatureLabel := widget.NewLabel(fmt.Sprintf("GPU: %d C", 28))
	gpuTemperatureLabel.Alignment = fyne.TextAlignCenter

	tGroup.Append(cpuTemperatureLabel)
	tGroup.Append(gpuTemperatureLabel)

	vbox := widget.NewVBox(iGroup, tGroup)

	// starts temperature monitoring
	go monitTemperatures(
		cpuTemperatureLabel,
		gpuTemperatureLabel,
		hostInfo,
	)

	w := a.NewWindow("ennone")
	w.SetContent(vbox)
	w.Resize(fyne.NewSize(400, 250))
	w.ShowAndRun()
}
