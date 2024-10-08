package main

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// Fetch CPU percentage
	seconds_to_calc_cpu := 3 * time.Second
	percent, _ := cpu.Percent(seconds_to_calc_cpu, false)
	fmt.Printf("CPU Percent: %v\n", percent)

	// Fetch local core count
	logical_cores, _ := cpu.Counts(true)
	fmt.Printf("CPU Cores (logical): %v\n", logical_cores)

	// Fetch physical core count
	physical_cores, _ := cpu.Counts(false)
	fmt.Printf("CPU Cores (physical): %v\n", physical_cores)

	// Fetch Memory usage
	v, _ := mem.VirtualMemory()
	fmt.Printf("Memory Usage: %v\n", v.UsedPercent)

	// Fetch Disk usage
	d, _ := disk.Usage("/")
	fmt.Printf("Disk Usage: %v\n", math.Round(d.UsedPercent))
}
