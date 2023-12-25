package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return computeWins(parseInput(1, input))
}

func part2(input []string) any {
	return computeWins(parseInput(2, input))
}

func computeWins(races [][2]float64) int {
	total := 1
	for _, race := range races {
		x1 := (race[0] - math.Sqrt(race[0]*race[0]-4*race[1])) / 2
		x2 := (race[0] + math.Sqrt(race[0]*race[0]-4*race[1])) / 2
		total *= int(math.Ceil(x2-1) - math.Floor(x1+1) + 1)
	}
	return total
}

func parseInput(part int, input []string) [][2]float64 {
	var races [][2]float64
	for i, line := range input {
		fields := strings.Fields(line)[1:]
		if part == 2 {
			fields = []string{strings.Join(fields, "")}
		}
		for j, number := range fields {
			if i == 0 {
				races = append(races, [2]float64{})
			}
			num, _ := strconv.Atoi(number)
			races[j][i] = float64(num)
		}
	}
	return races
}
