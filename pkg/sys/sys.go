package sys

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/wailsapp/wails"
	"math"
	"runtime"
	"time"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
}
// CpuStats .
type CpuStats struct {
	Usage int
	Count int
	Os string
	Arch string
	Swap *mem.SwapMemoryStat
	Mem *mem.VirtualMemoryStat
	CPUInfo []cpu.InfoStat
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")
	return nil
}
//GetStats returns a pointer of type CpuStats .
func (s *Stats) GetStats() CpuStats  {
	return CpuStats{
		Usage: s.GetCPUUsage(),
		Count:   s.GetCPUCount(),
		Os:      runtime.GOOS,
		Arch:    runtime.GOARCH,
		Swap:    s.GetSwapMemory(),
		Mem:     s.GetMemory(),
		CPUInfo: s.GetCPUInfo(),
	}
}
func (s *Stats) GetCPUInfo() []cpu.InfoStat  {
	cpuInfo,err := cpu.Info()
	if err != nil {
		s.log.Errorf("Unable to retrive cpu information")
	}
	return cpuInfo
}
func (s *Stats) GetCPUCount() int {
	count,err := cpu.Counts(true)
	if err != nil {
		s.log.Errorf("unable to retrive cpu count!")
	}
	return count
}

// GetCPUUsage .
func (s *Stats) GetCPUUsage()int {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return 0
	}
	return  int(math.Round(percent[0]))
}

//GetSwapMemory .
func (s *Stats)  GetSwapMemory() *mem.SwapMemoryStat{
	sms,err := mem.SwapMemory()
	if err != nil {
		s.log.Errorf("Unable to retrive swap memory")
		return &mem.SwapMemoryStat{}
	}
	return sms
}
// GetMemory .
func (s *Stats) GetMemory() *mem.VirtualMemoryStat {
	ms, err := mem.VirtualMemory()
	if err != nil {
		s.log.Errorf("Unable to retrive memory!")
	}
	return ms
}