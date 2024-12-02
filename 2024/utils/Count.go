package utils

// Count counts the occurences of elements that meet a given condition in a slice 
func Count[T any](slice []T, f func(T) bool) (count int) {
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return
}