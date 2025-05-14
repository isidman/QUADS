package main

import "fmt"

//QuadC prints a rectangle using 'A' for top corners, 'C' for bottom corners, 'B' for edges
func QuadC(x, y int) {
	// The QuadC function takes two integer parameters, x and y, which represent the width and height of the rectangle respectively.
	// If either x or y is less than or equal to 0, the function returns without printing anything.
	if x <= 0 || y <= 0 {
		return
	}
	// With a nested loop, we iterate through each row and column of the rectangle.

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			// For top or the botton line we check if the row is 1 or y
			// We start with the first row and check if the column is 1 or x
			if row == 1 {
				if col == 1 || col == x {
					fmt.Print("A")
				} else {
					fmt.Print("B")
				}
				// In the last row we check if the column is 1 or x and print C
			} else if row == y {
				if col == 1 || col == x {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
				// For the left or right side we check if the column is 1 or x and print B in the rest of the rows
			} else {
				if col == 1 || col == x {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}

			}
		}
		fmt.Print("\n")
		// We print a new line after each row to move to the next line.
	}

}
