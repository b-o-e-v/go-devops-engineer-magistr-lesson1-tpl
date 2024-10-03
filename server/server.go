package server

import (
	"fmt"
	"strconv"
	"strings"
)

// LoadAverage: текущий Load Average сервера
// TotalMemory: текущий объём оперативной памяти сервера в байтах
// UsedMemory: текущее потребление оперативной памяти сервера в байтах
// TotalDiskSpace: текущий объём дискового пространства сервера в байтах
// UsedDiskSpace: текущее потребление дискового пространства сервера в байтах
// NetworkBandwidth: текущая пропускная способность сети в байтах в секунду
// NetworkUsage: текущая загруженность сети в байтах в секунду
type ServerStatus struct {
	LoadAverage      int
	TotalMemory      int
	UsedMemory       int
	TotalDiskSpace   int
	UsedDiskSpace    int
	NetworkBandwidth int
	NetworkUsage     int
}

func (s *ServerStatus) CheckLoadAverage() {
	if s.LoadAverage > 30 {
		fmt.Printf("Load Average is too high: %d\n", s.LoadAverage)
	}
}

func (s *ServerStatus) CheckMemoryUsage() {
	usagePercent := s.UsedMemory * 100 / s.TotalMemory
	if usagePercent > 80 {
		fmt.Printf("Memory usage too high: %d%%\n", usagePercent)
	}
}

func (s *ServerStatus) CheckDiskSpace() {
	freeSpace := s.TotalDiskSpace - s.UsedDiskSpace
	if s.UsedDiskSpace*100/s.TotalDiskSpace > 90 {
		fmt.Printf("Free disk space is too low: %d Mb left\n", freeSpace/1024/1024)
	}
}

func (s *ServerStatus) CheckNetworkUsage() {
	freeBandwidth := s.NetworkBandwidth - s.NetworkUsage
	if s.NetworkUsage*100/s.NetworkBandwidth > 90 {
		fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", freeBandwidth*8/1024/1024)
	}
}

func ParseStatus(data string) (ServerStatus, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 7 {
		return ServerStatus{}, fmt.Errorf("invalid data format")
	}

	values := make([]int, 7)
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return ServerStatus{}, fmt.Errorf("invalid number: %s", part)
		}
		values[i] = value
	}

	return ServerStatus{
		LoadAverage:      values[0],
		TotalMemory:      values[1],
		UsedMemory:       values[2],
		TotalDiskSpace:   values[3],
		UsedDiskSpace:    values[4],
		NetworkBandwidth: values[5],
		NetworkUsage:     values[6],
	}, nil
}
