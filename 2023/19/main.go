package main

import (
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	wfs, parts := parseInput(input)
	total := 0
	for _, p := range parts {
		ranges := [4][2]int{}
		for i := 0; i < 4; i++ {
			ranges[i] = [2]int{p[i], p[i] + 1}
		}
		if 1 == computeCombinations(wfs, "in", ranges) {
			total += p[0] + p[1] + p[2] + p[3]
		}
	}
	return total
}

func part2(input []string) any {
	wfs, _ := parseInput(input)
	ranges := [4][2]int{{1, 4001}, {1, 4001}, {1, 4001}, {1, 4001}}
	return computeCombinations(wfs, "in", ranges)
}

func computeCombinations(wfs map[string][]string, wfName string, ranges [4][2]int) int {
	switch wfName {
	case "R":
		return 0
	case "A":
		c := 1
		for i := 0; i < 4; i++ {
			c *= ranges[i][1] - ranges[i][0]
		}
		return c
	}
	currentRanges := ranges
	wf := wfs[wfName]
	total := 0
	for _, r := range wf {
		m := strings.Split(r, ":")
		if len(m) == 1 {
			total += computeCombinations(wfs, r, currentRanges)
			break
		}
		catIdx := ratingCatNameToIdx(rune(m[0][0]))
		op := m[0][1]
		compVal, _ := strconv.Atoi(m[0][2:])
		catRange := currentRanges[catIdx]
		subRanges := [2][4][2]int{
			currentRanges,
			currentRanges,
		}
		switch {
		case compVal < catRange[0]:
			subRanges[0][catIdx] = [2]int{compVal, compVal}
			subRanges[1][catIdx] = [2]int{catRange[0], catRange[1]}
		case compVal > catRange[1]:
			subRanges[0][catIdx] = [2]int{catRange[0], catRange[1]}
			subRanges[1][catIdx] = [2]int{compVal, compVal}
		default:
			if op == '>' {
				compVal++
			}
			subRanges[0][catIdx] = [2]int{catRange[0], compVal}
			subRanges[1][catIdx] = [2]int{compVal, catRange[1]}
		}
		selectedSubRangeIdx := 0
		if op == '>' {
			selectedSubRangeIdx = 1
		}
		res := 0
		if subRanges[selectedSubRangeIdx][catIdx][0] < subRanges[selectedSubRangeIdx][catIdx][1] {
			res = computeCombinations(wfs, m[1], subRanges[selectedSubRangeIdx])
		}
		total += res
		currentRanges = subRanges[1-selectedSubRangeIdx]
	}
	return total
}

func ratingCatNameToIdx(catName rune) int {
	switch catName {
	case 'x':
		return 0
	case 'm':
		return 1
	case 'a':
		return 2
	default: // s
		return 3
	}
}

func parseInput(input []string) (map[string][]string, [][4]int) {
	wfs := map[string][]string{}
	var parts [][4]int
	workflows := true
	for _, line := range input {
		if line == "" {
			workflows = false
			continue
		}
		if workflows {
			p := strings.Split(line, "{")
			wfs[p[0]] = strings.Split(p[1][:len(p[1])-1], ",")
		} else {
			ps := strings.Split(line[1:len(line)-1], ",")
			m := [4]int{}
			for _, part := range ps {
				n, _ := strconv.Atoi(part[2:])
				m[ratingCatNameToIdx(rune(part[0]))] = n
			}
			parts = append(parts, m)
		}
	}
	return wfs, parts
}
