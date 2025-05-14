package main

import "fmt"

// Here is the function that will take values inside the parentheses and will give us a square with these dimensions:
// If the values are negative or zero, the function will not print anything.
func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for drow := 1; drow <= y; drow++ {
		for dcol := 1; dcol <= x; dcol++ {
			// For the top left or bottom left corner
			if (drow == 1 || drow == y) && (dcol == 1) {
				fmt.Print("A")
				// For the top right or bottom right corner
			} else if (drow == 1 || drow == y) && (dcol == x) {
				fmt.Print("C")
				// For the top or bottom side
			} else if drow == 1 || drow == y {
				fmt.Print("B")
				// For the left or right side
			} else if dcol == 1 || dcol == x {
				fmt.Print("B")
				// For inside the square
			} else {
				fmt.Print(" ")
			}
		}
		// Για να πάμε από κάτω
		fmt.Print("\n")
	}

}
