package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/gammazero/deque"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	seedRanges, steps := parseInput(1, input)
	return getMinimumLocation(seedRanges, steps)
}

func part2(input []string) any {
	seedRanges, steps := parseInput(2, input)
	return getMinimumLocation(seedRanges, steps)
}

func getMinimumLocation(seedRanges [][2]int, steps map[string]*Step) int {
	minLocation := math.MaxInt
	for _, seedRange := range seedRanges {
		for _, r := range mapRange(steps, seedRange, "seed") {
			if r[0] < minLocation {
				minLocation = r[0]
			}
		}
	}
	return minLocation
}

func mapRange(steps map[string]*Step, inputRange [2]int, inputType string) [][2]int {
	step, ok := steps[inputType]
	if !ok {
		return [][2]int{inputRange}
	}
	q := deque.New[[2]int]()
	q.PushBack(inputRange)
	subRanges := make([][2]int, 0, 100)

	availableRanges := map[string]struct{}{}
	for k := range step.ranges {
		availableRanges[k] = struct{}{}
	}

queueLoop:
	for q.Len() > 0 {
		toMap := q.PopFront()
		for k, r := range step.ranges {
			if _, ok := availableRanges[k]; !ok {
				continue
			}
			additionalSubRange := mapToSubRange(q, toMap, r)
			if additionalSubRange != nil {
				subRanges = append(subRanges, *additionalSubRange)
				delete(availableRanges, k)
				continue queueLoop
			}
		}
		// out of all ranges
		subRanges = append(subRanges, [2]int{toMap[0], toMap[1]})
	}

	var mappedRanges [][2]int
	for _, r := range subRanges {
		mappedRanges = append(mappedRanges, mapRange(steps, r, step.destType)...)
	}
	return mappedRanges
}

func mapToSubRange(q *deque.Deque[[2]int], toMap [2]int, r *StepRangeDefinition) *[2]int {
	if toMap[0] < r.sourceRange[0] && toMap[1] > r.sourceRange[1] {
		// wraps range
		q.PushBack([2]int{toMap[0], r.sourceRange[0]})
		q.PushBack([2]int{r.sourceRange[1], toMap[1]})
		return &[2]int{
			r.destRange[0],
			r.destRange[1],
		}
	}
	if toMap[0] >= r.sourceRange[0] && toMap[1] <= r.sourceRange[1] {
		// fully in range
		return &[2]int{
			r.destRange[0] + (toMap[0] - r.sourceRange[0]),
			r.destRange[0] + (toMap[1] - r.sourceRange[0]),
		}
	}
	if toMap[0] < r.sourceRange[0] && (toMap[1] < r.sourceRange[1] && toMap[1] >= r.sourceRange[0]) {
		// partially in range (right part)
		q.PushBack([2]int{toMap[0], r.sourceRange[0]})
		return &[2]int{
			r.destRange[0],
			r.destRange[0] + (toMap[1] - r.sourceRange[0]),
		}
	}
	if (toMap[0] >= r.sourceRange[0] && toMap[0] < r.sourceRange[1]) && toMap[1] > r.sourceRange[1] {
		// partially in range (left part)
		q.PushBack([2]int{r.sourceRange[1], toMap[1]})
		return &[2]int{
			r.destRange[0] + (toMap[0] - r.sourceRange[0]),
			r.destRange[1],
		}
	}
	return nil
}

type StepRangeDefinition struct {
	sourceRange [2]int
	destRange   [2]int
}

type Step struct {
	sourceType string
	destType   string
	ranges     map[string]*StepRangeDefinition
}

func parseInput(part int, input []string) ([][2]int, map[string]*Step) {
	descRxp := regexp.MustCompile(`^(.*)-to-(.*) map:$`)
	var seedRanges [][2]int
	seedStrs := strings.Fields(input[0][7:])
	if part == 1 {
		for i := 0; i < len(seedStrs); i++ {
			start, _ := strconv.Atoi(seedStrs[i])
			seedRanges = append(seedRanges, [2]int{start, start + 1})
		}
	} else {
		for i := 0; i < len(seedStrs); i += 2 {
			start, _ := strconv.Atoi(seedStrs[i])
			length, _ := strconv.Atoi(seedStrs[i+1])
			seedRanges = append(seedRanges, [2]int{start, start + length})
		}
	}
	steps := map[string]*Step{}
	var currentStep *Step
	for _, line := range input[2:] {
		if line == "" {
			continue
		}
		matches := descRxp.FindStringSubmatch(line)
		if len(matches) > 0 {
			if currentStep != nil {
				steps[currentStep.sourceType] = currentStep
			}
			currentStep = &Step{
				sourceType: matches[1],
				destType:   matches[2],
				ranges:     map[string]*StepRangeDefinition{},
			}
			continue
		}

		nums := strings.Fields(line)
		srcStart, _ := strconv.Atoi(nums[1])
		dstStart, _ := strconv.Atoi(nums[0])
		length, _ := strconv.Atoi(nums[2])

		stepRangeDef := &StepRangeDefinition{
			sourceRange: [2]int{srcStart, srcStart + length},
			destRange:   [2]int{dstStart, dstStart + length},
		}
		currentStep.ranges[fmt.Sprintf("%d-%d", stepRangeDef.sourceRange[0], stepRangeDef.sourceRange[1])] = stepRangeDef
	}
	if currentStep != nil {
		steps[currentStep.sourceType] = currentStep
	}

	return seedRanges, steps
}
