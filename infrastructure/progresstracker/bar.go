package progresstracker

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

type ProgressBarService struct{}

func (service *ProgressBarService) TrackProgressBar(message string, totalSizeInBytes int64, detailsChannel <-chan domain.ProgressDetails) {
	bar := progressbar.NewOptions(int(totalSizeInBytes),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan]%s.....[reset]", message)),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	var details domain.ProgressDetails
	previousValue := 0
	for details.End != true {
		details = <-detailsChannel
		bar.Add(int(details.Size) - previousValue)
		previousValue = int(details.Size)
	}
}
