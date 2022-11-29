package usecases

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
)

func ConfigureK3s() map[string]bool {
	config := make(map[string]bool)
	if domain.Config.Boot.K3s.Enabled {
		config["AUTO_SETUP_INSTALL_SOFTWARE_ID=193"] = true
	} else {
		fmt.Println("Skipping k3s setup")
	}
	return config
}
