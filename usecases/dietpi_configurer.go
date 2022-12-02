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
		multiKeyMaps := make(map[string][]string)
		for key := range config {
			configArrPair := strings.Split(key, "=")
			configKey := configArrPair[0]
			configValue := configArrPair[1]
			multiKeyMaps[configKey] = append(multiKeyMaps[configKey], configValue)
		}
		modifiedLines := []string{}
		existingConfigs := map[string]bool{}
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			lineKeyArr := strings.Split(line, "=")
			regularKey := lineKeyArr[0]
			commentedKey := strings.TrimPrefix(regularKey, "#")
			var key string
			var valuesArr []string
			if multiKeyMaps[regularKey] != nil {
				key = regularKey
				valuesArr = multiKeyMaps[regularKey]
			} else if multiKeyMaps[commentedKey] != nil {
				key = commentedKey
				valuesArr = multiKeyMaps[commentedKey]
			}
			if valuesArr != nil {
				for _, value := range valuesArr {
					newLine := fmt.Sprintf("%s=%s", key, value)
					if existingConfigs[newLine] == false {
						modifiedLines = append(modifiedLines, newLine)
						existingConfigs[newLine] = true
					}
				}
				continue
			}
			modifiedLines = append(modifiedLines, line)
			existingConfigs[line] = true
		}
		output := strings.Join(modifiedLines, "\n")
		err = ioutil.WriteFile(filePath, []byte(output), 0777)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Skipping configuration as file:%s does not exists\n", filePath)
	}
	return nil
}
