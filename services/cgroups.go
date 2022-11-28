package services

import (
	"fmt"
	"sync"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/usecases"
)

func ApplyCgroupsConfig(wg *sync.WaitGroup) {
	wg.Add(1)
	go doConfig(wg)
}

func doConfig(wg *sync.WaitGroup) {
	defer wg.Done()
	if domain.Config.Boot.Cgroups.Enabled {
		configText := fmt.Sprintf(" %s", domain.Config.Boot.Cgroups.ConfigText)
		filePath := fmt.Sprintf("/Volumes/%s/%s", domain.Config.Boot.Flash.DefaultVolumeName, domain.Config.Boot.Cgroups.ConfigFile)
		err := usecases.Exists(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = usecases.AppendFile(configText, filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Skipping setting cgroups as its disbaled")
	}
	return
}
