package fetch

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"

	"go.jlucktay.dev/adventofcode/aocautoself/pkg/model"
)

// Day will perform the necessary HTTP magic to get an Advent of Code day for the given year and date.
func Day(ctx context.Context, session *http.Cookie, year, day int) (*model.Day, error) {
	details, err := getFromWeb(ctx, year, day, session)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(details))
	if err != nil {
		return nil, err
	}

	output := model.NewDay(year, day)

	g := errgroup.Group{}
	g.SetLimit(1)

	doc.Find("article.day-desc").Each(
		func(i int, s *goquery.Selection) {
			g.Go(func() error {
				switch i {
				case 0:
					breakDescriptionDown(&output.Part1, s.Text())
				case 1:
					breakDescriptionDown(&output.Part2, s.Text())
				default:
					return fmt.Errorf("'%d' article.day-desc elements is too many", i+1)
				}

				output.Description += s.Text()

				return nil
			})
		},
	)

	if err := g.Wait(); err != nil {
		return nil, fmt.Errorf("parsing doc: %w", err)
	}

	return output, nil
}

func breakDescriptionDown(dd *model.DayDesc, desc string) {
	descSlice := strings.Split(desc, "\n")
	endsWithColon, endsWithQuestion := false, false

	for _, line := range descSlice {
		if len(line) > 0 {
			switch strings.TrimSpace(line)[len(line)-1:] {
			case ":":
				endsWithColon = true
			case "?":
				endsWithQuestion = true
			}
		}

		if !endsWithColon && !endsWithQuestion {
			dd.Fluff += line + "\n"
		} else if !endsWithQuestion {
			dd.Test += line + "\n"
		} else {
			dd.Stinger += line + "\n"
		}
	}

	dd.Fluff = strings.TrimSpace(dd.Fluff)
	dd.Test = strings.TrimSpace(dd.Test)
	dd.Stinger = strings.TrimSpace(dd.Stinger)
}
