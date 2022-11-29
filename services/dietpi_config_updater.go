package services

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/usecases"
)

func UpdateDietpiConfig() {
	configs := make(map[string]bool)
	nodeDetails := usecases.ConfigureNode()
	for k, v := range nodeDetails {
		configs[k] = v
	}
	wifiDietPiConfig := usecases.ConfigureWifi()
	for k, v := range wifiDietPiConfig.TxtConfig {
		configs[k] = v
	}
	k3Configs := usecases.ConfigureK3s()
	for k, v := range k3Configs {
		configs[k] = v
	}
	defaultConfigs := usecases.ConfigureDefaults()
	for k, v := range defaultConfigs {
		configs[k] = v
	}
	usecases.ConfigureDietPi(domain.DietpiConfig{
		TxtConfig:  configs,
		WifiConfig: wifiDietPiConfig.WifiConfig,
	})
	fmt.Println("Updated dietpi configurations")
}
