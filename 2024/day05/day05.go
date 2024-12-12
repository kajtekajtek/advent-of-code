package main

import (
	"strings"
	"strconv"

	"github.com/kajtekajtek/advent-of-code/2024/utils"
)

// parseData splits the input data into rules and updates
func parseData(data string) ([][]int, [][]int, error) {
	var rules, updates [][]int

	// split the data into two parts
	rulesAndUpdates := strings.Split(data, "\n\n")

	// split the parts into separate lines
	rawRules := strings.Split(rulesAndUpdates[0], "\n")
	rawUpdates := strings.Split(rulesAndUpdates[1], "\n")

	// parse the rules and updates into integer slices
	for _, rawRule := range rawRules {
		stringRule := strings.Split(rawRule, "|")
		rule := make([]int, 0)

		for _, s := range stringRule {
			if i, err := strconv.Atoi(s); err != nil {
				return nil, nil, err
			} else {
				rule = append(rule, i)
			}
		}
		rules = append(rules, rule)
	}

	for _, rawUpdate := range rawUpdates {
		stringUpdate := strings.Split(rawUpdate, ",")
		update := make([]int, 0)

		for _, s := range stringUpdate {
			if i, err := strconv.Atoi(s); err != nil {
				return nil, nil, err
			} else {
				update = append(update, i)
			}
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}

func filterOutInvalidUpdates(rules, updates [][]int) [][]int {
	validUpdates := make([][]int, 0)
	// rules map
	rulesMap := make(map[int][]int)

	// fill the rules map
	/* rulesMap[n] = [ numbers that must be printed after 
		n in the valid update ] */ 
	for _, rule := range rules {
		rulesMap[rule[0]] = append(rulesMap[rule[0]], rule[1])
	}

	// traverse every update reverse
	for _, update := range updates {
		isValid := true
		for i := len(update) - 1; i > 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if utils.Contains(rulesMap[update[i]], update[j]) {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}

func sumMiddleNumbers(updates [][]int) int {
	var sum int
	for _, update := range updates {
		if len(update) % 2 == 0 {
			continue
		}
		sum += update[len(update) / 2]
	}

	return sum
}
