package usecases

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/interfaces"
)

func ConfirmFlash() bool {
	fmt.Printf("Flashing image:[%s] to disk:[%s]\n", domain.Config.Boot.Flash.ImagePath, domain.Config.Boot.Flash.DiskPath)
	fmt.Print("Type 'y' or 'Yes' to continue: ")
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	if strings.EqualFold("yes", string(text)) == true || strings.EqualFold("y", string(text)) {
		return true
	} else {
		return false
	}
}

func UnmountDisk() error {
	resultChannel := make(chan domain.CommandResult)
	command := fmt.Sprintf("diskutil unmountDisk %s", domain.Config.Boot.Flash.DiskPath)
	fmt.Printf("Unmounting disk..%s\n", domain.Config.Boot.Flash.DiskPath)
	go interfaces.CommandInstance.Run(command, resultChannel)
	errMsg := fmt.Sprintf("Failed to unmount disk:%s, Make sure to stop all disk operations and try again.\n", domain.Config.Boot.Flash.DiskPath)
	return handleChannelResponse(resultChannel, errMsg)
}

func FlashAndTrackProgress() error {
	imageSize, err := getImageSize()
	if err != nil {
		return err
	}
	ddCommandOutput := make(chan domain.CommandResult)
	command := fmt.Sprintf("dd bs=%s if=%s of=%s",
		domain.Config.Boot.Flash.WriteBs,
		domain.Config.Boot.Flash.ImagePath,
		domain.Config.Boot.Flash.DiskPath)
	go interfaces.CommandInstance.Run(command, ddCommandOutput)
	pkillInfoCommandOutput := make(chan domain.CommandResult)
	command = fmt.Sprintf("while pkill -INFO -x dd; do sleep %d; done", domain.Config.Boot.Flash.TrackIntervalInSeconds)
	go interfaces.CommandInstance.Run(command, pkillInfoCommandOutput)
	progressDetailsChannel := make(chan domain.ProgressDetails)
	go mapDDoutputToPogressBarSize(ddCommandOutput, progressDetailsChannel)
	interfaces.ProgressBarInstance.TrackProgressBar(
		"Flashing.",
		imageSize,
		progressDetailsChannel)
	fmt.Println("\nFinished flashing.")
	return nil
}

func RenameVolume() error {
	resultChannel := make(chan domain.CommandResult)
	command := fmt.Sprintf("diskutil mountDisk %s", domain.Config.Boot.Flash.DiskPath)
	fmt.Printf("Mounting disk..%s\n", domain.Config.Boot.Flash.DiskPath)
	go interfaces.CommandInstance.Run(command, resultChannel)
	errMsg := fmt.Sprintf("Failed to mount disk:%s, Make sure to stop all disk operations and try again.\n", domain.Config.Boot.Flash.DiskPath)
	handleChannelResponse(resultChannel, errMsg)
	resultChannel = make(chan domain.CommandResult)
	command = fmt.Sprintf("diskutil rename %ss1 %s", domain.Config.Boot.Flash.DiskPath, domain.Config.Boot.Flash.DefaultVolumeName)
	go interfaces.CommandInstance.Run(command, resultChannel)
	errMsg = fmt.Sprintf("Failed to rename volume on disk:%s, Make sure to stop all disk operations and try again.\n", domain.Config.Boot.Flash.DiskPath)
	return handleChannelResponse(resultChannel, errMsg)
}

func mapDDoutputToPogressBarSize(ddOutputChannel <-chan domain.CommandResult,
	progressDetailsChannel chan<- domain.ProgressDetails) {
	defer close(progressDetailsChannel)
	var ddResult domain.CommandResult
	pattern := regexp.MustCompile("[\\d]+ bytes")
	for ddResult.End != true && ddResult.Error == nil {
		ddResult = <-ddOutputChannel
		if ddResult.Output != nil {
			bytesTransferredStr := pattern.FindString(*ddResult.Output)
			if bytesTransferredStr != "" {
				bytesTransferred, _ := strconv.ParseInt((strings.Fields(bytesTransferredStr)[0]), 10, 64)
				progressDetailsChannel <- domain.ProgressDetails{
					Size: bytesTransferred,
				}
			}
		}
	}
	var ddResultErr domain.CommandResult
	for ddResult.Error == nil && ddResultErr.End != true && ddResultErr.Error == nil {
		ddResultErr = <-ddOutputChannel
	}
	progressDetailsChannel <- domain.ProgressDetails{End: true}
}

func getImageSize() (int64, error) {
	fi, err := os.Stat(domain.Config.Boot.Flash.ImagePath)
	if err != nil {
		return -1, err
	}
	// get the size
	return fi.Size(), nil
}

func handleChannelResponse(resultChannel <-chan domain.CommandResult, errMsg string) error {
	var sb strings.Builder
	var result domain.CommandResult
	for result.End != true && result.Error == nil {
		result = <-resultChannel
		if result.Output != nil {
			sb.WriteString(*result.Output)
		}
	}
	var errResult domain.CommandResult
	for result.Error == nil && errResult.End != true && errResult.Error == nil {
		errResult = <-resultChannel
		if errResult.Output != nil {
			sb.WriteString(*errResult.Output)
		}
	}
	var err error
	if result.Error != nil {
		err = result.Error
	} else {
		err = errResult.Error
	}
	if err != nil {
		fmt.Printf(errMsg)
		return err
	} else {
		fmt.Println(sb.String())
		return nil
	}
}
