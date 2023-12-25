package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	math2 "github.com/Swop/advent-of-code/pkg/math"
	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return compute(parseInput(1, input))
}

func part2(input []string) any {
	return compute(parseInput(2, input))
}

func compute(moves []*Move) any {
	var vertices [][2]int
	current := [2]int{0, 0}
	perimeter := 0
	for _, m := range moves {
		vertices = append(vertices, current)
		perimeter += m.Steps
		current = [2]int{current[0] + m.Direction[0]*m.Steps, current[1] + m.Direction[1]*m.Steps}
	}
	// Using Pick's theorem
	return int(math2.Shoelace(vertices)+math.Floor(float64(perimeter)/2)) + 1
}

type Move struct {
	Direction [2]int
	Steps     int
}

func parseInput(part int, input []string) []*Move {
	directions := map[string][2]int{
		"R": {1, 0},
		"D": {0, 1},
		"L": {-1, 0},
		"U": {0, -1},
	}
	dirMappings := []string{"R", "D", "L", "U"}

	var moves []*Move
	for _, line := range input {
		f := strings.Fields(line)
		m := &Move{}
		if part == 1 {
			m.Direction = directions[f[0]]
			s, _ := strconv.Atoi(f[1])
			m.Steps = s
		} else {
			steps, err := strconv.ParseUint(f[2][2:len(f[2])-2], 16, 32)
			if err != nil {
				panic(fmt.Errorf("invalid input: %s", line))
			}
			dirIdx, _ := strconv.Atoi(string(f[2][len(f[2])-2]))
			m.Direction = directions[dirMappings[dirIdx]]
			m.Steps = int(steps)
		}
		moves = append(moves, m)
	}
	return moves
}
