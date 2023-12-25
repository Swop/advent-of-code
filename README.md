# Advent of Code - Golang

https://adventofcode.com/

> Advent of Code (credits: [Eric Wastl](http://was.tl/)) is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

- **Language**: Golang (+ some Python for some days)
- **Years**:
  - 2023

## Requirements
- Go >= 1.21
- Python >= 3.8, with additional packages:
  - z3-solver lib (for some days which required advanced problem solvers)
- graphviz (for some days which required graph visualization)

## Run tests
This will run all solvers against sample data AND input data, and check expected results.
The command will also run utils package tests.

**Warning:**
- The tests require all the input data to be present next to each solver, in a file named `input.txt`, e.g. `2023/day1/input.txt` (they're not included in the repo, as required by AoC rules).
- For test cases that target input data: the expected results are the ones matching my own input data. If you want to run the tests against your own input data, you'll need to update the expected results in the test cases to match yours (same if the inputs has been rotated).

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
