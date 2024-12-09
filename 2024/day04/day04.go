package main

import (
	"strings"
	"errors"
)

/* checkWord checks if the target string is present in the grid
	(r,c) - starting position
	(dr, dc) - direction */
var checkWord = func(target string, grid []string, r,c, dr, dc int) bool {
	var word string	
	rows, cols := len(grid), len(grid[0])

	// for every char in target
	for i := range len(target) {
		// calculate the char's position	
		nr, nc := r + i * dr, c + i * dc
		// if the position is in bounds
		if 0 <= nr && nr < rows && 0 <= nc && nc < cols {
			// append the char to the word
			word += string(grid[nr][nc])
		} else {
			return false
		} 
	
	}
	return word == target
}

/* countWord counts the number of times the target word appears in the 
	grid horizontally, vertically, and diagonally in both directions */
func countWord(data, target string) int {
	grid := strings.Split(data, "\n") 
	rows, cols := len(grid), len(grid[0])
	dirs := []int{-1, 0, 1}
	var count int

	// for every cell in the grid
	for r := range rows {
		for c := range cols {
			// if the cell contains the first char of the target
			if grid[r][c] == target[0] {
				// for every direction
				for _, dr := range dirs {
					for _, dc := range dirs {
						if dr == 0 && dc == 0 {
							continue
						}
						if checkWord(target, grid, r, c, dr, dc) {
							count++
						}
					}
				}
			}
			
		}
	}
	return count
}

/* countXWord counts the number of times two target words in the shape 
	of an X appear in the grid */
func countXWord(data, target string) (int, error) {
	if len(target) % 2 == 0 {
		return -1, errors.New("Target word can't have an odd number")
	}	

	wingLen := len(target) / 2
	midChar := target[wingLen]
	grid := strings.Split(data, "\n")
	rows, cols := len(grid), len(grid[0])
	dirs := []int{-1, 1}
	var count int

	/* checkXWord checks if the cell is the center of two words in 
		the shape of an X */
	var checkXWord = func(r, c int ) bool {
		wingsMatched := 0

		// calculate the wings ending positions
		var wingEnds [][]int
		for _, dr := range dirs {
			for _, dc := range dirs {
				wingEnds = append(wingEnds, []int{r + dr * wingLen, c + dc * wingLen})
			}
		}
		// for every corner of the X
		for _, wingEnd := range wingEnds {
			/* check if the wing matches the target starting from every 
				corner of the X in the opposite direction */
			wr, wc := wingEnd[0], wingEnd[1]
			dr, dc := -wr + r, -wc + c
			if wingsMatched < 2 && checkWord(target, grid, wr, wc, dr, dc) {
				wingsMatched++
			}
		}

		if wingsMatched == 2 {
			return true
		}
		return false
	}

	// for every cell in the grid
	for r := range rows {
		for c := range cols {
			// if the cell contains the middle char of the target
			if grid[r][c] == midChar {
				if checkXWord(r, c) {
					count++
				}
			}
		}
	}

	return count, nil
}