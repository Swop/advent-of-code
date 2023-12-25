# Advent of Code - Golang

https://adventofcode.com/

> Advent of Code (credits: [Eric Wastl](http://was.tl/)) is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

- **Language**:
  - Golang (+ some Python for some days)
- **Years**:
  - 2023

## Requirements
- **Go** >= 1.21
- **Python** >= 3.8
- **graphviz** _(optional - for some days which required graph visualization)_

```shell
go mod download
pip install -r requirements.txt
```

## Run a specific solver

> [!WARNING]
> My input files are encrypted (cf [Run tests](#run-tests) section).
>
> Please replace the input file with your own input file, or decrypt mine with `git-crypt unlock`.

```bash
$ cd {YEAH}/{DAY} && cat {INPUT_FILE} | go run main.go -part={PART}
```
_Example for year 2023, day 1, part 1:_
```bash
$ cd 2023/01 && cat input.txt | go run main.go -part=1
```

## Run tests
This will run all solvers against sample data AND input data, and check expected results.
The command will also run utils package tests.

> [!WARNING]
> - The tests require all the input data to be present next to each solver, in a file named `input.txt`, e.g. `2023/day1/input.txt`
> - My input files are included in the repo, but they're `git-crypt` encrypted (to be compliant with AoC rules, which forbid to share publicly input files).
> - For test cases that target input data: the expected results are the ones matching my own input data. If you want to run the tests against your own input data, you'll need to update the expected results in the test cases to match yours (same if the inputs has been rotated).
**Warning:**


### All tests
```bash
$ make test
```

### Solvers only, for a specific year
```bash
$ go test ./2023/...
```

## Run all solvers (gives a pretty overview of the results & timings)
### Current year
```bash
$ make run-all
```

### Specific year
```bash
$ go run cmd/run_all.go -path="." -year=2023
```

## Run benchmarks

### All benchmarks
```bash
$ make bench
```

### Specific year
```bash
$ go test ./2023/... -bench=.
```
