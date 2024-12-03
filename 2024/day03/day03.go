package main

import (
	"regexp"
)

func parseDataToNumberPairs(data string) []string {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	instructions := re.FindAllString(data, -1)

	var numberPairs [][]int
	for _, instr := range instructions {
		re = regexp.MustCompile(`[0-9]{1,3}`)
		numberPairs = append(numberPairs, re.FindAllString(instr,-1))
	}

	return numberPairs
}