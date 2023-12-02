package main

import (
	"context"
	"os"
	"os/signal"

	"go.jlucktay.dev/adventofcode/aocautoself/cmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	exitStatus := cmd.Execute(ctx, os.Stdout, os.Stderr, os.Args)

	stop()
	os.Exit(exitStatus)
}
