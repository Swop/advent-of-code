package main

import (
	"fmt"

	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	parseInput(input)
	return compute()
}

func part2(input []string) any {
	parseInput(input)
	return compute()
}

func compute() any {
	return 0
}

func parseInput(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
}
