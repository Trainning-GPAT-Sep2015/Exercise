package sudoku_test

import (
	"GoDay2/sudoku"
	"sort"
	"testing"
)

var (
	sample_board = [9][9]int{
		{0, 7, 8, 1, 0, 0, 0, 2, 0},
		{1, 0, 0, 0, 6, 2, 0, 0, 3},
		{5, 0, 0, 0, 9, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 4, 0, 6},
		{0, 6, 1, 0, 7, 0, 0, 9, 0},
		{0, 9, 0, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 5, 0, 4, 2, 0, 7},
		{6, 0, 0, 0, 8, 0, 0, 3, 0},
		{0, 5, 0, 7, 0, 0, 9, 0, 0}}
	sample_board2 = [9][9]int{
		{0, 7, 8, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 6, 2, 0, 0, 0},
		{5, 0, 0, 0, 9, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 9, 0},
		{0, 9, 0, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 5, 0, 4, 2, 0, 0},
		{6, 0, 0, 0, 8, 0, 0, 3, 0},
		{0, 5, 0, 7, 0, 0, 9, 0, 0}}
	sample_board_solution = [9][9]int{
		{3, 7, 8, 1, 4, 5, 6, 2, 9},
		{1, 4, 9, 8, 6, 2, 7, 5, 3},
		{5, 2, 6, 3, 9, 7, 1, 4, 8},
		{8, 3, 5, 9, 2, 1, 4, 7, 6},
		{2, 6, 1, 4, 7, 3, 8, 9, 5},
		{7, 9, 4, 6, 5, 8, 3, 1, 2},
		{9, 8, 3, 5, 1, 4, 2, 6, 7},
		{6, 1, 7, 2, 8, 9, 5, 3, 4},
		{4, 5, 2, 7, 3, 6, 9, 8, 1}}
)

func TestInitSudoku(t *testing.T) {
	output := sudoku.InitSudoku(sample_board)
	var expect sudoku.Sudoku
	for i, v := range sample_board {
		for j, d := range v {
			if d != 0 {
				expect[i][j] = sudoku.Block{Val: d}
			} else {
				expect[i][j] = sudoku.Block{Val: d, Possible: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
			}
		}
	}
	for i, r := range expect {
		for j, v := range r {
			if output[i][j].Val != v.Val {
				t.Errorf("Wrong output from InitSudoku")
			}
		}
	}
}

func TestUnique(t *testing.T) {
	sample := sudoku.Block{Val: 0, Possible: []int{6, 9, 8, 7, 1}}
	have := []int{5, 6, 9, 8}
	expect := sudoku.Block{Val: 0, Possible: []int{1, 7}}
	sample.Unique(have)
	if len(sample.Possible) != len(expect.Possible) {
		t.Errorf("Wrong length")
	}
	sort.Ints(sample.Possible)
	sort.Ints(expect.Possible)
	for i, e := range expect.Possible {
		if sample.Possible[i] != e {
			t.Errorf("Wrong Calculation, expect %v, but got %v", expect, sample)
		}
	}
}

func TestScanRow(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	expect := [9][]int{[]int{7, 8, 1, 2}, []int{1, 6, 2, 3}, []int{5, 9}, []int{8, 4, 6}, []int{6, 1, 7, 9}, []int{9, 3}, []int{5, 4, 2, 7}, []int{6, 8, 3}, []int{5, 7, 9}}
	for i := 0; i < sudoku.SUDOKU_LENGTH; i++ {
		actual := sudoku.ScanRow(&sample_sudoku, i)
		sort.Ints(actual)
		sort.Ints(expect[i])
		for j, a := range actual {
			if a != expect[i][j] {
				t.Errorf("Wrong ScanRow output at line %v", i)
			}
		}
	}
}

func TestScanCol(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	expect := [9][]int{[]int{1, 5, 8, 6}, []int{7, 6, 9, 5}, []int{8, 1}, []int{1, 5, 7}, []int{6, 9, 7, 8}, []int{2, 4}, []int{4, 3, 2, 9}, []int{2, 9, 3}, []int{3, 6, 7}}
	for i := 0; i < sudoku.SUDOKU_LENGTH; i++ {
		actual := sudoku.ScanCol(&sample_sudoku, i)
		sort.Ints(actual)
		sort.Ints(expect[i])
		for j, a := range actual {
			if a != expect[i][j] {
				t.Errorf("Wrong ScanCol output at column %v", i)
			}
		}
	}
}

func TestScanRec(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	expect := [9][]int{[]int{7, 8, 1, 5}, []int{1, 6, 2, 9}, []int{2, 3}, []int{8, 6, 1, 9}, []int{7}, []int{4, 6, 9, 3}, []int{6, 5}, []int{5, 4, 8, 7}, []int{2, 7, 3, 9}}
	for i := 0; i < sudoku.SUDOKU_LENGTH; i++ {
		actual := sudoku.ScanRec(&sample_sudoku, i)
		sort.Ints(actual)
		sort.Ints(expect[i])
		for j, a := range actual {
			if a != expect[i][j] {
				t.Errorf("Wrong ScanRec output at Rectangle %v", i)
			}
		}
	}
}

func TestScanBoard(t *testing.T) {

}

func TestCanWrite(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	sudoku.ScanBoard(&sample_sudoku)
	value := []int{2, 3, 6, 5, 8, 4, 5, 2, 4}
	expect := []bool{false, false, false, false, false, false, true, false, true}
	for i := 0; i < len(value); i++ {
		actual := sample_sudoku[0][i].CanWrite(value[i])
		if actual != expect[i] {
			t.Errorf("CanWrite function fail at block %v, expect %v, but got %v", i, expect[i], actual)
		}
	}
}

func TestCanWriteBlock(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	value := []int{3, 4, 5, 6, 5, 6, 2, 5, 5}
	expect := []bool{true, false, false, false, true, false, false, false, true}
	for i := 0; i < len(value); i++ {
		actual := sudoku.CanWriteBlock(i, value[i], sample_sudoku)
		if expect[i] != actual {
			t.Errorf("CanWriteBlock give wrong output at block %v", i)
		}
	}
}

func TestIsSolved(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	expect := false
	actual := sudoku.IsSolved(sample_sudoku)
	if expect != actual {
		t.Errorf("IsSolved give wrong output")
	}
}

func TestIsValid(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board2)
	expect := true
	actual := sudoku.IsValid(sample_sudoku)
	if expect != actual {
		t.Errorf("IsValid give wrong output")
	}
}

func TestWriteBlock(t *testing.T) {

}

func TestUnWriteBlock(t *testing.T) {

}

func TestUnAssignBlock(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	row, col := sudoku.UnAssignBlock(sample_sudoku)
	expect_row, expect_col := 0, 0
	if row != expect_row || col != expect_col {
		t.Errorf("UnAssignBlock give wrong output expect row %v,col %v but got %v, %v", expect_row, expect_col, row, col)
	}
}

func TestBacktrackSolve(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	sample_sudoku_solution := sudoku.InitSudoku(sample_board_solution)
	result := sudoku.BacktrackSolve(&sample_sudoku)
	if result == false {
		t.Errorf("BacktrackSolve dont give solution")
	}
	for i, r := range sample_sudoku {
		for j, a := range r {
			if sample_sudoku_solution[i][j].Val != a.Val {
				t.Errorf("BacktrackSolve give wrong solution at block %v", 9*i+j)
			}
		}
	}
}

func TestGetSolutions(t *testing.T) {
	sample_sudoku := sudoku.InitSudoku(sample_board)
	single_sudoku_solution := sudoku.InitSudoku(sample_board_solution)
	var solution_list []sudoku.Sudoku
	sudoku.GetSolutions(&sample_sudoku, &solution_list)
	if len(solution_list) != 1 {
		t.Errorf("GetSolutions dont give any solution or more give more than 1")
	}
	for i, r := range single_sudoku_solution {
		for j, a := range r {
			if solution_list[0][i][j].Val != a.Val {
				t.Errorf("RecursiveSolve give wrong solution at block %v", 9*i+j)
			}
		}
	}
}

func BenchmarkBacktrackSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sample_sudoku := sudoku.InitSudoku(sample_board)
		sudoku.BacktrackSolve(&sample_sudoku)
	}
}

func BenchmarkGetSolutions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var solution_list []sudoku.Sudoku
		sample_sudoku := sudoku.InitSudoku(sample_board)
		sudoku.GetSolutions(&sample_sudoku, &solution_list)
	}
}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sample := sudoku.Block{Val: 0, Possible: []int{6, 9, 8, 7, 1}}
		have := []int{5, 6, 9, 8}
		sample.Unique(have)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sudoku.GenerateSudoku()
	}
}
