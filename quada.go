package main

import "fmt"

// Here is the function that will take values inside the parentheses and will give us a square with these dimensions:
// If the values are negative or zero, the function will not print anything.
func QuadA(x, y int) {
	if x <= 0 && y <= 0 {
		return
	}
	// row = row, col = column
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			// For the top left or bottom left corner:
			if (row == 1 || row == y) && (col == 1) {
				fmt.Print("o")
				//For the top right or bottom right corner:
			} else if (row == 1 || row == y) && (col == x) {
				fmt.Print("o")
				// For the top or bottom side:
			} else if row == 1 || row == y {
				fmt.Print("-")
				// For the left or right side:
			} else if col == 1 || col == x {
				fmt.Print("|")
				// For inside the square:
			} else {
				fmt.Print(" ")
			}
		}
		// For going on to the next line:
		fmt.Print("\n")
	}
}
