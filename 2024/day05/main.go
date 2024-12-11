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
	rules, updates := parseData(data)
	fmt.Println(rules)
	fmt.Println(updates)
}