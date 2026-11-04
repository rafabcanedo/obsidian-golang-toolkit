package ptr

func ToPtr[T any](v T) *T {
	return &v
}

func FromPtr[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}

	return *p
}

func FromPtrOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}

	return *p
}
