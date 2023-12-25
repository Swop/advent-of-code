package main

import (
	"regexp"

	"github.com/Swop/advent-of-code/pkg/math"
	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	graph, directions, _ := parseInput(input)
	return computeSteps(graph, directions, []string{"AAA"})
}

func part2(input []string) any {
	graph, directions, start := parseInput(input)
	return computeSteps(graph, directions, start)
}

func computeSteps(graph map[string][2]string, dirs []uint8, n []string) int {
	lcmChan := make(chan int)
	for i := 0; i < len(n); i++ {
		go func(n string) {
			dIdx := 0
			steps := 0
			for {
				if n[len(n)-1] == 'Z' {
					lcmChan <- steps
					return
				}
				n = graph[n][dirs[dIdx]]
				dIdx = (dIdx + 1) % len(dirs)
				steps++
			}
		}(n[i])
	}
	var steps int
	for i := 0; i < len(n); i++ {
		if i == 0 {
			steps = <-lcmChan
			continue
		}
		steps = math.LCM(steps, <-lcmChan)
	}
	return steps
}

func parseInput(input []string) (map[string][2]string, []uint8, []string) {
	directions := make([]uint8, len(input[0]))
	for i, c := range input[0] {
		directions[i] = 1
		if c == 'L' {
			directions[i] = 0
		}
	}
	rxp := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)
	graph := map[string][2]string{}
	var start []string
	for _, line := range input[2:] {
		matches := rxp.FindStringSubmatch(line)
		graph[matches[1]] = [2]string{matches[2], matches[3]}
		if matches[1][len(matches[1])-1] == 'A' {
			start = append(start, matches[1])
		}
	}
	return graph, directions, start
}
