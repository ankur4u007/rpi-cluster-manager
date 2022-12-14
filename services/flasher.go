package services

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/usecases"
)

func FlashDisk() {
	err := flash()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Flashing aborted!")
	}
}

func flash() error {
	if domain.Config.Boot.Flash.Enabled && usecases.ConfirmFlash() {
		err := usecases.ValidateConfigs()
		if err != nil {
			return err
		}
		err = usecases.UnmountDisk()
		if err != nil {
			return err
		}
		err = usecases.FlashAndTrackProgress()
		if err != nil {
			return err
		}
		err = usecases.RenameVolume()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Skipping flash...")
	}
	return nil
}
