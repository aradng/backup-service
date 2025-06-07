package utils

func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range input {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Apply[T any, U any](input []T, predicate func(T) U) []U {
	var result []U
	for _, item := range input {
		result = append(result, predicate(item))
	}
	return result
}

func Reduce[T any, U any](function func(U, T) U, input []T, accumulator U) U{
	for _, item := range input {
		accumulator = function(accumulator, item)
	}
	return accumulator
}