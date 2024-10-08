package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metrics struct {
	CPUUsage          float64 `json:"cpu_usage"`
	LogicalCoreCount  int     `json:"cpu_logical_core_count"`
	PhysicalCoreCount int     `json:"cpu_physical_core_count"`
	MemUsedPercent    float64 `json:"mem_used_percent"`
	DiskUsedPercent   float64 `json:"disk_used_percent"`
	Timestamp         string  `json:"timestamp"`
}

func getMetrics(seconds_to_calc_cpu time.Duration) (*Metrics, error) {
	// Fetch CPU percentage
	percent, err := cpu.Percent(seconds_to_calc_cpu, false)
	if err != nil {
		return nil, err
	}

	// Fetch local core count
	cpuLogicalCoreCount, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	// Fetch physical core count
	cpuPhysicalCoreCount, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}

	// Fetch Memory usage
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// Fetch Disk usage
	disk, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	// Create Metrics struct
	metrics := &Metrics{
		CPUUsage:          percent[0],
		LogicalCoreCount:  cpuLogicalCoreCount,
		PhysicalCoreCount: cpuPhysicalCoreCount,
		MemUsedPercent:    math.Round(memory.UsedPercent),
		DiskUsedPercent:   math.Round(disk.UsedPercent),
		Timestamp:         time.Now().Format(time.RFC3339),
	}

	return metrics, nil
}

func main() {
	// Seconds to wait for calculating CPU usage
	seconds_to_calc_cpu := 3 * time.Second
	metrics, err := getMetrics(seconds_to_calc_cpu)
	if err != nil {
		log.Fatalf("Error getting metrics: %v", err)
	}

	fmt.Printf("Metrics collected: %+v\n", metrics)
}
