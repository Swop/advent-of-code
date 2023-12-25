package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	games := parseInput(input)
	total := 0
	for _, game := range games {
		if game.valid {
			total += game.number
		}
	}
	return total
}

func part2(input []string) any {
	games := parseInput(input)
	total := 0
	for _, game := range games {
		total += game.power
	}
	return total
}

type Color string

const (
	red   Color = "red"
	green Color = "green"
	blue  Color = "blue"
)

type game struct {
	number int
	valid  bool
	power  int
}

func parseInput(input []string) []game {
	var games []game
	r1 := regexp.MustCompile(`Game (\d+): (.*)`)
	r2 := regexp.MustCompile(`(\d+) (blue|red|green)`)
	maxs := map[Color]int{
		red:   12,
		green: 13,
		blue:  14,
	}
	for _, line := range input {
		mins := map[Color]int{
			red:   0,
			green: 0,
			blue:  0,
		}
		m := r1.FindAllStringSubmatch(line, -1)
		gameNumber, _ := strconv.Atoi(m[0][1])
		gameRounds := strings.Split(m[0][2], ";")
		validGame := true
		for _, part := range gameRounds {
			colorSet := strings.Split(part, ", ")
			for _, colorInfo := range colorSet {
				m := r2.FindStringSubmatch(colorInfo)
				maxColor := maxs[Color(m[2])]
				currentColor, _ := strconv.Atoi(m[1])
				if currentColor > maxColor {
					validGame = false
				}
				if currentColor > mins[Color(m[2])] {
					mins[Color(m[2])] = currentColor
				}
			}
		}
		games = append(games, game{
			number: gameNumber,
			valid:  validGame,
			power:  mins[red] * mins[green] * mins[blue],
		})
	}
	return games
}
