package timeutil_test

import (
	"testing"
	"time"

	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/util/timeutil"
)

func TestBeginOf(t *testing.T) {
	input := time.Date(2024, 6, 15, 14, 32, 45, 0, time.UTC)
	want := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	testutil.Equal(t, timeutil.BeginOf(input), want)
}

func TestEndOf(t *testing.T) {
	input := time.Date(2024, 6, 15, 14, 32, 45, 0, time.UTC)
	want := time.Date(2024, 6, 15, 23, 59, 59, 999999999, time.UTC)
	testutil.Equal(t, timeutil.EndOf(input), want)
}

func TestIsBetween(t *testing.T) {
	start := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		t    time.Time
		want bool
	}{
		{"inside range", time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC), true},
		{"equal to start", start, true},
		{"equal to end", end, true},
		{"before start", time.Date(2024, 5, 31, 0, 0, 0, 0, time.UTC), false},
		{"after end", time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, timeutil.IsBetween(tt.t, start, end), tt.want)
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{"zero", 0, "0s"},
		{"seconds only", 45 * time.Second, "45s"},
		{"minutes only", 30 * time.Minute, "30m"},
		{"hours only", 2 * time.Hour, "2h"},
		{"minutes and seconds", 30*time.Minute + 5*time.Second, "30m 5s"},
		{"hours and minutes", 2*time.Hour + 30*time.Minute, "2h 30m"},
		{"hours and seconds", 2*time.Hour + 5*time.Second, "2h 5s"},
		{"hours minutes seconds", 2*time.Hour + 30*time.Minute + 5*time.Second, "2h 30m 5s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equal(t, timeutil.FormatDuration(tt.d), tt.want)
		})
	}
}
