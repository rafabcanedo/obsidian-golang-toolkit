package maputil_test

import (
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/util/maputil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/util/sliceutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := maputil.Keys(m)
	testutil.Equal(t, len(keys), 3)
	testutil.Equal(t, sliceutil.Contains(keys, "a"), true)
	testutil.Equal(t, sliceutil.Contains(keys, "b"), true)
	testutil.Equal(t, sliceutil.Contains(keys, "c"), true)
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	values := maputil.Values(m)
	testutil.Equal(t, len(values), 3)
	testutil.Equal(t, sliceutil.Contains(values, 1), true)
	testutil.Equal(t, sliceutil.Contains(values, 2), true)
	testutil.Equal(t, sliceutil.Contains(values, 3), true)
}

func TestMerge(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"c": 3, "b": 99}
	result := maputil.Merge(a, b)
	testutil.Equal(t, len(result), 3)
	testutil.Equal(t, result["a"], 1)
	testutil.Equal(t, result["c"], 3)
	testutil.Equal(t, result["b"], 99)
}

func TestFilter(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	result := maputil.Filter(m, func(k string, v int) bool { return v > 1 })
	testutil.Equal(t, len(result), 2)
	testutil.Equal(t, result["b"], 2)
	testutil.Equal(t, result["c"], 3)
}
