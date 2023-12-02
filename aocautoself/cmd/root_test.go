package cmd_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/aocautoself/cmd"
)

func TestExecuteNoCommandName(t *testing.T) {
	is := is.New(t)

	actual := cmd.Execute(context.Background(), io.Discard, io.Discard, []string{})
	is.Equal(cmd.ExitNoCommandName, actual) // exit status should match when no command name is set
}

func TestExecuteCommandNameWithCanceledContext(t *testing.T) {
	is := is.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	const cmdName = "commandName"

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	actual := cmd.Execute(ctx, stdout, stderr, []string{cmdName})
	is.Equal(5, strings.Count(stdout.String(), cmdName))
	is.Equal(0, len(stderr.String()))
	is.Equal(cmd.ExitSuccess, actual)
}
