package fetch

import "github.com/orsinium-labs/enum"

type aocWebResource enum.Member[string]

var (
	awrDay   = aocWebResource{"day"}
	awrInput = aocWebResource{"input"}

	aocWebResources = enum.New(awrDay, awrInput)
)
