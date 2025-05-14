package main

import "fmt"

// QuadB prints a rectangle using the following characters:
// Top-left corner: '/'
// Top-right corner: '\'
// Bottom-left corner: '\'
// Bottom-right corner: '/'
// Horizontal edges: '*'
// Vertical edges: '*'
// Inside: space (' ')
func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		// If width or height is zero or negative, do nothing
		return
	}

	// Loop through each row of the rectangle
	for row := 0; row < y; row++ {
		// Loop through each column (character) in the current row
		for col := 0; col < x; col++ {

			// First row (top)
			if row == 0 {
				if col == 0 {
					fmt.Print("/") // Top-left corner
				} else if col == x-1 {
					fmt.Print("\\") // Top-right corner
				} else {
					fmt.Print("*") // Top edge (middle)
				}

				// Last row (bottom)
			} else if row == y-1 {
				if col == 0 {
					fmt.Print("\\") // Bottom-left corner
				} else if col == x-1 {
					fmt.Print("/") // Bottom-right corner
				} else {
					fmt.Print("*") // Bottom edge (middle)
				}

				// Middle rows (not top or bottom)
			} else {
				if col == 0 || col == x-1 {
					fmt.Print("*") // Left and right edges
				} else {
					fmt.Print(" ") // Inside of the rectangle
				}
			}
		}

		// After printing one full row, move to the next line
		fmt.Println()
	}
}
