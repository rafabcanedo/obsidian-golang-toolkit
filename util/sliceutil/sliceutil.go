package sliceutil

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}

	return false
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	result := []T{}

	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))

	for i, v := range slice {
		result[i] = fn(v)
	}

	return result
}

func Unique[T comparable](slice []T) []T {
	seen := map[T]struct{}{}

	result := []T{}

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	result := [][]T{}

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}

		result = append(result, slice[i:end])
	}

	return result
}
