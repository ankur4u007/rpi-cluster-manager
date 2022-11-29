package usecases

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ConfigureNode() map[string]bool {
	configs := make(map[string]bool)
	if domain.Config.Boot.NodeDetails.Enabled {
		if utilities.IsValueNonEmpty(domain.Config.Boot.NodeDetails.Hostname) {
			hostnameConfig := fmt.Sprintf("AUTO_SETUP_NET_HOSTNAME=%s", domain.Config.Boot.NodeDetails.Hostname)
			configs[hostnameConfig] = true
		}
		if utilities.IsValueNonEmpty(domain.Config.Boot.NodeDetails.Password) {
			passwordConfig := fmt.Sprintf("AUTO_SETUP_GLOBAL_PASSWORD=%s", domain.Config.Boot.NodeDetails.Password)
			configs[passwordConfig] = true
		}
	} else {
		fmt.Println("Skipping node details setup")
	}
	return configs
}
