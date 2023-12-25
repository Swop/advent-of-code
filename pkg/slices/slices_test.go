package slices

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSlices_Combinations(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	tests := []struct {
		name    string
		setSize int
		want    [][]int
	}{
		{
			name:    "set of 2",
			setSize: 2,
			want:    [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5}},
		},
		{
			name:    "set of 3",
			setSize: 3,
			want:    [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combinations(input, tt.setSize)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
