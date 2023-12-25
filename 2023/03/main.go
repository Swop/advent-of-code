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
	numbers, symbols := parseInput(input)
	total := 0
loop:
	for _, num := range numbers {
		for i := num.start - 1; i <= num.end; i++ {
			if check(symbols, num.line-1, i) {
				total += num.val
				continue loop
			}
		}
		if check(symbols, num.line, num.start-1) {
			total += num.val
			continue loop
		}
		if check(symbols, num.line, num.end) {
			total += num.val
			continue loop
		}
		for i := num.start - 1; i <= num.end; i++ {
			if check(symbols, num.line+1, i) {
				total += num.val
				continue loop
			}
		}
	}
	return total
}

func part2(input []string) any {
	numbers, symbols := parseInput(input)
	total := 0
	for _, num := range numbers {
		num := num
		knownSymbols := map[string]struct{}{}
		for i := num.start - 1; i <= num.end; i++ {
			_, ok := knownSymbols[strconv.Itoa(num.line-1)+"-"+strconv.Itoa(i)]
			if check(symbols, num.line-1, i) && !ok {
				knownSymbols[strconv.Itoa(num.line-1)+"-"+strconv.Itoa(i)] = struct{}{}
				addAdjacent(symbols, num.line-1, i, &num)
			}
		}
		_, ok := knownSymbols[strconv.Itoa(num.line)+"-"+strconv.Itoa(num.start-1)]
		if check(symbols, num.line, num.start-1) && !ok {
			knownSymbols[strconv.Itoa(num.line)+"-"+strconv.Itoa(num.start-1)] = struct{}{}
			addAdjacent(symbols, num.line, num.start-1, &num)
		}
		_, ok = knownSymbols[strconv.Itoa(num.line)+"-"+strconv.Itoa(num.end)]
		if check(symbols, num.line, num.end) && !ok {
			knownSymbols[strconv.Itoa(num.line)+"-"+strconv.Itoa(num.end)] = struct{}{}
			addAdjacent(symbols, num.line, num.end, &num)
		}
		for i := num.start - 1; i <= num.end; i++ {
			_, ok := knownSymbols[strconv.Itoa(num.line+1)+"-"+strconv.Itoa(i)]
			if check(symbols, num.line+1, i) && !ok {
				knownSymbols[strconv.Itoa(num.line+1)+"-"+strconv.Itoa(i)] = struct{}{}
				addAdjacent(symbols, num.line+1, i, &num)
			}
		}
	}
	for _, symb := range symbols {
		if symb.val != "*" || len(symb.adg) != 2 {
			continue
		}
		ratio := 1
		for _, adg := range symb.adg {
			ratio *= adg.val
		}
		total += ratio
	}
	return total
}

func check(symbols map[string]*Symb, line int, pos int) bool {
	_, ok := symbols[strconv.Itoa(line)+"-"+strconv.Itoa(pos)]
	return ok
}

func addAdjacent(symbols map[string]*Symb, line int, pos int, num *Num) {
	symb := symbols[strconv.Itoa(line)+"-"+strconv.Itoa(pos)]
	symb.adg = append(symb.adg, num)
}

type Symb struct {
	pos  int
	line int
	val  string
	adg  []*Num
}

type Num struct {
	start int
	end   int
	line  int
	val   int
}

func parseInput(input []string) ([]Num, map[string]*Symb) {
	symbols := map[string]*Symb{}
	var numbers []Num
	rnum := regexp.MustCompile(`(\d+)`)
	rsymb := regexp.MustCompile(`([^\d.])`)
	for i, line := range input {
		line = strings.TrimSpace(line)
		for _, m := range rnum.FindAllStringSubmatchIndex(line, -1) {
			start, end := m[0], m[1]
			val, _ := strconv.Atoi(line[start:end])
			numbers = append(numbers, Num{
				start: start, end: end, line: i, val: val,
			})
		}
		for _, m := range rsymb.FindAllStringSubmatchIndex(line, -1) {
			start, end := m[0], m[1]
			val := line[start:end]
			symbols[strconv.Itoa(i)+"-"+strconv.Itoa(start)] = &Symb{start, i, val, nil}
		}
	}
	return numbers, symbols
}
