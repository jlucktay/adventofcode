# TODO list

## Aim

- Create Go code templates, one for each day of the Advent of Code
- Start by aiming at 2017 only
- Eventually go back through previous years as well
  - 2015
  - 2016

## Misc

- Define `Day` and `DayDesc` structs, and fill them up with all the things
- CLI argument(s)
  - year
  - date(s)
- each day has 2 parts
  - initially, only part 1 will be visible
  - an unauthenticated HTTP request won't ever see part 2
    - add functionality to allow a session cookie value to be specified
      - ~~see `cookiemonster` subdirectory for a one-shot script that can pull the appropriate cookie out~~
      - ✅ achieved with `browserutils/kooky` package
- grab every code block from the puzzle description page and make it into a table test input with `//go:embed`

## Modular bits

- **HTTP GET** fetches and returns a `Day`:
  - inputs:
    - date
    - year
    - optional session cookie
      - this would also allow us to fetch puzzle inputs, since they are user-specific
  - outputs:
    - Day
- Refine a `DayDesc` from the raw `Day` into subcomponents:
  - intro
    - everything from the start, up to the test case example(s)
  - test case example(s)
    - the first line that ends with a `:` up to the second-last line
  - stinger
    - the last line, which (usually?) ends with a `?` to summarise and pose the challenge
- Arrange and execute text template, based on `Day` input
- Write template output to disk
  - implement the `Writer`(?) interface
