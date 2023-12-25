package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/slices"
	strings2 "github.com/Swop/advent-of-code/pkg/strings"
)

//go:embed z3_solver.py
var z3Solver []byte

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return resolvePart1(input, [2]float64{200000000000000, 400000000000000})
}

func resolvePart1(input []string, area [2]float64) any {
	hs := parseInput(input)
	total := 0
	for _, c := range slices.Combinations(hs, 2) {
		// y = a1*x + b1 AND y = a2*x + b2 (equations of the lines, x/y being the coordinates of the intersection
		// => a1*x + b1 = a2*x + b2
		// so:
		// => x = (b2 - b1) / (a1 - a2)
		// => y = a1*x + b1
		a1 := float64(c[0][4]) / float64(c[0][3])
		b1 := float64(c[0][1]) - a1*float64(c[0][0])
		a2 := float64(c[1][4]) / float64(c[1][3])
		b2 := float64(c[1][1]) - a2*float64(c[1][0])
		if a1 == a2 {
			continue // lines are parallel
		}
		// intersection point
		x := (b2 - b1) / (a1 - a2)
		y := a1*x + b1

		// time of the intersection (for each line)
		t1 := (x - float64(c[0][0])) / float64(c[0][3])
		t2 := (x - float64(c[1][0])) / float64(c[1][3])

		if t1 >= 0 && t2 >= 0 && x >= area[0] && x <= area[1] && y >= area[0] && y <= area[1] {
			// both times are in the future and the intersection is in the area
			total++
		}
	}
	return total
}

func part2(input []string) any {
	// for this part, it's harder to manually code the equation resolution by hand.
	// Some assumptions:
	// - If we manage to find a line crossing 3 hailstones lines (A, B, C), we can assume that all the other hailstones
	// will cross this line too.
	// Therefore, we can find the line coordinates by solving a problem with 3 constraints
	// (looked-for line being called L):
	// - L cross the first line at some point:
	//     xL + vxL * t = xA + vxA * t
	//     yL + vyL * t = yA + vyA * t
	//     zL + vzL * t = zA + vzA * t
	// - L cross the second line at some point:
	//     xL + vxL * t = xB + vxB * t
	//     yL + vyL * t = yB + vyB * t
	//     zL + vzL * t = zB + vzB * t
	// - L cross the third line at some point:
	//     xL + vxL * t = xC + vxC * t
	//     yL + vyL * t = yC + vyC * t
	//     zL + vzL * t = zC + vzC * t
	// (vxX being the X hailstone's velocity on the X axis, vyX being the X hailstone's velocity on the Y axis, etc.)
	// All the x, y, z, vx, vy, vz of the A, B, C lines are known values.
	// We have 9 equations and 9 unknowns (xL, yL, zL, vxL, vyL, vzL).
	// We can solve this problem with a solver (z3 in this case).
	//
	// Since I didn't find good and/or unarchived implementation (or bindings) of a good solver in Go, I decided to use
	// python3 and z3. The python3 script is in the same directory as this file: z3_solver.py
	// Therefore, this part is not really a Go solution, but a Go + python3 solution, and requires python3 to be
	// installed, in addition to the Python z3-solver lib.
	//
	// This part (Go-side) is only about formatting the data to send to the python script, and parsing the result back,
	// to finally perform the sum to get the final answer.

	hs := parseInput(input)
	// we only need the first 3 hailstones to determine a line crossing the 3 of them and all the other ones.
	xyz := z3Solve(hs[:3])
	return xyz[0] + xyz[1] + xyz[2]
}

func z3Solve(hs [][6]int) [3]int {
	var in string
	// Formatting the data to send to the python script through stdin.
	for i, c := range hs {
		in += fmt.Sprintf("%d %d %d %d %d %d", c[0], c[1], c[2], c[3], c[4], c[5])
		if i < 2 {
			in += "\n"
		}
	}

	// extract python script to a temporary file
	file, err := os.CreateTemp(os.TempDir(), "24-solver.py")
	if err != nil {
		panic(fmt.Errorf("error when creating temporary file for z3 solver: %w", err))
	}
	defer func() { _ = os.Remove(file.Name()) }()
	_, err = file.Write(z3Solver)
	if err != nil {
		panic(fmt.Errorf("error when writing temporary file for z3 solver: %w", err))
	}

	cmd := exec.Command("python3", file.Name()) //nolint:gosec
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	cmd.Stdin = strings.NewReader(in)
	err = cmd.Run()
	out := b.String()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "OUT: %s\n", out)
		panic(fmt.Errorf("error when running z3 solver: %w", err))
	}
	fields := strings.Fields(out)
	return [3]int{
		strings2.ToInt(fields[0]),
		strings2.ToInt(fields[1]),
		strings2.ToInt(fields[2]),
	}
}

func parseInput(input []string) [][6]int {
	var hailstones [][6]int
	for _, line := range input {
		p := strings.ReplaceAll(line, " @ ", ", ")
		p = strings.ReplaceAll(p, ", ", " ")
		p2 := strings.Fields(p)
		hailstones = append(hailstones, [6]int{
			strings2.ToInt(p2[0]),
			strings2.ToInt(p2[1]),
			strings2.ToInt(p2[2]),
			strings2.ToInt(p2[3]),
			strings2.ToInt(p2[4]),
			strings2.ToInt(p2[5]),
		})
	}
	return hailstones
}
