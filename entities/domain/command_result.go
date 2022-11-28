package domain

type CommandResult struct {
	Output *string
	Error  error
	End    bool
}

type ProgressDetails struct {
	Size int64
	End  bool
}
