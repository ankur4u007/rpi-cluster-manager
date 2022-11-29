package usecases

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ConfigureDietPi(deitpiConfig domain.DietpiConfig) {
	updateConfig("dietpi.txt", deitpiConfig.TxtConfig)
	updateConfig("dietpi-wifi.txt", deitpiConfig.WifiConfig)
}

func updateConfig(fileName string, config map[string]bool) error {
	filePath := fmt.Sprintf("/Volumes/%s/%s", domain.Config.Boot.Flash.DefaultVolumeName, fileName)
	err := utilities.Exists(filePath)
	if err == nil {
		input, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			for key := range config {
				configKey := strings.Split(key, "=")
				if strings.Contains(line, configKey[0]) && !strings.Contains(line, " ") {
					lines[i] = key
				}
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filePath, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Skipping configuration as file:%s does not exists\n", filePath)
	}
	return nil
}
