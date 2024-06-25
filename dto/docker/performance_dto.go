package docker

import "time"

type ContainerStats struct {
	Read      time.Time `json:"read"`
	PidsStats struct {
		Current int `json:"current"`
	} `json:"pids_stats"`
	Networks struct {
		Eth0 struct {
			RxBytes   int `json:"rx_bytes"`
			RxDropped int `json:"rx_dropped"`
			RxErrors  int `json:"rx_errors"`
			RxPackets int `json:"rx_packets"`
			TxBytes   int `json:"tx_bytes"`
			TxDropped int `json:"tx_dropped"`
			TxErrors  int `json:"tx_errors"`
			TxPackets int `json:"tx_packets"`
		} `json:"eth0"`
		Eth5 struct {
			RxBytes   int `json:"rx_bytes"`
			RxDropped int `json:"rx_dropped"`
			RxErrors  int `json:"rx_errors"`
			RxPackets int `json:"rx_packets"`
			TxBytes   int `json:"tx_bytes"`
			TxDropped int `json:"tx_dropped"`
			TxErrors  int `json:"tx_errors"`
			TxPackets int `json:"tx_packets"`
		} `json:"eth5"`
	} `json:"networks"`
	MemoryStats struct {
		Stats struct {
			TotalPgmajfault         int `json:"total_pgmajfault"`
			Cache                   int `json:"cache"`
			MappedFile              int `json:"mapped_file"`
			TotalInactiveFile       int `json:"total_inactive_file"`
			Pgpgout                 int `json:"pgpgout"`
			Rss                     int `json:"rss"`
			TotalMappedFile         int `json:"total_mapped_file"`
			Writeback               int `json:"writeback"`
			Unevictable             int `json:"unevictable"`
			Pgpgin                  int `json:"pgpgin"`
			TotalUnevictable        int `json:"total_unevictable"`
			Pgmajfault              int `json:"pgmajfault"`
			TotalRss                int `json:"total_rss"`
			TotalRssHuge            int `json:"total_rss_huge"`
			TotalWriteback          int `json:"total_writeback"`
			TotalInactiveAnon       int `json:"total_inactive_anon"`
			RssHuge                 int `json:"rss_huge"`
			HierarchicalMemoryLimit int `json:"hierarchical_memory_limit"`
			TotalPgfault            int `json:"total_pgfault"`
			TotalActiveFile         int `json:"total_active_file"`
			ActiveAnon              int `json:"active_anon"`
			TotalActiveAnon         int `json:"total_active_anon"`
			TotalPgpgout            int `json:"total_pgpgout"`
			TotalCache              int `json:"total_cache"`
			InactiveAnon            int `json:"inactive_anon"`
			ActiveFile              int `json:"active_file"`
			Pgfault                 int `json:"pgfault"`
			InactiveFile            int `json:"inactive_file"`
			TotalPgpgin             int `json:"total_pgpgin"`
		} `json:"stats"`
		MaxUsage int `json:"max_usage"`
		Usage    int `json:"usage"`
		Failcnt  int `json:"failcnt"`
		Limit    int `json:"limit"`
	} `json:"memory_stats"`
	BlkioStats struct {
	} `json:"blkio_stats"`
	CPUStats struct {
		CPUUsage struct {
			PercpuUsage       []int `json:"percpu_usage"`
			UsageInUsermode   int   `json:"usage_in_usermode"`
			TotalUsage        int   `json:"total_usage"`
			UsageInKernelmode int   `json:"usage_in_kernelmode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
		OnlineCpus     int   `json:"online_cpus"`
		ThrottlingData struct {
			Periods          int `json:"periods"`
			ThrottledPeriods int `json:"throttled_periods"`
			ThrottledTime    int `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"cpu_stats"`
	PrecpuStats struct {
		CPUUsage struct {
			PercpuUsage       []int `json:"percpu_usage"`
			UsageInUsermode   int   `json:"usage_in_usermode"`
			TotalUsage        int   `json:"total_usage"`
			UsageInKernelmode int   `json:"usage_in_kernelmode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
		OnlineCpus     int   `json:"online_cpus"`
		ThrottlingData struct {
			Periods          int `json:"periods"`
			ThrottledPeriods int `json:"throttled_periods"`
			ThrottledTime    int `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"precpu_stats"`
}

type PerformanceDTO struct {
	UsedMemory      int
	AvailableMemory int
	MemoryUsage     float32
	CpuUsage        float32
	CpuNumber       int
}

func (o *ContainerStats) ToPerformanceDTO() PerformanceDTO {
	usedMemory := o.MemoryStats.Usage - o.MemoryStats.Stats.Cache
	availableMemory := o.MemoryStats.Limit
	memoryUsage := float32(usedMemory) / float32(availableMemory) * 100.0
	cpuDelta := o.CPUStats.SystemCPUUsage - int64(o.PrecpuStats.CPUUsage.TotalUsage)
	systemCpuDelta := o.CPUStats.SystemCPUUsage - o.PrecpuStats.SystemCPUUsage
	numberCpus := o.CPUStats.OnlineCpus
	cpuUsage := float32((cpuDelta / systemCpuDelta)) * float32(numberCpus) * 100.0

	return PerformanceDTO{
		UsedMemory:      usedMemory,
		AvailableMemory: o.MemoryStats.Limit,
		MemoryUsage:     memoryUsage,
		CpuUsage:        cpuUsage,
		CpuNumber:       int(numberCpus),
	}
}
