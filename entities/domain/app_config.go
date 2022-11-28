package domain

var Config AppConfiguration = AppConfiguration{
	Boot: BootConfiguration{
		Flash: FlashConfiguration{
			WriteBs:                "10m",
			TrackIntervalInSeconds: 1,
			DefaultVolumeName:      "DIETPI",
		},
		Cgroups: CgroupsConfiguration{
			Enabled:    true,
			ConfigText: "cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory",
			ConfigFile: "cmdline.txt",
		},
	},
}

// AppConfiguration exported
type AppConfiguration struct {
	Boot BootConfiguration
}

// BootConfiguration exported
type BootConfiguration struct {
	Flash                FlashConfiguration
	Cgroups              CgroupsConfiguration
	NodeDetails          NodeDetailsConfiguration
	Wifi                 WifiConfiguration
	K3s                  K3sConfiguration
	CopyFiles            CopyFilesConfiguration
	FirstBootExecutables FirstBootExecutablesConfiguration
}

// FlashConfiguration exported
type FlashConfiguration struct {
	ImagePath              string
	DiskPath               string
	WriteBs                string
	TrackIntervalInSeconds int
	DefaultVolumeName      string
}

// CgroupsConfiguration exported
type CgroupsConfiguration struct {
	Enabled    bool
	ConfigText string
	ConfigFile string
}

// NodeDetailsConfiguration exported
type NodeDetailsConfiguration struct {
	Enabled  bool
	Hostname string
	User     string
	Password string
}

// WifiConfiguration exported
type WifiConfiguration struct {
	Enable   bool
	Name     string
	Password string
}

// K3sConfiguration exported
type K3sConfiguration struct {
	Enable               bool
	DisableStartAsServer bool
}

// CopyFilesConfiguration exported
type CopyFilesConfiguration struct {
	Enable bool
	Paths  []string
}

// FirstBootExecutablesConfiguration exported
type FirstBootExecutablesConfiguration struct {
	Enable bool
	Paths  []string
}
