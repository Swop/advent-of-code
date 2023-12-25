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
	return compute(1, input)
}

func part2(input []string) any {
	return compute(2, input)
}

func compute(part int, input []string) int {
	spelledOut := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	res := 0
	for _, line := range input {
		if part == 2 {
			for i, so := range spelledOut {
				line = strings.ReplaceAll(line, so, so+strconv.Itoa(i)+so)
			}
		}

		l := []byte(line)
		nums := make([]byte, 0)
		for _, c := range l {
			if c >= 48 && c <= 57 {
				nums = append(nums, c)
			}
		}
		nums2 := make([]byte, 2)
		nums2[0] = nums[0] //nolint:gosec
		nums2[1] = nums[len(nums)-1]
		num, _ := strconv.Atoi(string(nums2))
		res += num
	}
	return res
}
