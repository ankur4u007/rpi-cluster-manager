package interfaces

import "github.com/ankur4u007/dietpi-image-flasher/entities/domain"

var CommandInstance CommandRunner

type CommandRunner interface {
	Run(ccommand string, result chan<- domain.CommandResult)
}
