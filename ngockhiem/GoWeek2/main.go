package main

import (
	"GoTraining/ngockhiem/GoWeek2/sudoku"
	"fmt"
)

func main() {
	var input = [9][9]int{
		{0, 7, 8, 1, 0, 0, 0, 2, 0},
		{1, 0, 0, 0, 6, 2, 0, 0, 3},
		{5, 0, 0, 0, 9, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 4, 0, 6},
		{0, 6, 1, 0, 7, 0, 0, 9, 0},
		{0, 9, 0, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 5, 0, 4, 2, 0, 7},
		{6, 0, 0, 0, 8, 0, 0, 3, 0},
		{0, 5, 0, 7, 0, 0, 9, 0, 0}}
	var input2 = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
	_ = input2
	var output = [9][9]int{
		{3, 7, 8, 1, 4, 5, 6, 2, 9},
		{1, 4, 9, 8, 6, 2, 7, 5, 3},
		{5, 2, 6, 3, 9, 7, 1, 4, 8},
		{8, 3, 5, 9, 2, 1, 4, 7, 6},
		{2, 6, 1, 4, 7, 3, 8, 9, 5},
		{7, 9, 4, 6, 5, 8, 3, 1, 2},
		{9, 8, 3, 5, 1, 4, 2, 6, 7},
		{6, 1, 7, 2, 8, 9, 5, 3, 4},
		{4, 5, 2, 7, 3, 6, 9, 8, 1}}

	// Init the input sudoku
	board := sudoku.InitSudoku(input)

	// Init the solution sudoku
	solution := sudoku.InitSudoku(output)
	_ = solution
	// Starting the Algorithm
	sudoku.ScanBoard(&board)
	var solution_list []sudoku.Sudoku
	sudoku.GetSolutions(&board, &solution_list)
	for _, s := range solution_list {
		fmt.Println(s)
	}

}
