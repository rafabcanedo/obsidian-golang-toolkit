package contextutil_test

import (
	"context"
	"testing"

	"github.com/rafabcanedo/obsidian-golang-toolkit/contextutil"
	"github.com/rafabcanedo/obsidian-golang-toolkit/testutil"
)

func TestWithValue(t *testing.T) {
	ctx := context.Background()
	ctx = contextutil.WithValue(ctx, contextutil.Key("user_id"), 42)

	val, ok := contextutil.Value[int](ctx, contextutil.Key("user_id"))
	testutil.Equal(t, ok, true)
	testutil.Equal(t, val, 42)
}

func TestValue_missingKey(t *testing.T) {
	ctx := context.Background()

	val, ok := contextutil.Value[int](ctx, contextutil.Key("missing"))
	testutil.Equal(t, ok, false)
	testutil.Equal(t, val, 0)
}

func TestValue_wrongType(t *testing.T) {
	ctx := context.Background()
	ctx = contextutil.WithValue(ctx, contextutil.Key("name"), "rafael")

	val, ok := contextutil.Value[int](ctx, contextutil.Key("name"))
	testutil.Equal(t, ok, false)
	testutil.Equal(t, val, 0)
}
