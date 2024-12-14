# Advent of Code

[![GHA workflow](https://github.com/jlucktay/adventofcode/actions/workflows/go.yaml/badge.svg)](https://github.com/jlucktay/adventofcode/actions/workflows/go.yaml)
[![Package reference](https://pkg.go.dev/badge/go.jlucktay.dev/adventofcode.svg)](https://pkg.go.dev/go.jlucktay.dev/adventofcode)

[![Coverage](https://github.com/jlucktay/adventofcode/raw/refs/heads/main/docs/coverage.svg)](https://github.com/jlucktay/adventofcode/blob/main/.octocov.yml)
[![Code to test ratio](https://github.com/jlucktay/adventofcode/raw/refs/heads/main/docs/ratio.svg)](https://github.com/jlucktay/adventofcode/blob/main/.octocov.yml)
[![Test execution time](https://github.com/jlucktay/adventofcode/raw/refs/heads/main/docs/time.svg)](https://github.com/jlucktay/adventofcode/blob/main/.octocov.yml)

## Events

- 2015
- 2016
- [2017](2017/)
- [2018](2018/)
- 2019
- [2020](2020/)
- 2021
- [2022](2022/)
- [2023](2023/)
- [2024](2024/)

## [Helper CLI](aocautoself/)

## Templates

Stored in [the `template` directory](template/) and rendered on-demand by `task go:template`.

The year and day that the `go:template` task renders for can be overridden by setting the `AOC_YEAR` and `AOC_DAY`
environment variables when calling the task, like so:

```shell
AOC_YEAR=2016 AOC_DAY=2 task go:template
```

## `crunchy` Library

This is where the reusable maths live.
