package services

import (
	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/usecases"
)

func EjectWhenDone() error {
	if domain.Config.Boot.EjectWhenDone {
		err := usecases.EjectDisk()
		if err != nil {
			return err
		}
	}
	return nil
}
