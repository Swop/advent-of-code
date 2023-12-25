package main

import (
	"sort"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/set"
	strings2 "github.com/Swop/advent-of-code/pkg/strings"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	bCount, holds, heldBy := parseInput(input)
	total := 0
	// from bottom to top
	for i := 0; i < bCount; i++ {
		someBricksOnlyHoldCurrent := false
		for _, j := range holds[i].Values() {
			// if at least one of the brick (H) that is holding the current brick (C) is only holding this current brick (C)
			// then if we remove C, some other bricks will fall
			if len(heldBy[j].Values()) == 1 {
				someBricksOnlyHoldCurrent = true
				break
			}
		}
		if !someBricksOnlyHoldCurrent {
			total++
		}
	}
	return total
}

func part2(input []string) any {
	bCount, _, heldBy := parseInput(input)
	total := 0
	for i := 0; i < bCount; i++ {
		// we simulate the fact that this brick is disintegrated
		disintegrated := set.NewWithValues[int](i)
		for j := i + 1; j < bCount; j++ { // for all the bricks above the current one
			// if all the supporting bricks has been disintegrated, then this brick will fall as well
			if heldBy[j].IsSubset(disintegrated) {
				disintegrated.Add(j)
			}
		}
		// we count the number of bricks that will fall if we disintegrate the current brick
		// (minus 1 which corresponds to the initial brick, which we should not consider in the total)
		total += disintegrated.Size() - 1
	}
	return total
}

func parseInput(input []string) (int, map[int]set.Set[int], map[int]set.Set[int]) {
	var bricks [][2][3]int
	for _, line := range input {
		coords := strings.Split(line, "~")
		b := [2][3]int{
			strings2.To3Int(coords[0], ","),
			strings2.To3Int(coords[1], ","),
		}
		bricks = append(bricks, b)
	}
	// order brinks based on their z coord
	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i][0][2] < bricks[j][0][2]
	})

	holds := map[int]set.Set[int]{}  // brickId -> heldBrickIds
	heldBy := map[int]set.Set[int]{} // brickId -> holdingBrickIds
	heights := map[[2]int][2]int{}   // [x,y] -> [height, brickId]

	for i, brick := range bricks {
		minX := brick[0][0]
		maxX := brick[1][0]
		minY := brick[0][1]
		maxY := brick[1][1]
		minZ := brick[0][2]
		maxZ := brick[1][2]

		bHeldBy := set.New[int]()
		maxH := 0
		// simulate fall down
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				hSet, ok := heights[[2]int{x, y}]
				if !ok {
					// by default, we consider that we don't have any bricks for that coordinates
					// h = 0
					// hBrickID = -1 (no brick)
					hSet = [2]int{0, -1}
					heights[[2]int{x, y}] = hSet
				}
				h, hBrickID := hSet[0], hSet[1]
				if h > maxH {
					maxH = h
					// The brick is held by another brick higher than the previous one
					bHeldBy = set.New[int]()
					bHeldBy.Add(hBrickID)
				} else if h == maxH {
					// The brick is held by another brick at the same height as the previous one
					bHeldBy.Add(hBrickID)
				}
			}
		}
		heldBy[i] = bHeldBy
		for _, b := range bHeldBy.Values() {
			// report inverse holding relation between the two bricks
			if holds[b] == nil {
				holds[b] = set.New[int]()
			}
			holds[b].Add(i)
		}
		z := maxH + maxZ - minZ + 1 // new height of the brick = max height of the bricks below + height of the brick
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				// for the whole horizontal surface (x,y) of the brick, we set the height to the top of the brick,
				// and we declare that this new max height is owned by the current brick
				heights[[2]int{x, y}] = [2]int{z, i}
			}
		}
	}
	// we don't need to return the whole brick list. We just need the number of bricks to iterate on the indices
	// since holds and heldBy are maps, referring to brick IDs (indices)
	return len(bricks), holds, heldBy
}
