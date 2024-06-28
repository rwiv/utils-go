package seq

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, elem := range slice {
		ok := fn(elem)
		if ok {
			result = append(result, elem)
		}
	}
	return result
}

func Map[T any, R any](slice []T, fn func(T) R) []R {
	var result []R
	for _, elem := range slice {
		result = append(result, fn(elem))
	}
	return result
}
