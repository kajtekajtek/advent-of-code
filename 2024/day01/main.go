package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"slices"

	"github.com/kajtekajtek/advent-of-code/2024/utils"
)

func parseDataToLists(data string) ([]int, []int, error) {
	var li1, li2 []int
	lines := strings.Split(data, "\n")

	// for line in lines
	for _, line := range lines {
		// split the line into two parts
		a, b, found := strings.Cut(line, "   ")

		if found {
			// convert the two parts into integers and append them to the lists
			aInt, err := strconv.Atoi(a)
			if err != nil {
				return nil, nil, err
			}

			bInt, err := strconv.Atoi(b)
			if err != nil {
				return nil, nil, err
			}

			li1 = append(li1, aInt)
			li2 = append(li2, bInt)
		}
	}

	return li1, li2, nil
}

func calculateDistance(l1, l2 []int) (dist int) {
	for i, _ := range l1 {
		if l1[i] > l2[i] {
			dist += l1[i] - l2[i]
		} else {
			dist += l2[i] - l1[i] 
		}
	}

	return
}

/* similarity score = sum of each number from the left list multiplied by 
	the number of times it appears in the right list */
func similarityScore(l1, l2 []int) (score int) {
	for _, num := range l1 {
		score += num * utils.Count(l2, func(n int) bool {
				return num == n
		})
	}	
	return
}

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