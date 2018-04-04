# TODO list

## Aim

- Create Go code templates, one for each day of the Advent of Code
- Start by aiming at 2017 only
- Eventually go back through previous years as well

## Misc

- Define `Day` and `DayDesc` structs, and fill them up with all of the things
- CLI arg(s)
  - year
  - date(s)

## Modular bits

- `HTTP GET` creates a `Day`
- Refine a `DayDesc` from the raw `Day` into sub-components:
  - intro
  - test case examples
  - stinger
- Arrange and execute text template, based on `Day` input
- Write template output to disk (implement the `Writer`(?) interface)
