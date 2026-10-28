package contextutil

import "context"

type Key string

func WithValue(ctx context.Context, k Key, val any) context.Context {
	return context.WithValue(ctx, k, val)
}

func Value[T any](ctx context.Context, k Key) (T, bool) {
	val, ok := ctx.Value(k).(T)
	return val, ok
}
