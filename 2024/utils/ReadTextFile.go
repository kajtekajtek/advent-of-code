package utils

import (
	"os"
)

// ReadTextFile returns the contents of a text file as a string
func ReadTextFile(path string) (string, error) { 
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
