package services

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/usecases"
)

func UpdateDietpiConfig() {
	configs := make(map[string]bool)
	nodeConfig := usecases.ConfigureNode()
	for k, v := range nodeConfig {
		configs[k] = v
	}
	deitpiWifiConfig := usecases.ConfigureWifi()
	for k, v := range deitpiWifiConfig.TxtConfig {
		configs[k] = v
	}
	sshKeysConfig := usecases.ConfigureSshKeys()
	for k, v := range sshKeysConfig {
		configs[k] = v
	}
	defaultConfigs := usecases.ConfigureDefaults()
	for k, v := range defaultConfigs {
		configs[k] = v
	}
	usecases.ConfigureDietPi(domain.DietpiConfig{
		TxtConfig:  configs,
		WifiConfig: deitpiWifiConfig.WifiConfig,
	})
	fmt.Println("Updated dietpi configurations")
}
