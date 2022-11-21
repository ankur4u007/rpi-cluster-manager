package service

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
)

func Boo() {
	fmt.Println("hellow from boo")
}

func Install(toolName string) bool {
	commandToExecute := fmt.Sprintf("apt-get -y install %s", toolName)
	_, err := exec.Command(commandToExecute).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Str("command", commandToExecute).
			Msg("Error happened while running command")
		return false
	}
	return true
}
