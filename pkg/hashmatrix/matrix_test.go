package hashmatrix

import (
	"testing"

	"github.com/Swop/advent-of-code/pkg/ptr"
	"github.com/google/go-cmp/cmp"
)

type dataPoint struct {
	x, y int
	val  int
}

func TestMatrix_Neighbors_2D(t *testing.T) {
	// 1 . 3
	// 4 5 6
	// 7 8 .
	data := []dataPoint{
		{0, 0, 1},
		{2, 0, 3},

		{0, 1, 4},
		{1, 1, 5},
		{2, 1, 6},

		{0, 2, 7},
		{1, 2, 8},
	}
	tests := []struct {
		name             string
		pos              Pos2D
		includeDiagonals bool
		data             []dataPoint
		want             []Neighbor[Pos2D, int]
	}{
		{
			name:             "center_with_diagonals",
			pos:              Pos2D{1, 1},
			data:             data,
			includeDiagonals: true,
			want: []Neighbor[Pos2D, int]{
				{Position: Pos2D{0, 0}, Vector: Pos2D{-1, -1}, Value: ptr.To(1)},
				{Position: Pos2D{1, 0}, Vector: Pos2D{0, -1}, Value: nil},
				{Position: Pos2D{2, 0}, Vector: Pos2D{1, -1}, Value: ptr.To(3)},
				{Position: Pos2D{0, 1}, Vector: Pos2D{-1, 0}, Value: ptr.To(4)},
				{Position: Pos2D{2, 1}, Vector: Pos2D{1, 0}, Value: ptr.To(6)},
				{Position: Pos2D{0, 2}, Vector: Pos2D{-1, 1}, Value: ptr.To(7)},
				{Position: Pos2D{1, 2}, Vector: Pos2D{0, 1}, Value: ptr.To(8)},
				{Position: Pos2D{2, 2}, Vector: Pos2D{1, 1}, Value: nil},
			},
		},
		{
			name:             "center_without_diagonals",
			pos:              Pos2D{1, 1},
			data:             data,
			includeDiagonals: false,
			want: []Neighbor[Pos2D, int]{
				{Position: Pos2D{1, 0}, Vector: Pos2D{0, -1}, Value: nil},
				{Position: Pos2D{0, 1}, Vector: Pos2D{-1, 0}, Value: ptr.To(4)},
				{Position: Pos2D{2, 1}, Vector: Pos2D{1, 0}, Value: ptr.To(6)},
				{Position: Pos2D{1, 2}, Vector: Pos2D{0, 1}, Value: ptr.To(8)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New[Pos2D, int]()
			for _, d := range tt.data {
				m.Set(Pos2D{d.x, d.y}, d.val)
			}
			got := m.Neighbors(tt.pos, tt.includeDiagonals)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
