package main

import (
	"fmt"
	"log"

	"github.com/kajtekajtek/advent-of-code/2024/utils"
)

func main() {
	filePath := "input.txt"
	data, err := utils.ReadTextFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// part 1
	rules, updates, err := parseData(data)
	if err != nil {
		log.Fatalf("Failed to parse data: %v", err)
	}

	validUpdates := filterOutInvalidUpdates(rules, updates)
	fmt.Println(sumMiddleNumbers(validUpdates))
}