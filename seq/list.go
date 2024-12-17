package seq

import "fmt"

func Concat[T any](origin []T, list []T) []T {
	for _, elem := range list {
		origin = append(origin, elem)
	}
	return origin
}

func SubList[T any](list []T, n int) ([][]T, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n should be greater than 0")
	}

	var result [][]T
	var current []T

	for _, elem := range list {
		if len(current) >= n {
			result = append(result, current)
			current = []T{}
		}
		current = append(current, elem)
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result, nil
}
