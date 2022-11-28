package modules

import (
	"github.com/ankur4u007/dietpi-image-flasher/entities/interfaces"
	"github.com/ankur4u007/dietpi-image-flasher/infrastructure/commands"
	"github.com/ankur4u007/dietpi-image-flasher/infrastructure/configs"
	"github.com/ankur4u007/dietpi-image-flasher/infrastructure/logger"
	"github.com/ankur4u007/dietpi-image-flasher/infrastructure/progresstracker"
)

func Initialize() {
	logger.Initialize()
	interfaces.CommandInstance = &commands.CommandRunnerService{}
	interfaces.ProgressBarInstance = &progresstracker.ProgressBarService{}
	configs.LoadConfig()
}
