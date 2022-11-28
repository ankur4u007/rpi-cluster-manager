package commands

import (
	"bufio"
	"io"
	"os/exec"
	"time"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/rs/zerolog/log"
)

type CommandRunnerService struct{}

func (runner *CommandRunnerService) Run(command string, result chan<- domain.CommandResult) {
	defer close(result)
	log.Debug().
		Str("command", command).
		Msg("Running command")
	cmd := exec.Command("bash", "-c", command)
	pipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	err := cmd.Start()
	err = getCommandOutput(command, pipe, result, err)
	if err != nil {
		return
	}
	err = getCommandOutput(command, stderrPipe, result, err)
	if err != nil {
		return
	}
	cmd.Wait()
	log.Debug().
		Str("command", command).
		Msg("Command ran successfully")

	// waiting additional 5 seconds for all channles to close safely
	time.Sleep(5 * time.Second)
}

func getCommandOutput(command string, pipe io.ReadCloser, result chan<- domain.CommandResult, err error) error {
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Str("command", command).
			Msg("Error happened while running the command")
		result <- domain.CommandResult{
			Error: err,
			End:   true,
		}
		return err
	}
	go streamOutput(pipe, result)
	return nil
}

func streamOutput(pipe io.ReadCloser, result chan<- domain.CommandResult) {
	reader := bufio.NewReader(pipe)
	line, eof, err := reader.ReadLine()
	lineString := string(line)
	for err == nil && eof != true {
		lineString = string(line)
		result <- domain.CommandResult{
			Output: &lineString,
			End:    eof,
		}
		line, eof, err = reader.ReadLine()
	}
	if err == io.EOF || eof == true {
		result <- domain.CommandResult{
			End: true,
		}
	} else {
		result <- domain.CommandResult{
			Error: err,
			End:   true,
		}
	}
}
