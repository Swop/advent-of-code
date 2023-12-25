package main

import (
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/dynprog"
	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/slices"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return compute(parseInput(1, input))
}

func part2(input []string) any {
	return compute(parseInput(5, input))
}

func compute(rec []record) any {
	total := 0

	var memoized dynprog.MemoizedFunc[int]
	memoizeHashFunc := func(args ...any) string {
		return args[0].(string) + "|" + strings.Join(slices.IntsToStrings(args[1].([]int)), ",")
	}
	memoizeFunc := func(args ...any) int {
		return arrangementCount(memoized, args[0].(string), args[1].([]int))
	}
	memoized = dynprog.Memoize(memoizeHashFunc, memoizeFunc)
	for _, r := range rec {
		v := memoized(r.spring, r.data)
		total += v
	}
	return total
}

func arrangementCount(memoized dynprog.MemoizedFunc[int], s string, d []int) int {
	if len(s) == 0 {
		if len(d) == 0 {
			return 1
		}
		return 0
	}
	switch {
	case strings.HasPrefix(s, "."):
		return memoized(s[1:], d)
	case strings.HasPrefix(s, "?"):
		return memoized(s[1:], d) + // "."
			memoized("#"+s[1:], d) // "#"
	case strings.HasPrefix(s, "#"):
		if len(d) == 0 {
			return 0
		}
		if len(s) < d[0] {
			return 0
		}
		for _, c := range s[:d[0]] {
			if c == '.' {
				return 0
			}
		}
		if len(d) > 1 {
			if len(s) < d[0]+1 || s[d[0]] == '#' {
				return 0
			}
			return memoized(s[d[0]+1:], d[1:])
		}
		return memoized(s[d[0]:], d[1:])
	default:
		panic("impossible branch")
	}
}

type record struct {
	spring string
	data   []int
}

func parseInput(unfoldMultiplier int, input []string) []record {
	var records []record
	for _, line := range input {
		f := strings.Fields(line)
		var spring string
		for i := 0; i < unfoldMultiplier; i++ {
			spring += f[0]
			if i < unfoldMultiplier-1 {
				spring += "?"
			}
		}
		dataF := strings.Split(f[1], ",")
		data := make([]int, len(dataF)*unfoldMultiplier)
		for i := 0; i < unfoldMultiplier; i++ {
			for dfi, df := range dataF {
				d, _ := strconv.Atoi(df)
				data[i*len(dataF)+dfi] = d
			}
		}
		records = append(records, record{
			spring: spring,
			data:   data,
		})
	}
	return records
}
