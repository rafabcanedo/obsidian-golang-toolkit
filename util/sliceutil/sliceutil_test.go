package sliceutil_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/util/sliceutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestContains(t *testing.T) {
    tests := []struct {
        name  string
        slice []int
        item  int
        want  bool
    }{
        {"item exists", []int{1, 2, 3}, 2, true},
        {"item not exists", []int{1, 2, 3}, 5, false},
        {"empty slice", []int{}, 1, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            testutil.Equal(t, sliceutil.Contains(tt.slice, tt.item), tt.want)
        })
    }
}

func TestFilter(t *testing.T) {
    result := sliceutil.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
    testutil.Equal(t, len(result), 2)
    testutil.Equal(t, result[0], 2)
    testutil.Equal(t, result[1], 4)
}

func TestMap(t *testing.T) {
    result := sliceutil.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
    testutil.Equal(t, len(result), 3)
    testutil.Equal(t, result[0], 2)
    testutil.Equal(t, result[1], 4)
    testutil.Equal(t, result[2], 6)
}

func TestUnique(t *testing.T) {
    result := sliceutil.Unique([]int{1, 2, 2, 3, 1})
    testutil.Equal(t, len(result), 3)
    testutil.Equal(t, result[0], 1)
    testutil.Equal(t, result[1], 2)
    testutil.Equal(t, result[2], 3)
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		size    int
		wantLen int
		wantFirst int
		wantLast  int
	}{
		{"even split", []int{1, 2, 3, 4}, 2, 2, 1, 3},
		{"odd split", []int{1, 2, 3, 4, 5}, 2, 3, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceutil.Chunk(tt.slice, tt.size)
			testutil.Equal(t, len(result), tt.wantLen)
			testutil.Equal(t, result[0][0], tt.wantFirst)
			testutil.Equal(t, result[len(result)-1][0], tt.wantLast)
		})
	}
}

func TestChunk_sizeZero(t *testing.T) {
	result := sliceutil.Chunk([]int{1, 2, 3}, 0)
	testutil.Equal(t, len(result), 0)
}
