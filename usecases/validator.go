package usecases

import (
	"fmt"
	"os"
	"strings"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
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

func Exists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("Path:%s does not exists", path)
	}
	return fmt.Errorf("Path:%s cannot be loaded", path)
}

func doImageValidation(imagePath string) error {
	err := Exists(imagePath)
	if err != nil {
		return err
	}
	if strings.Contains(imagePath, ".img") != true {
		return fmt.Errorf("Image:%s is not valid, must end with .img", imagePath)
	}
	return nil
}

func doDiskValidation(diskPath string) error {
	err := Exists(diskPath)
	if err != nil {
		return err
	}
	return nil
}
