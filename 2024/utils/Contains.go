package utils

// Contains checks if a slice contains a given element
func Contains[T comparable](slice []T, elem T) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}
	return false
}