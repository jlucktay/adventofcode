package cmd

// Exit status codes returned by the root command.
const (
	// ExitSuccess when everything goes to plan.
	ExitSuccess = iota

	// ExitUnknown if the cause of the error is not defined.
	ExitUnknown

	// ExitNoCommandName if Execute is passed a zero-length string slice, without a command name as the first element.
	ExitNoCommandName

	// ExitParsingArguments if parsing arguments goes awry.
	ExitParsingArguments
)
