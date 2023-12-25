package main

import (
	"strings"

	"github.com/Swop/advent-of-code/pkg/math"
	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/set"
	"github.com/gammazero/deque"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	mods, fromMemory, inverseEdges := parseInput(input)
	return compute(1, mods, fromMemory, inverseEdges)
}

func part2(input []string) any {
	mods, fromMemory, inverseEdges := parseInput(input)
	return compute(2, mods, fromMemory, inverseEdges)
}

const (
	low  = "low"
	high = "high"
)

type state struct {
	pushes int // keep track of the number of button pushes
	// The watched modules are probably at the end of their respective loops (which probably have different sizes).
	// And we're looking for the moment when all the loops synchronises.
	// we're going to keep track of the nodes that received a low pulse, to detect a loop (i.e. when count == 2):
	receivedLows map[string]int
	// and also keep track of last button press count, when that happens, to deduce the size of each loop the watched
	// modules are in:
	previousBtnPressCount map[string]int
	loopSizes             []int                        // This will store the individual loop sizes
	fromMemory            map[string]map[string]string // conjunction mods memory
	onFlipFlops           set.Set[string]              // keep track of which flip-flops are on
	highs                 int                          // This will store the total number of high pulses received
	lows                  int                          // This will store the total number of low pulses received
}

func compute(part int, mods map[string][]string, fromMemory map[string]map[string]string, inverseEdges map[string][]string) any {
	q := deque.New[[3]string]()
	s := &state{
		receivedLows:          map[string]int{},
		previousBtnPressCount: map[string]int{},
		fromMemory:            fromMemory,
		onFlipFlops:           set.New[string](),
	}

	// checks for part 2
	// rx needs to receive a low pulse
	// by looking at the input, we can see that rx has only one parent, which is a conjunction module
	// this parent conjunction module has N parents, all of which are also conjunction modules.
	// Then, rx will only receive a low pulse when its parent emits a high pulse.
	// This will only happen when any of its grandparents synchronously (= at same timing) emit a low pulse.
	// We're going to watch the pulses targeting the grandparents of rx
	watchedMods := set.New[string]()
	if part == 2 {
		for _, grandparent := range inverseEdges[inverseEdges["rx"][0]] {
			watchedMods.Add(grandparent)
		}
	}
	for {
		if total, done := pushButton(part, mods, q, watchedMods, s); done {
			return total
		}
	}
}

func pushButton(part int, mods map[string][]string, q *deque.Deque[[3]string], watchedMods set.Set[string], s *state) (int, bool) {
	if part == 1 && s.pushes == 1000 {
		return s.highs * s.lows, true
	}
	s.pushes++
	q.PushBack([3]string{"broadcaster", "button", low})
	for q.Len() > 0 {
		pulseData := q.PopFront()
		dst := pulseData[0]
		src := pulseData[1]
		pulse := pulseData[2]

		if total, shouldExit := checkExitConditions(part, pulse, s, dst, watchedMods); shouldExit {
			return total, true
		}

		switch pulse {
		case high:
			s.highs++
		case low:
			s.lows++
		}

		if _, ok := mods[dst]; !ok {
			// end node
			continue
		}
		handleMod(q, mods, dst, src, pulse, s)
	}
	return 0, false
}

func handleMod(q *deque.Deque[[3]string], mods map[string][]string, dst, src, pulse string, s *state) {
	switch dst[0] {
	case '%': // flip flop
		if pulse == high {
			// ignore high pulses
			return
		}
		if !s.onFlipFlops.Has(dst) {
			s.onFlipFlops.Add(dst)
			pulse = high
		} else {
			s.onFlipFlops.Remove(dst)
			pulse = low
		}
		for _, e := range mods[dst] {
			q.PushBack([3]string{e, dst, pulse})
		}
	case '&': // conjunction
		s.fromMemory[dst][src] = pulse
		pulse = low
		for _, memPulse := range s.fromMemory[dst] {
			if memPulse == low {
				pulse = high
				break
			}
		}
		for _, e := range mods[dst] {
			q.PushBack([3]string{e, dst, pulse})
		}
	default: // broadcaster
		for _, e := range mods[dst] {
			q.PushBack([3]string{e, dst, pulse})
		}
	}
}

func checkExitConditions(part int, pulse string, s *state, dst string, watchedMods set.Set[string]) (int, bool) {
	if part != 2 {
		return 0, false
	}
	if pulse == low {
		// check if we're closing a loop on the watched modules
		previousPushes, alreadySeen := s.previousBtnPressCount[dst]
		if watchedMods.Has(dst) && alreadySeen && s.receivedLows[dst] == 2 {
			// we're closing a loop: we keep track of the size of the loop, to determine the LCM of them all
			s.loopSizes = append(s.loopSizes, s.pushes-previousPushes)
		}
		s.previousBtnPressCount[dst] = s.pushes
		s.receivedLows[dst]++
	}
	// if we managed to get the size of all the loops, we can compute the LCM of them all
	if len(s.loopSizes) == watchedMods.Size() {
		return math.LCM(s.loopSizes...), true
	}
	return 0, false
}

func parseInput(input []string) (map[string][]string, map[string]map[string]string, map[string][]string) {
	mods := map[string][]string{}                // store modules (nodes)
	fromMemory := map[string]map[string]string{} // store memory of each module (given src pulse type)
	inverseEdges := map[string][]string{}        // store inverse edges

	types := map[string]string{}
	for _, line := range input {
		p := strings.Split(line, " -> ")
		mods[p[0]] = strings.Split(p[1], ", ")
		types[p[0][1:]] = string(p[0][0])
	}
	fullType := func(s string) string {
		if t, ok := types[s]; ok {
			return t + s
		}
		return s
	}
	for n, v := range mods { // remap edges to include module type + build reverse edges
		var edges []string
		for _, e := range v {
			fe := fullType(e)
			edges = append(edges, fe)
			if fe[0] == '&' {
				if _, ok := fromMemory[fe]; !ok {
					fromMemory[fe] = map[string]string{}
				}
				fromMemory[fe][n] = low
			}
			inverseEdges[fe] = append(inverseEdges[fe], n)
		}
		mods[n] = edges
	}
	return mods, fromMemory, inverseEdges
}
