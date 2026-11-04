package timeutil

import (
	"fmt"
	"time"
)

func BeginOf(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOf(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func IsBetween(t, start, end time.Time) bool {
	return !t.Before(start) && !t.After(end)
}

func FormatDuration(d time.Duration) string {
	if d == 0 {
		return "0s"
	}

	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	switch {
	case h > 0 && m > 0 && s > 0:
		return fmt.Sprintf("%dh %dm %ds", h, m, s)
	case h > 0 && m > 0:
		return fmt.Sprintf("%dh %dm", h, m)
	case h > 0 && s > 0:
		return fmt.Sprintf("%dh %ds", h, s)
	case h > 0:
		return fmt.Sprintf("%dh", h)
	case m > 0 && s > 0:
		return fmt.Sprintf("%dm %ds", m, s)
	case m > 0:
		return fmt.Sprintf("%dm", m)
	default:
		return fmt.Sprintf("%ds", s)
	}
}
