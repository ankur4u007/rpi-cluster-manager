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
			"AUTO_SETUP_SSH_SERVER_INDEX=-2",
		},
		EjectWhenDone: true,
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
	Node                 NodeConfiguration
	Wifi                 WifiConfiguration
	SshKeys              SshKeysConfiguration
	DefaultDietPiConfigs []string
	EjectWhenDone        bool
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

// WifiConfiguration exported
type NodeConfiguration struct {
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

// SshKeysConfiguration exported
type SshKeysConfiguration struct {
	Enabled               bool
	PublicKeyPath         string
	DisablePasswordLogins bool
}
