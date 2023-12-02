package cmd

import (
	"time"

	"github.com/spf13/pflag"
)

// Flags to pass parsed values forward to command logic.
var (
	flagYearOverride, flagDateOverride int
)

// localFlags returns a [pflag.FlagSet] to be added to the root command.
func localFlags() *pflag.FlagSet {
	pfs := &pflag.FlagSet{}
	now := time.Now()
	pfs.IntVarP(&flagYearOverride, "year", "y", now.Year(), "use this year instead of parsing from right now")
	pfs.IntVarP(&flagDateOverride, "date", "d", now.Day(), "use this day of the month instead of parsing from right now")

	return pfs
}
