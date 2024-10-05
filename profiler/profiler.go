package profiler

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	loadThreshold    = 30
	memoryThreshold  = 80
	diskThreshold    = 90
	networkThreshold = 90
)

// LoadAverage: текущий Load Average сервера
// TotalMemory: текущий объём оперативной памяти сервера в байтах
// UsedMemory: текущее потребление оперативной памяти сервера в байтах
// TotalDiskSpace: текущий объём дискового пространства сервера в байтах
// UsedDiskSpace: текущее потребление дискового пространства сервера в байтах
// NetworkBandwidth: текущая пропускная способность сети в байтах в секунду
// NetworkUsage: текущая загруженность сети в байтах в секунду
type Profiler struct {
	LoadAverage      int
	TotalMemory      int
	UsedMemory       int
	TotalDiskSpace   int
	UsedDiskSpace    int
	NetworkBandwidth int
	NetworkUsage     int
}

func (s *Profiler) CheckLoadAverage() {
	if s.LoadAverage > loadThreshold {
		fmt.Printf("Load Average is too high: %d\n", s.LoadAverage)
	}
}

func (s *Profiler) CheckMemoryUsage() {
	usagePercent := s.UsedMemory * 100 / s.TotalMemory
	if usagePercent > memoryThreshold {
		fmt.Printf("Memory usage too high: %d%%\n", usagePercent)
	}
}

func (s *Profiler) CheckDiskSpace() {
	usagePercent := s.UsedDiskSpace * 100 / s.TotalDiskSpace
	freeSpace := (s.TotalDiskSpace - s.UsedDiskSpace) / 1024 / 1024
	if usagePercent > diskThreshold {
		fmt.Printf("Free disk space is too low: %d Mb left\n", freeSpace)
	}
}

func (s *Profiler) CheckNetworkUsage() {
	usagePercent := s.NetworkUsage * 100 / s.NetworkBandwidth
	freeBandwidth := (s.NetworkBandwidth - s.NetworkUsage) / 1000 / 1000
	if usagePercent > networkThreshold {
		fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", freeBandwidth)
	}
}

func Parse(data string) (Profiler, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 7 {
		return Profiler{}, fmt.Errorf("invalid data format")
	}

	values := make([]int, 7)
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return Profiler{}, fmt.Errorf("invalid number: %s", part)
		}
		values[i] = value
	}

	return Profiler{
		LoadAverage:      values[0],
		TotalMemory:      values[1],
		UsedMemory:       values[2],
		TotalDiskSpace:   values[3],
		UsedDiskSpace:    values[4],
		NetworkBandwidth: values[5],
		NetworkUsage:     values[6],
	}, nil
}
