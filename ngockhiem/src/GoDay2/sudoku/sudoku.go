package sudoku

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	SUDOKU_LENGTH = 9
)

type Sudoku [9][9]Block
type Block struct {
	Val      int
	Possible []int
}

func (this Sudoku) String() string {
	var out string = ""
	for _, r := range this {
		for _, v := range r {
			out = out + strconv.Itoa(v.Val) + "  "
		}
		out = out + "\n"
	}
	return out
}

func InitSudoku(input [9][9]int) Sudoku {
	var board Sudoku
	for i, v := range input {
		for j, d := range v {
			if d != 0 {
				board[i][j] = Block{Val: d}
			} else {
				board[i][j] = Block{Val: d, Possible: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
			}
		}
	}
	return board
}

// Check whether we can write a value to a block
func (this *Block) CanWrite(value int) bool {
	if this.Val != 0 {
		return false
	}
	for _, val := range this.Possible {
		if val == value {
			return true
		}
	}
	return false
}

func CanWriteBlock(block, value int, sudoku Sudoku) bool {
	tmp := sudoku
	row, col := block/9, block%9
	if tmp[row][col].Val != 0 {
		return false
	}
	RowHave := ScanRow(&tmp, row)
	for _, r := range RowHave {
		if r == value {
			return false
		}
	}
	ColHave := ScanCol(&tmp, col)
	for _, r := range ColHave {
		if r == value {
			return false
		}
	}
	rec := 3*(row/3) + (col / 3)
	RecHave := ScanRec(&tmp, rec)
	for _, r := range RecHave {
		if r == value {
			return false
		}
	}
	return true
}

func ValidValue(block int, sudoku Sudoku) []int {
	valid := []int{}
	tmp := sudoku
	row, col := block/9, block%9
	rec := 3*(row/3) + (col / 3)
	if tmp[row][col].Val != 0 {
		return valid
	}
	RowHave := ScanRow(&tmp, row)
	ColHave := ScanCol(&tmp, col)
	RecHave := ScanRec(&tmp, rec)
	not_valid := append(append(RecHave, RowHave...), ColHave...)
	for i := 1; i < 10; i++ {
		valid_flag := true
		for _, iv := range not_valid {
			if iv == i {
				valid_flag = false
				break
			}
		}
		if valid_flag {
			valid = append(valid, i)
		}
	}
	return valid
}

// remove all element in Block that appear in b
func (this *Block) Unique(b []int) {
	for _, vb := range b {
		for i, dup := range this.Possible {
			if dup == vb {
				this.Possible = append(this.Possible[:i], this.Possible[i+1:]...)
				break
			}
		}
	}
}

// Write to a block
func (this *Block) WriteBlock(value int) {
	if this.Val != 0 {
		panic(fmt.Sprintf("Cant Write to Block value %v", value))
	}
	this.Val = value
}

func WriteBlock2(sudoku Sudoku, block, value int) Sudoku {
	row, col := block/9, block%9
	if sudoku[row][col].Val != 0 || (sudoku[row][col].Val == 0 && !sudoku[row][col].CanWrite(value)) {
		panic(fmt.Sprintf("Cant Write to Block %v, %v, value %v", row, col, value))
	}
	sudoku[row][col].Val = value
	sudoku[row][col].Possible = []int{}
	return sudoku
}

func UnWriteBlock(sudoku *Sudoku, row, col int) {

	sudoku[row][col].Val = 0
}

func IsSolved(sudoku Sudoku) bool {
	//Check columns and rows
	for i := 0; i < SUDOKU_LENGTH; i++ {
		//Check column
		var tmpCol map[int]int = make(map[int]int)
		for j := 0; j < SUDOKU_LENGTH; j++ {
			if sudoku[j][i].Val == 0 {
				return false
			}
			_, ok := tmpCol[sudoku[j][i].Val]
			if ok {
				return false
			} else {
				tmpCol[sudoku[j][i].Val] = 1
			}
		}
		//Check row
		var tmpRow map[int]int = make(map[int]int)
		for j := 0; j < SUDOKU_LENGTH; j++ {
			_, ok := tmpRow[sudoku[i][j].Val]
			if ok {
				return false
			} else {
				tmpRow[sudoku[i][j].Val] = 1
			}
		}
	}
	//Check rectangle
	for i := 0; i < SUDOKU_LENGTH; i += 3 {
		for h := 0; h < SUDOKU_LENGTH; h += 3 {
			var tmpRec map[int]int = make(map[int]int)
			for j := i; j < i+3; j++ {
				for k := h; k < h+3; k++ {
					_, ok := tmpRec[sudoku[j][k].Val]
					if ok {
						return false
					} else {
						tmpRec[sudoku[j][k].Val] = 1
					}
				}
			}
		}

	}
	return true
}

func IsValid(sudoku Sudoku) bool {
	for i := 0; i < SUDOKU_LENGTH; i++ {
		rowNums := ScanRow(&sudoku, i)
		sort.Ints(rowNums)
		for i := 0; i < len(rowNums)-1; i++ {
			if rowNums[i+1] == rowNums[i] {
				return false
			}
		}
		colNums := ScanCol(&sudoku, i)
		sort.Ints(colNums)
		for i := 0; i < len(colNums)-1; i++ {
			if colNums[i+1] == colNums[i] {
				return false
			}
		}
		recNums := ScanRec(&sudoku, i)
		sort.Ints(recNums)
		for i := 0; i < len(recNums)-1; i++ {
			if recNums[i+1] == recNums[i] {
				return false
			}
		}
	}
	return true
}

func ScanRow(sudoku *Sudoku, row int) (have []int) {
	if row > 8 || row < 0 {
		panic("Wrong row input in ScanRow function")
	}
	for j := 0; j < SUDOKU_LENGTH; j++ {
		if sudoku[row][j].Val != 0 {
			have = append(have, sudoku[row][j].Val)
		}
	}
	for i := 0; i < SUDOKU_LENGTH; i++ {
		if sudoku[row][i].Val == 0 {
			sudoku[row][i].Unique(have)
		}
	}
	return have
}

func ScanCol(sudoku *Sudoku, col int) (have []int) {
	if col > 8 || col < 0 {
		panic("Wrong col input in ScanCol function")
	}
	for j := 0; j < SUDOKU_LENGTH; j++ {
		if sudoku[j][col].Val != 0 {
			have = append(have, sudoku[j][col].Val)
		}
	}
	for i := 0; i < SUDOKU_LENGTH; i++ {
		if sudoku[i][col].Val == 0 {
			sudoku[i][col].Unique(have)
		}
	}
	return have
}

func ScanRec(sudoku *Sudoku, rec int) (have []int) {
	row, col := 2*(rec/3)+rec/3, 2*(rec%3)+rec%3
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if sudoku[i][j].Val != 0 {
				have = append(have, sudoku[i][j].Val)
			}
		}
	}
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if sudoku[i][j].Val == 0 {
				sudoku[i][j].Unique(have)
			}
		}
	}
	return have
}

func ScanBoard(sudoku *Sudoku) {
	for i := 0; i < SUDOKU_LENGTH; i++ {
		ScanRow(sudoku, i)
		ScanCol(sudoku, i)
		ScanRec(sudoku, i)
	}
}

func UnAssignBlock(sudoku Sudoku) (int, int) {
	for i, r := range sudoku {
		for j, c := range r {
			if c.Val == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func BacktrackSolve(sudoku *Sudoku) bool {
	row, col := UnAssignBlock(*sudoku)
	if row == -1 {
		return true
	}

	for i := 1; i < 10; i++ {
		if CanWriteBlock(9*row+col%9, i, *sudoku) {
			sudoku[row][col].WriteBlock(i)
			if BacktrackSolve(sudoku) {
				return true
			}
			UnWriteBlock(sudoku, row, col)
		}
	}

	return false
}

func GetSolutions(sudoku *Sudoku, solution_list *[]Sudoku) {
	row, col := UnAssignBlock(*sudoku)
	if row == -1 {
		*solution_list = append(*solution_list, *sudoku)
		return
	}

	for i := 1; i < 10; i++ {
		if CanWriteBlock(9*row+col%9, i, *sudoku) {
			sudoku[row][col].WriteBlock(i)
			GetSolutions(sudoku, solution_list)
			UnWriteBlock(sudoku, row, col)
		}
	}
	return
}
