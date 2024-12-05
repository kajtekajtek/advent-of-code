package main

import (
	"regexp"
	"strconv"
)

func sumPairs(pairs [][]int) int {
	var sum int
	for _, pair := range pairs {
		sum += pair[0] * pair[1]
	}
	return sum
}

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

func parseEnabledDataToIntegerPairs(data string) ([][]int, error) {
	var enabled, rest, allEnabled string

	// find the first don't instruction
	enabled, rest = findDont(data)
	allEnabled += enabled

	// while there are no more don't's nor do's
	for rest != "" {
		// find the next do instruction
		_, rest = findDo(rest)
		// if found
		if rest != "" {
			// find the next don't instruction
			enabled, rest = findDont(rest)
			allEnabled += enabled
		}
	}
	return parseDataToIntegerPairs(allEnabled)
}

/* findDont returns the data string split in two parts at 
	the nearest don't instruction */
func findDont(data string) (string, string) {
	re := regexp.MustCompile(`don't\(\)`)
	nearestDont := re.FindStringIndex(data)
	if nearestDont == nil {
		return data, ""
	}
	return data[:nearestDont[0]], data[nearestDont[1]:]
}

/* findDo returns the data string split in two parts at
	the nearest do instruction */
func findDo(data string) (string, string) {
	re := regexp.MustCompile(`do\(\)`)
	nearestDo := re.FindStringIndex(data)
	if nearestDo == nil {
		return data, ""
	}
	return data[:nearestDo[0]], data[nearestDo[1]:]
}