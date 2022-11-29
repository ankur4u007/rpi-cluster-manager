package usecases

import (
	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ConfigureDefaults() map[string]bool {
	configs := make(map[string]bool)
	for _, entry := range domain.Config.Boot.DefaultDietPiConfigs {
		if utilities.IsValueNonEmpty(entry) {
			configs[entry] = true
		}
	}
	return configs
}
