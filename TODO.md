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
- each day has 2 parts
  - initially, only part 1 will be visible
  - an unauthenticated HTTP request won't ever see part 2
    - add functionality to allow a session cookie value to be specified
    - this would also allow us to fetch puzzle inputs, since they are user-specific

## Modular bits

- **HTTP GET** creates a `Day`
- Refine a `DayDesc` from the raw `Day` into sub-components:
  - intro
    - everything from the start up to the test case example(s)
  - test case example(s)
    - the first line that ends with a `:` up to the second-last line
  - stinger
    - the last line, which (usually?) ends with a `?`
- Arrange and execute text template, based on `Day` input
- Write template output to disk (implement the `Writer`(?) interface)
