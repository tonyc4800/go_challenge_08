package sudoku

import (
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	var cases = []struct {
		fpath string
		sS    string
	}{
		{
			"./input/puzzle_03.txt",
			"yes",
		},
	}
	for _, c := range cases {
		sS, err := solveSudoku(c.fpath)
		if err != nil {
			t.Errorf("Error retriving solved Sudoku: %v", err)
		}
		if sS != c.sS {
			t.Errorf("solved Sudoku does not match target\n")
		}
	}
}
