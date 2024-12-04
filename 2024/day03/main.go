package main

import (
	"fmt"
	"log"

	"github.com/kajtekajtek/advent-of-code/2024/utils"
)

func main() {
	filePath := "input_test.txt"
	data, err := utils.ReadTextFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// part 1

	pairs, err := parseDataToIntegerPairs(data)
	if err != nil {
		log.Fatalf("Failed to parse data: %v", err)
	}

	var sum int
	for _, pair := range pairs {
		sum += pair[0] * pair[1]
	}

	fmt.Println(sum)
}