package usecases

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ConfigureNode() map[string]bool {
	configs := make(map[string]bool)
	if domain.Config.Boot.Node.Enabled {
		if utilities.IsValueNonEmpty(domain.Config.Boot.Node.Hostname) {
			hostnameConfig := fmt.Sprintf("AUTO_SETUP_NET_HOSTNAME=%s", domain.Config.Boot.Node.Hostname)
			configs[hostnameConfig] = true
		}
		if utilities.IsValueNonEmpty(domain.Config.Boot.Node.Password) {
			passwordConfig := fmt.Sprintf("AUTO_SETUP_GLOBAL_PASSWORD=%s", domain.Config.Boot.Node.Password)
			configs[passwordConfig] = true
		}
	} else {
		fmt.Println("Skipping nodes setup")
	}
	return configs
}
