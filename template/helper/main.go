// Helper will copy code templates and change their contents (via the 'text/template' package) to match today's date.
// This only really works from inside my 'adventofcode' repo, and should be executed directly with 'go run'.
package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/lmittmann/tint"
)

type tmplData struct {
	Day, Year int
}

func main() {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				os.Stderr,
				&tint.Options{TimeFormat: time.RFC3339},
			)))

	gitTopCmd := exec.Command("git", "rev-parse", "--show-toplevel")

	gitTopOutput, err := gitTopCmd.CombinedOutput()
	if err != nil {
		slog.Error("running git command to find repo's top level",
			slog.String("command", gitTopCmd.String()),
			slog.Any("err", err),
			slog.String("output", string(gitTopOutput)))

		os.Exit(1)
	}

	gitTop := strings.TrimSpace(string(gitTopOutput))

	tmplDir := filepath.Join(gitTop, "template")

	tmplFiles := filepath.Join(tmplDir, "*.go.tmpl")

	tmpl, err := template.ParseGlob(tmplFiles)
	if err != nil {
		slog.Error("parsing template files",
			slog.Any("err", err),
			slog.String("files", tmplFiles))

		os.Exit(1)
	}

	now := time.Now()
	td := tmplData{
		Day:  now.Day(),
		Year: now.Year(),
	}

	for _, t := range tmpl.Templates() {
		dayLZ := fmt.Sprintf("%02d", now.Day())

		targetFilename := strings.TrimSuffix(t.Name(), ".tmpl")
		targetFilename = strings.ReplaceAll(targetFilename, "00", dayLZ)

		targetDirectory := filepath.Join(gitTop, strconv.Itoa(now.Year()), "day"+dayLZ)

		if strings.HasSuffix(targetFilename, "main.go") {
			targetDirectory = filepath.Join(targetDirectory, "cmd")
		}

		targetDirFile := filepath.Join(targetDirectory, targetFilename)

		if _, err := os.Stat(targetDirFile); err == nil {
			slog.Warn("target already exists, continuing with next template", slog.String("target", targetDirFile))
			continue
		}

		if err := os.MkdirAll(targetDirectory, 0o750); err != nil {
			slog.Error("making target directories",
				slog.Any("err", err),
				slog.String("target", targetDirectory))

			os.Exit(1)
		}

		slog.Info("target", slog.String("file", targetDirFile))

		f, err := os.Create(targetDirFile)
		if err != nil {
			slog.Error("creating target file",
				slog.Any("err", err),
				slog.String("file", targetDirFile))

			os.Exit(1)
		}
		defer f.Close()

		if err := tmpl.ExecuteTemplate(f, t.Name(), td); err != nil {
			slog.Error("executing template files",
				slog.Any("err", err),
				slog.String("files", tmplFiles))

			os.Exit(1)
		}
	}
}
