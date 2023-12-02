// This app will create some Advent of Code templates for you.
package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/spf13/cobra"

	"go.jlucktay.dev/adventofcode/aocautoself/pkg/fetch"
)

func Execute(ctx context.Context, stdout, stderr io.Writer, args []string) int {
	if len(args) < 1 {
		return ExitNoCommandName
	}

	cmdName := filepath.Base(args[0])

	version := fmt.Sprintf("%s built on %s from git SHA %s",
		versioninfo.Version, versioninfo.LastCommit.UTC().Format(time.RFC3339), versioninfo.Revision)

	if versioninfo.DirtyBuild {
		version += " (dirty)"
	}

	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use: cmdName,

		Short: "Create some Advent of Code templates.",
		Long: `Create some Advent of Code templates.
Looks up the adventofcode.com session cookie from the default profile of the local install of Firefox.`,

		Example: `  ` + cmdName + ` --year 2020
  ` + cmdName + ` -d=12`,

		Version: version,

		RunE: root(&flagYearOverride, &flagDateOverride),
	}

	// Add flags to the root command.
	rootCmd.Flags().AddFlagSet(rootFlags())

	// Wire in the arguments passed to this func.
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(args[1:])

	// Execute command with the given context.
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		switch {
		case errors.Is(err, ErrUnknownArguments):
			return ExitParsingArguments
		default:
			return ExitUnknown
		}
	}

	return ExitSuccess
}

func root(year, date *int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// Check for any remaining unparsed arguments.
		if len(args) > 0 {
			return fmt.Errorf("%w: '%s'", ErrUnknownArguments, strings.Join(args, "', '"))
		}

		// Display help text when passed no options.
		// Cf. https://clig.dev/#help
		if cmd.Flags().HasFlags() && cmd.Flags().NFlag() == 0 {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n%s", cmd.Long, cmd.UsageString())

			return nil
		}

		// Core logic for root command commences from here.
		cookie, err := fetch.FirefoxCookie()
		if err != nil {
			return err
		}

		day, err := fetch.Day(cmd.Context(), cookie, *year, *date)
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), day.String())

		return nil
	}
}
