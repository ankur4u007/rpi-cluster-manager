package services

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/usecases"
)

func FlashDisk() error {
	if usecases.ConfirmFlash() == true {
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
