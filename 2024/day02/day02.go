package main

import (
	//"fmt"
	"strings"
	"strconv"
)

/* parseDataToLists parses the given data string 
	and returns a list of lists of ints */
func parseDataToLists(data string) ([][]int, error) {
	var listOfLists [][]int
	lines := strings.Split(data, "\n")

	// for line in data
	for _, line := range lines {
		var list []int
		// split the line in to numbers
		stringNumbers := strings.Split(line, " ")

		// convert string numbers to integers and append them to the list
		for _, s := range stringNumbers {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}		
			list = append(list, i)
		}

		// append the list to the list of lists
		listOfLists = append(listOfLists, list)
	}

	return listOfLists, nil
}

func countSafeReports(lists [][]int) int {
	var count int

	// for every report
	for _, report := range lists {
		// check whether report is safe
		if reportIsSafe := isReportSafe(report); reportIsSafe {
			count++
		// if it's not
		} else {
			// for every index in report
			for id := range report {
				// make a copy of report without one of the levels
				copyNumbers := make([]int, len(report))
				copy(copyNumbers, report)
				copyNumbers = append(copyNumbers[:id], copyNumbers[id+1:]...)
				// check whether it's safe now
				reportIsSafe = isReportSafe(copyNumbers)
				if reportIsSafe {
					break
				}
			}
			if reportIsSafe {
				count++
			}
		}
	}

	return count
}

func isReportSafe(report []int) bool {
	asc := false
	// for every number in the report
	for i := range (len(report) - 1) {
		/* if it's the first number, check if the report is 
			ascending or descending */
		if i == 0 {
			if report[i] < report[i + 1] {
				asc = true
			}
		}

		diff := report[i] - report[i + 1]
		// report is not safe if:
		// previous differences were positive and current is negative
		if asc && diff >= 0 {
			return false	
		// previous differences were negative and current is positive
		} else if !asc && diff <= 0 {
			return false
		// difference is greater than 3
		} else if diff < -3 || diff > 3 {
			return false
		}
	}

	return true
}