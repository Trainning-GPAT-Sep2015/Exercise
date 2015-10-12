package sudoku

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	empty_board = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}
)

func shuffle(s []int) []int {
	index := rand.Perm(len(s))
	result := make([]int, len(s))
	for i := range result {
		result[i] = s[index[i]]
	}
	return result
}

func IsSingleSolution(sudoku *Sudoku, sol_count *int) (int, bool) {
	if CountAssignedBlock(*sudoku) < 15 {
		return 2, false
	}
	row, col := UnAssignBlock(*sudoku)
	if row == -1 {
		*sol_count++
		return *sol_count, true
	}
	for i := 1; i < 10; i++ {
		if CanWriteBlock(9*row+col%9, i, *sudoku) {
			sudoku[row][col].WriteBlock(i)
			IsSingleSolution(sudoku, sol_count)
			UnWriteBlock(sudoku, row, col)
			if *sol_count > 1 {
				break
			}
		}

	}
	if *sol_count == 1 {
		return *sol_count, true
	}
	return *sol_count, false
}

func CheckSingle(sudoku *Sudoku) (sol_num int, is_single bool) {
	var sol_count int = 0
	sol_num, is_single = IsSingleSolution(sudoku, &sol_count)
	return
}
func GenerateSudoku() Sudoku {
	rand.Seed(time.Now().UnixNano())
	var generate Sudoku
	generate = InitSudoku(empty_board)
	var sol_count int = 0
	addBlocks(&generate, &sol_count)
	return generate
}

func CountEmptyBlockWithMultiVal(sudoku Sudoku) (num int) {
	num = 0
	for i, r := range sudoku {
		for j, c := range r {
			if c.Val == 0 && len(ValidValue(9*i+j, sudoku)) > 1 {
				num++
			}
		}
	}
	return num
}
func CountAssignedBlock(sudoku Sudoku) int {
	num := 0
	for _, r := range sudoku {
		for _, v := range r {
			if v.Val != 0 {
				num++
			}
		}
	}
	return num
}
func RandomUnAssignBlock(sudoku Sudoku) (int, int) {
	random_block := rand.Intn(CountEmptyBlockWithMultiVal(sudoku))
	for i, r := range sudoku {
		for j, c := range r {
			if c.Val == 0 && len(ValidValue(9*i+j, sudoku)) > 1 {
				random_block--
			}
			if random_block == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func addBlocks(sudoku *Sudoku, sol_count *int) {
	if CountAssignedBlock(*sudoku) > 15 {
		sol_num, is_single := CheckSingle(sudoku)
		if is_single {
			*sol_count = 1
			return
		} else if sol_num == 0 {
			*sol_count = -1
			return
		}
	}
	row, col := RandomUnAssignBlock(*sudoku)
	if row == -1 {
		*sol_count = -1
		return
	}
	valid_value := ValidValue(9*row+col, *sudoku)
	valid_value = shuffle(valid_value)
	for _, i := range valid_value {
		if CanWriteBlock(9*row+col%9, i, *sudoku) {
			sudoku[row][col].WriteBlock(i)
			fmt.Println("try", row, col, i)
			addBlocks(sudoku, sol_count)
			if *sol_count == 1 {
				break
			}
			if *sol_count == -1 {
				UnWriteBlock(sudoku, row, col)
			}
		}
	}
	return
}
