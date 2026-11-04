package ptr_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/util/ptr"
)

func TestToPtr(t *testing.T) {
	p := ptr.ToPtr(42)
	testutil.Equal(t, p != nil, true)
	testutil.Equal(t, *p, 42)

	s := ptr.ToPtr("hello")
	testutil.Equal(t, s != nil, true)
	testutil.Equal(t, *s, "hello")
}

func TestFromPtr(t *testing.T) {
	tests := []struct {
		name string
		p    *int
		want int
	}{
		{"nil returns zero", nil, 0},
		{"non-nil returns value", ptr.ToPtr(99), 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, ptr.FromPtr(tt.p), tt.want)
		})
	}
}

func TestFromPtrOr(t *testing.T) {
	tests := []struct {
		name     string
		p        *int
		fallback int
		want     int
	}{
		{"nil returns fallback", nil, 7, 7},
		{"non-nil ignores fallback", ptr.ToPtr(3), 7, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, ptr.FromPtrOr(tt.p, tt.fallback), tt.want)
		})
	}
}
