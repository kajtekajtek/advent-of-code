package main

import (
	"regexp"
	"fmt"
	"strconv"
)

func parseDataToIntegerPairs(data string) ([][]int, error) {
	// find all instructions in the data
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	instructions := re.FindAllString(data, -1)

	// extract pairs of numbers from the instructions
	var stringPairs [][]string
	for _, instr := range instructions {
		re = regexp.MustCompile(`[0-9]{1,3}`)
		stringPairs = append(stringPairs, re.FindAllString(instr,-1))
	}

	// convert string pairs to integer pairs
	var intPairs [][]int
	for _, pair := range stringPairs {
		var intPair []int
		s1, s2 := pair[0], pair[1]

		if i1, err := strconv.Atoi(s1); err != nil {
			return nil, err
		} else {
			intPair = append(intPair, i1)
		}
		if i2, err := strconv.Atoi(s2); err != nil {
			return nil, err
		} else {
			intPair = append(intPair, i2)
		}

		intPairs = append(intPairs, intPair)
	} 

	return intPairs, nil
}