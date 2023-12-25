package main

import (
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/slices"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	cards := parseInput(input)
	total := 0
	for _, c := range cards {
		total += c.Worth()
	}
	return total
}

func part2(input []string) any {
	cards := parseInput(input)
	finalCardSet := map[int]int{}
	for _, c := range cards {
		if _, ok := finalCardSet[c.number]; !ok {
			finalCardSet[c.number] = 1
		} else {
			finalCardSet[c.number]++
		}
		times := finalCardSet[c.number]
		wins := len(c.Intersect())
		for i := 1; i <= wins; i++ {
			if _, ok := finalCardSet[c.number+i]; !ok {
				finalCardSet[c.number+i] = times
			} else {
				finalCardSet[c.number+i] += times
			}
		}
	}
	total := 0
	for _, count := range finalCardSet {
		total += count
	}
	return total
}

type card struct {
	number     int
	winNumbers []int
	myNumbers  []int
}

func (c *card) Intersect() []int {
	return slices.Intersect(c.winNumbers, c.myNumbers)
}

func (c *card) Worth() int {
	n := len(c.Intersect())
	if n == 0 {
		return 0
	}
	total := 1
	for i := 1; i < n; i++ {
		total *= 2
	}
	return total
}

func parseInput(input []string) []*card {
	var cards []*card
	for i, line := range input {
		parts := strings.Split(line, ":")
		numbers := strings.Split(parts[1], "|")
		winNumbers := strings.Split(numbers[0], " ")
		myNumbers := strings.Split(numbers[1], " ")
		c := &card{
			number: i + 1,
		}
		for _, num := range winNumbers {
			if num == "" {
				continue
			}
			num, _ := strconv.Atoi(num)
			c.winNumbers = append(c.winNumbers, num)
		}
		for _, num := range myNumbers {
			if num == "" {
				continue
			}
			num, _ := strconv.Atoi(num)
			c.myNumbers = append(c.myNumbers, num)
		}
		cards = append(cards, c)
	}
	return cards
}
