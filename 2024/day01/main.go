package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/kajtekajtek/advent-of-code/2024/utils"
)

func main() {
	filePath := "input.txt"
	data, err := utils.ReadTextFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	list1, list2, err := parseDataToLists(data)
    if err != nil {
        log.Fatalf("Failed to parse data: %v", err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Println(calculateDistance(list1, list2))
	fmt.Println(similarityScore(list1, list2))
}