# Sudoku Solver

Go program to solve sudoku puzzles.

## Usage

go run main.go \<sudoku info json file\>

This will display the initial state and the solved sudoku on the screen.

## Sudoku info json file format

The json file contains an array the initial values of the sudoku.
Each item is and array of [zero based row, zero based comumn,sudoku value]. 

For example :

[
	[0, 3, 6],  
	  ..... 
]

On first row from upper left (0) and 4th column (3) place value 6.


