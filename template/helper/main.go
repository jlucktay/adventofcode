// Helper will copy code templates and change their contents (via the 'text/template' package) to match today's date.
// This only really works from inside my 'adventofcode' repo, and should be executed directly with 'go run'.
package main

import (
	"context"
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

	"go.jlucktay.dev/adventofcode/aocautoself/pkg/fetch"
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

	tmplFiles := filepath.Join(tmplDir, "*.tmpl")

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

	envDay := os.Getenv("AOC_DAY")
	if envDay != "" {
		convDay, err := strconv.Atoi(envDay)
		if err != nil {
			slog.Error("converting AOC_DAY environment variable",
				slog.Any("err", err),
				slog.String("raw", envDay))

			os.Exit(1)
		}

		td.Day = convDay
	}

	envYear := os.Getenv("AOC_YEAR")
	if envYear != "" {
		convYear, err := strconv.Atoi(envYear)
		if err != nil {
			slog.Error("converting AOC_YEAR environment variable",
				slog.Any("err", err),
				slog.String("raw", envYear))

			os.Exit(1)
		}

		td.Year = convYear
	}

	dayLZ := fmt.Sprintf("%02d", td.Day)
	targetDirectory := filepath.Join(gitTop, strconv.Itoa(td.Year), "day"+dayLZ)

	for _, t := range tmpl.Templates() {
		targetFilename := strings.TrimSuffix(t.Name(), ".tmpl")
		targetFilename = strings.ReplaceAll(targetFilename, "00", dayLZ)

		templateTargetDir := targetDirectory

		if strings.HasSuffix(targetFilename, "main.go") || strings.HasSuffix(targetFilename, "embed.go") {
			templateTargetDir = filepath.Join(templateTargetDir, "cmd")
		}

		targetDirFile := filepath.Join(templateTargetDir, targetFilename)

		if _, err := os.Stat(targetDirFile); err == nil {
			slog.Warn("target already exists, continuing with next template", slog.String("target", targetDirFile))
			continue
		}

		if err := os.MkdirAll(templateTargetDir, 0o750); err != nil {
			slog.Error("making target directories",
				slog.Any("err", err),
				slog.String("target", templateTargetDir))

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

	inputTxtFilePath := filepath.Join(targetDirectory, "cmd", "input.txt")

	if _, err := os.Stat(inputTxtFilePath); err == nil {
		slog.Warn("input file already exists, exiting", slog.String("target", inputTxtFilePath))
		return
	}

	ffxCookie, err := fetch.FirefoxCookie()
	if err != nil {
		slog.Error("getting cookie from Firefox",
			slog.Any("err", err))

		os.Exit(1)
	}

	slog.Info("fetching input for day",
		slog.Int("day", td.Day), slog.Int("year", td.Year))

	inputTxt, err := fetch.Input(context.TODO(), ffxCookie, td.Year, td.Day)
	if err != nil {
		slog.Error("getting input text from AOC",
			slog.Any("err", err))

		os.Exit(1)
	}

	slog.Info("writing today's input to file",
		slog.String("path", inputTxtFilePath),
		slog.Int("day", td.Day), slog.Int("year", td.Year))

	if err := os.WriteFile(inputTxtFilePath, []byte(inputTxt), 0o640); err != nil {
		slog.Error("writing input text to file",
			slog.Any("err", err))

		os.Exit(1)
	}

	slog.Info("input written to file OK",
		slog.String("path", inputTxtFilePath),
		slog.Int("day", td.Day), slog.Int("year", td.Year))
}
