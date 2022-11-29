package domain

var Config AppConfiguration = AppConfiguration{
	Boot: BootConfiguration{
		Flash: FlashConfiguration{
			Enabled:                true,
			WriteBs:                "10m",
			TrackIntervalInSeconds: 1,
			DefaultVolumeName:      "DIETPI",
		},
		Cgroups: CgroupsConfiguration{
			Enabled:    true,
			ConfigText: "cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory",
			ConfigFile: "cmdline.txt",
		},
		DefaultDietPiConfigs: []string{
			"SURVEY_OPTED_IN=0",
			"AUTO_SETUP_AUTOMATED=1",
			"AUTO_SETUP_BROWSER_INDEX=0",
			"AUTO_SETUP_HEADLESS=1",
			"AUTO_SETUP_SWAPFILE_SIZE=0",
			"AUTO_SETUP_SWAPFILE_SIZE=1",
			"AUTO_SETUP_SWAPFILE_LOCATION=/var/swap",
			"AUTO_UNMASK_LOGIND=0"},
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
	DefaultDietPiConfigs []string
}

// FlashConfiguration exported
type FlashConfiguration struct {
	Enabled                bool
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
	Password string
}

// WifiConfiguration exported
type WifiConfiguration struct {
	Enabled  bool
	Name     string
	Password string
}

// K3sConfiguration exported
type K3sConfiguration struct {
	Enabled bool
}

// CopyFilesConfiguration exported
type CopyFilesConfiguration struct {
	Enabled bool
	Paths   []string
}

// FirstBootExecutablesConfiguration exported
type FirstBootExecutablesConfiguration struct {
	Enabled bool
	Paths   []string
}
