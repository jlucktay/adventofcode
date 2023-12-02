package cmd

import "errors"

// Static errors returned by the root command.
var (
	ErrUnknownArguments = errors.New("unknown arguments passed in")
)
