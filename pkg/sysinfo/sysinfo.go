package sysinfo

type Info struct {
	Name        string
	ModelName   string
	Temperature float64
}

type HardwareInfo interface {
	LoadsInfo() error
	String() string
}
