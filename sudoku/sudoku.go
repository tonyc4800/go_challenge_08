package sudoku

import "fmt"
import "io/ioutil"

func crossIndex(A string, N string) []string {
	var k string
	var indexes []string
	for _, a := range A {
		for _, n := range N {
			k = string(a) + string(n)
			indexes = append(indexes, k)
		}
	}
	return indexes
}

func solveSudoku(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v\n", string(data))

	// Global board information.  The Sudoku board is assumed to be a standard
	// 9x9 (A-I)x(1-9) grid -- where the first index (upper left) would be `A1`
	// and the last index (lower right) would be `I9`.
	rows := "ABCDEFGHI"
	cols := "123456789"

	ind := crossIndex(rows, cols)
	fmt.Println(ind)

	// Create units. units in this context will be groups of grid indexes that
	// can only contain one instance of a number 1-9.  In Sudoku, a unit will be
	// considered a `rowUnits`(horizontal), a `colUnits`(vertical), and a
	// `blockUnits`(3x3 grid).

	// Create `rowUnits`.
	var rowUnits [][]string
	for _, r := range rows {
		cUnit := crossIndex(string(r), cols)
		rowUnits = append(rowUnits, cUnit)
	}
	fmt.Println(rowUnits)

	// Create `colUnits`.

	// Create `blockUnits`.

	// Convert to grid

	//
	return "n", nil
}
