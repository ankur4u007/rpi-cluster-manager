package main

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/infrastructure/modules"
	"github.com/ankur4u007/dietpi-image-flasher/services"
)

func main() {
	modules.Initialize()
	err := services.FlashDisk()
	if err == nil {
		fmt.Println("Completed!")
	} else {
		fmt.Println(err)
		fmt.Println("Aborted!")
	}

}
