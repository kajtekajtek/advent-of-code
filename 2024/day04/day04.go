package main

import (
	"strings"
)

/* countWord counts the number of times the target word appears in the 
	grid horizontally, vertically, and diagonally in both directions */
func countWord(data, target string) int {
	grid := strings.Split(data, "\n") 
	rows, cols := len(grid), len(grid[0])
	dirs := []int{-1, 0, 1}
	var count int

	/* checkWord checks if the target is present in the grid
	(r,c) - starting position
	(dr, dc) - direction */
	var checkWord = func(r,c, dr, dc int) bool {
		var word string	

		// for every char in target
		for i := range len(target) {
			// calculate the char's position	
			nr, nc := r + i * dr, c + i * dc
			// if the position is out of bounds
			if 0 <= nr && nr < rows && 0 <= nc && nc < cols {
				// append the char to the word
				word += string(grid[nr][nc])
			} else {
				return false
			} 
		
		}
		return word == target
	}

	// for every cell in the grid
	for r := range rows {
		for c := range cols {
			if grid[r][c] == target[0] {
				// for every direction
				for _, dr := range dirs {
					for _, dc := range dirs {
						if dr == 0 && dc == 0 {
							continue
						}
						if checkWord(r, c, dr, dc) {
							count++
						}
					}
				}
			}
			
		}
	}
	return count
}

