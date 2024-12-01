package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"slices"
)

func readTextFile(path string) string { 
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func splitTwoLists(data string) (list1, list2 []int) {
	// split the data into lines
	lines := strings.Split(data, "\n")

	// for line in lines
	for _, line := range lines {
		// split the line into two parts
		a, b, found := strings.Cut(line, "   ")

		if found {
			// convert the two parts into integers and append them to the lists
			if aInt, err := strconv.Atoi(a); err == nil {
				list1 = append(list1, aInt)
			}	
			if bInt, err := strconv.Atoi(b); err == nil {
				list2 = append(list2, bInt)
			}
		}
	}

	return
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
		score += num * count(l2, func(n int) bool {
				return num == n
		})
	}	
	return
}

/* generic function that counts the occurrences of elements that meet 
	given condition in a given slice */
func count[T any](slice []T, f func(T) bool) (count int) {
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return
}

func main() {
	lists := readTextFile("input.txt")
	list1, list2 := splitTwoLists(lists)

	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Println(calculateDistance(list1, list2))

	fmt.Println(similarityScore(list1, list2))
}