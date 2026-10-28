package maputil

func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}

func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	result := map[K]V{}
	for k, v := range m {
		if fn(k, v) {
			result[k] = v
		}
	}
	return result
}
