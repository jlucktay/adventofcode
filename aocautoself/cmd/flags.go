package cmd

import (
	"time"

	"github.com/spf13/pflag"
)

// Flags to pass parsed values forward to command logic.
var (
	flagYearOverride, flagDateOverride int
)

// rootFlags returns a [pflag.FlagSet] to be added to the root command.
func rootFlags() *pflag.FlagSet {
	pfs := &pflag.FlagSet{}
	now := time.Now()
	pfs.IntVarP(&flagYearOverride, "year", "y", now.Year(), "use given year instead of defaulting to this year")
	pfs.IntVarP(&flagDateOverride, "date", "d", now.Day(), "use given day of the month instead of defaulting to today")

	return pfs
}
