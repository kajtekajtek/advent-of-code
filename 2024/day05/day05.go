package main

import (
	"strings"
)

// parseData splits the input data into rules and updates
func parseData(data string) ([][]string, [][]string) {
	// split the data into two parts
	rulesAndUpdates := strings.Split(data, "\n\n")
	// split the parts into separate lines
	rawRules := strings.Split(rulesAndUpdates[0], "\n")
	rawUpdates := strings.Split(rulesAndUpdates[1], "\n")

	// split the rules and updates into slices of strings
	var rules, updates [][]string
	for _, rule := range rawRules {
		rules = append(rules, strings.Split(rule, "|"))
	}
	for _, update := range rawUpdates {
		updates = append(updates, strings.Split(update, ","))
	}

	return rules, updates
}