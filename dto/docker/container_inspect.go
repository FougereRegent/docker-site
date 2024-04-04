package docker

import (
	"time"
)

type DockerContainerInspect struct {
	AppArmorProfile string   `json:"AppArmorProfile"`
	Args            []string `json:"Args"`
	Config          struct {
		AttachStderr bool     `json:"AttachStderr"`
		AttachStdin  bool     `json:"AttachStdin"`
		AttachStdout bool     `json:"AttachStdout"`
		Cmd          []string `json:"Cmd"`
		Domainname   string   `json:"Domainname"`
		Env          []string `json:"Env"`
		Healthcheck  struct {
			Test []string `json:"Test"`
		} `json:"Healthcheck"`
		Hostname string `json:"Hostname"`
		Image    string `json:"Image"`
		Labels   struct {
			ComExampleVendor  string `json:"com.example.vendor"`
			ComExampleLicense string `json:"com.example.license"`
			ComExampleVersion string `json:"com.example.version"`
		} `json:"Labels"`
		MacAddress      string `json:"MacAddress"`
		NetworkDisabled bool   `json:"NetworkDisabled"`
		OpenStdin       bool   `json:"OpenStdin"`
		StdinOnce       bool   `json:"StdinOnce"`
		Tty             bool   `json:"Tty"`
		User            string `json:"User"`
		Volumes         struct {
			VolumesData struct {
			} `json:"/volumes/data"`
		} `json:"Volumes"`
		WorkingDir  string `json:"WorkingDir"`
		StopSignal  string `json:"StopSignal"`
		StopTimeout int    `json:"StopTimeout"`
	} `json:"Config"`
	Created    time.Time `json:"Created"`
	Driver     string    `json:"Driver"`
	ExecIDs    []string  `json:"ExecIDs"`
	HostConfig struct {
		MaximumIOps       int `json:"MaximumIOps"`
		MaximumIOBps      int `json:"MaximumIOBps"`
		BlkioWeight       int `json:"BlkioWeight"`
		BlkioWeightDevice []struct {
		} `json:"BlkioWeightDevice"`
		BlkioDeviceReadBps []struct {
		} `json:"BlkioDeviceReadBps"`
		BlkioDeviceWriteBps []struct {
		} `json:"BlkioDeviceWriteBps"`
		BlkioDeviceReadIOps []struct {
		} `json:"BlkioDeviceReadIOps"`
		BlkioDeviceWriteIOps []struct {
		} `json:"BlkioDeviceWriteIOps"`
		ContainerIDFile    string        `json:"ContainerIDFile"`
		CpusetCpus         string        `json:"CpusetCpus"`
		CpusetMems         string        `json:"CpusetMems"`
		CPUPercent         int           `json:"CpuPercent"`
		CPUShares          int           `json:"CpuShares"`
		CPUPeriod          int           `json:"CpuPeriod"`
		CPURealtimePeriod  int           `json:"CpuRealtimePeriod"`
		CPURealtimeRuntime int           `json:"CpuRealtimeRuntime"`
		Devices            []interface{} `json:"Devices"`
		DeviceRequests     []struct {
			Driver       string     `json:"Driver"`
			Count        int        `json:"Count"`
			DeviceIDs    []string   `json:"DeviceIDs"`
			Capabilities [][]string `json:"Capabilities"`
			Options      struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Options"`
		} `json:"DeviceRequests"`
		IpcMode           string `json:"IpcMode"`
		Memory            int    `json:"Memory"`
		MemorySwap        int    `json:"MemorySwap"`
		MemoryReservation int    `json:"MemoryReservation"`
		OomKillDisable    bool   `json:"OomKillDisable"`
		OomScoreAdj       int    `json:"OomScoreAdj"`
		NetworkMode       string `json:"NetworkMode"`
		PidMode           string `json:"PidMode"`
		PortBindings      struct {
		} `json:"PortBindings"`
		Privileged      bool `json:"Privileged"`
		ReadonlyRootfs  bool `json:"ReadonlyRootfs"`
		PublishAllPorts bool `json:"PublishAllPorts"`
		RestartPolicy   struct {
			MaximumRetryCount int    `json:"MaximumRetryCount"`
			Name              string `json:"Name"`
		} `json:"RestartPolicy"`
		LogConfig struct {
			Type string `json:"Type"`
		} `json:"LogConfig"`
		Sysctls struct {
			NetIpv4IPForward string `json:"net.ipv4.ip_forward"`
		} `json:"Sysctls"`
		Ulimits []struct {
		} `json:"Ulimits"`
		VolumeDriver string `json:"VolumeDriver"`
		ShmSize      int    `json:"ShmSize"`
	} `json:"HostConfig"`
	HostnamePath    string `json:"HostnamePath"`
	HostsPath       string `json:"HostsPath"`
	LogPath         string `json:"LogPath"`
	ID              string `json:"Id"`
	Image           string `json:"Image"`
	MountLabel      string `json:"MountLabel"`
	Name            string `json:"Name"`
	NetworkSettings struct {
		Bridge                 string `json:"Bridge"`
		SandboxID              string `json:"SandboxID"`
		HairpinMode            bool   `json:"HairpinMode"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		SandboxKey             string `json:"SandboxKey"`
		EndpointID             string `json:"EndpointID"`
		Gateway                string `json:"Gateway"`
		GlobalIPv6Address      string `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int    `json:"GlobalIPv6PrefixLen"`
		IPAddress              string `json:"IPAddress"`
		IPPrefixLen            int    `json:"IPPrefixLen"`
		IPv6Gateway            string `json:"IPv6Gateway"`
		MacAddress             string `json:"MacAddress"`
		Networks               struct {
			Bridge struct {
				NetworkID           string `json:"NetworkID"`
				EndpointID          string `json:"EndpointID"`
				Gateway             string `json:"Gateway"`
				IPAddress           string `json:"IPAddress"`
				IPPrefixLen         int    `json:"IPPrefixLen"`
				IPv6Gateway         string `json:"IPv6Gateway"`
				GlobalIPv6Address   string `json:"GlobalIPv6Address"`
				GlobalIPv6PrefixLen int    `json:"GlobalIPv6PrefixLen"`
				MacAddress          string `json:"MacAddress"`
			} `json:"bridge"`
		} `json:"Networks"`
	} `json:"NetworkSettings"`
	Path           string `json:"Path"`
	ProcessLabel   string `json:"ProcessLabel"`
	ResolvConfPath string `json:"ResolvConfPath"`
	RestartCount   int    `json:"RestartCount"`
	State          struct {
		Error      string    `json:"Error"`
		ExitCode   int       `json:"ExitCode"`
		FinishedAt time.Time `json:"FinishedAt"`
		Health     struct {
			Status        string `json:"Status"`
			FailingStreak int    `json:"FailingStreak"`
			Log           []struct {
				Start    time.Time `json:"Start"`
				End      time.Time `json:"End"`
				ExitCode int       `json:"ExitCode"`
				Output   string    `json:"Output"`
			} `json:"Log"`
		} `json:"Health"`
		OOMKilled  bool      `json:"OOMKilled"`
		Dead       bool      `json:"Dead"`
		Paused     bool      `json:"Paused"`
		Pid        int       `json:"Pid"`
		Restarting bool      `json:"Restarting"`
		Running    bool      `json:"Running"`
		StartedAt  time.Time `json:"StartedAt"`
		Status     string    `json:"Status"`
	} `json:"State"`
	Mounts []struct {
		Name        string `json:"Name"`
		Source      string `json:"Source"`
		Destination string `json:"Destination"`
		Driver      string `json:"Driver"`
		Mode        string `json:"Mode"`
		RW          bool   `json:"RW"`
		Propagation string `json:"Propagation"`
	} `json:"Mounts"`
}
