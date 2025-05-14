package main

import "fmt"

// QuadE prints a rectangle using 'A' for top corners, 'C' for bottom corners, 'B' for edges
func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return // Do nothing if dimensions are not positive
	}

	for row := 0; row < y; row++ { // Loop through each row
		for col := 0; col < x; col++ { // Loop through each column
			if row == 0 { // Top row
				if col == 0 {
					fmt.Print("A") // Top-left corner
				} else if col == x-1 {
					fmt.Print("C") // Top-right corner
				} else {
					fmt.Print("B") // Top edge
				}
			} else if row == y-1 { // Bottom row
				if col == 0 {
					fmt.Print("C") // Bottom-left corner
				} else if col == x-1 {
					fmt.Print("A") // Bottom-right corner
				} else {
					fmt.Print("B") // Bottom edge
				}
			} else { // Middle rows
				if col == 0 || col == x-1 {
					fmt.Print("B") // Side borders
				} else {
					fmt.Print(" ") // Inside space
				}
			}
		}
		fmt.Println() // New line after each row
	}
}
