The purpose of this exercise is to create rectangle shapes with variable dimentions using specific elements to print them, as long as the dimention units provider are integers and larger than 0 or not negative.

The way to test the usage of the quads 
is to run the testquads.go program. 

Within the testquads.go you can change the parameters in each quad seperately. 

The dimentions are set within the function with the units "X" and "Y" in all of the quads and they all follow a specific set of code with the variables changing in each quad.


Example:__Quad A:__

Using the "o" unit for corners.
The " | " is being used for the left and right side.
The " - " is being used for the top and the bottom side.

Depending on the dimentions set within the test program, the only constant will be the corners and the sides will change on how many characters are being shown.

Example: __Quad B:__

Using the " / " and " \ " unit for corners.
The " * " is being used for the left and right side.
The " * " is being used for the top and the bottom side.

If we set a rectangle with dimentions of 6,4 in the function it means that the index of the x parameter is 6 and the index of the y parameter is 4. In that fashion, the indexes 1 and 6 of the x parameter and the indexes 1 and 4 of the y parameter will be set as " / " and " \ " since they are the corners of the rectangle.

The indexes 2,3,4,5 of the x parameter on the first row and fourth row and the indexes of 2,3 on the first column and the sixth column of the y parameter will be set as " * ". 

The indexes of 2,3,4,5 of the second and third row and the indexes of 2 and 3 of the second, the third, the fourth and the fifth column will be set as empty since we enter the "space" character in them.

