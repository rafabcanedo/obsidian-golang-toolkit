package randutil_test

import (
	"regexp"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/randutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestString_length(t *testing.T) {
	testutil.Equal(t, len(randutil.String(10)), 10)
	testutil.Equal(t, len(randutil.String(0)), 0)
}

func TestString_charset(t *testing.T) {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	testutil.Equal(t, re.MatchString(randutil.String(50)), true)
}

func TestInt_range(t *testing.T) {
	for range 100 {
		v := randutil.Int(5, 10)
		testutil.Equal(t, v >= 5 && v <= 10, true)
	}
}

func TestUUID_format(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	testutil.Equal(t, re.MatchString(randutil.UUID()), true)
}

func TestUUID_unique(t *testing.T) {
	a, b := randutil.UUID(), randutil.UUID()
	testutil.Equal(t, a != b, true)
}
