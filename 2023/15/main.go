package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/hashmap"
	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	lenses := parseInput(input)
	tot := 0
	for _, l := range lenses {
		tot += hash(l.String())
	}
	return tot
}

func (l *lens) String() string {
	if l.focalLength != nil {
		return fmt.Sprintf("%s=%d", l.label, *l.focalLength)
	}
	return fmt.Sprintf("%s-", l.label)
}

func part2(input []string) any {
	lenses := parseInput(input)

	hCache := map[string]int{}
	hm := hashmap.New[*lens, int, *lens](func(l *lens) int {
		h, ok := hCache[l.label]
		if !ok {
			h = hash(l.label)
			hCache[l.label] = h
		}
		return h
	})
	for _, l := range lenses {
		if l.focalLength != nil {
			hm.SetWithReplace(l, func(bl *lens) (bool, bool) {
				return bl.label == l.label, true
			}, l)
			continue
		}
		hm.Unset(l, func(bl *lens) bool {
			return bl.label != l.label
		})
	}
	tot := 0
	for bIdx, b := range hm.Enumerate() {
		lIdx := 1
		b.TraverseForward(func(bl *lens) {
			tot += (1 + bIdx) * lIdx * *bl.focalLength
			lIdx++
		})
	}

	return tot
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}

type lens struct {
	label       string
	focalLength *int
}

func parseInput(input []string) []*lens {
	var lenses []*lens
	re := regexp.MustCompile(`([a-z]+)[-=](\d+)?`)
	seq := strings.Split(input[0], ",")
	for _, s := range seq {
		m := re.FindStringSubmatch(s)
		var fl *int
		flInt, _ := strconv.Atoi(m[2])
		if flInt > 0 {
			fl = &flInt
		}
		lenses = append(lenses, &lens{label: m[1], focalLength: fl})
	}
	return lenses
}
