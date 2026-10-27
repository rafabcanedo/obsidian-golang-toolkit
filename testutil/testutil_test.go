package testutil_test

import (
	"errors"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestEqual(t *testing.T) {
	testutil.Equal(t, "a", "a")
	testutil.Equal(t, 42, 42)
}

func TestNoError(t *testing.T) {
	testutil.NoError(t, nil)
}

func TestError(t *testing.T) {
	testutil.Error(t, errors.New("something went wrong"))
}

func TestEqualError(t *testing.T) {
	testutil.EqualError(t, errors.New("boom"), "boom")
}
