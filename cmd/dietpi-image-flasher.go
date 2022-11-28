package main

import (
	"sync"

	"github.com/ankur4u007/dietpi-image-flasher/infrastructure/modules"
	"github.com/ankur4u007/dietpi-image-flasher/services"
)

func main() {
	modules.Initialize()
	services.FlashDisk()
	var wg sync.WaitGroup
	services.ApplyCgroupsConfig(&wg)
	wg.Wait()
}
