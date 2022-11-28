package interfaces

import "github.com/ankur4u007/dietpi-image-flasher/entities/domain"

var ProgressBarInstance ProgressBar

type ProgressBar interface {
	TrackProgressBar(message string, totalSizeInBytes int64, progressDetailsChannel <-chan domain.ProgressDetails)
}
