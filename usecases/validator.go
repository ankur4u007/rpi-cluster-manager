package usecases

import (
	"fmt"
	"strings"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ValidateConfigs() error {
	imagePath := domain.Config.Boot.Flash.ImagePath
	diskPath := domain.Config.Boot.Flash.DiskPath
	err := doImageValidation(imagePath)
	if err != nil {
		return err
	}
	err = doDiskValidation(diskPath)
	if err != nil {
		return err
	}
	return nil
}

func doImageValidation(imagePath string) error {
	err := utilities.Exists(imagePath)
	if err != nil {
		return err
	}
	if strings.Contains(imagePath, ".img") != true {
		return fmt.Errorf("Image:%s is not valid, must end with .img", imagePath)
	}
	return nil
}

func doDiskValidation(diskPath string) error {
	err := utilities.Exists(diskPath)
	if err != nil {
		return err
	}
	return nil
}
