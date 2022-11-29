package usecases

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ConfigureWifi() domain.DietpiConfig {
	wifiConfig := make(map[string]bool)
	txtConfig := make(map[string]bool)
	if domain.Config.Boot.Wifi.Enabled {
		if utilities.IsValueNonEmpty(domain.Config.Boot.Wifi.Name) && utilities.IsValueNonEmpty(domain.Config.Boot.Wifi.Password) {
			nameConfig := fmt.Sprintf("aWIFI_SSID[0]='%s'", domain.Config.Boot.Wifi.Name)
			wifiConfig[nameConfig] = true
			passwordConfig := fmt.Sprintf("aWIFI_KEY[0]='%s'", domain.Config.Boot.Wifi.Password)
			wifiConfig[passwordConfig] = true
			txtConfig["AUTO_SETUP_NET_WIFI_ENABLED=1"] = true
		}
	} else {
		fmt.Println("Skipping wifi setup")
	}
	return domain.DietpiConfig{
		TxtConfig:  txtConfig,
		WifiConfig: wifiConfig,
	}
}
