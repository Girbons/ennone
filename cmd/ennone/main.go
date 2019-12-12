package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/Girbons/ennone/pkg/sysinfo"
)

func main() {
	cpuInfo := new(sysinfo.CPUInfo)
	cpuInfo.LoadsInfo()

	memInfo := new(sysinfo.MemInfo)
	memInfo.LoadsInfo()

	hostInfo := new(sysinfo.HostInfo)
	hostInfo.LoadsInfo()

	a := app.New()
	hostInfoLabel := widget.NewLabel(hostInfo.String())
	cpuInfoLabel := widget.NewLabel(cpuInfo.String())
	memoryInfoLabel := widget.NewLabel(memInfo.String())

	vbox := widget.NewVBox(
		hostInfoLabel,
		cpuInfoLabel,
		memoryInfoLabel,
	)

	w := a.NewWindow("ennone")
	w.SetContent(vbox)
	w.Resize(fyne.NewSize(400, 250))
	w.ShowAndRun()
}
