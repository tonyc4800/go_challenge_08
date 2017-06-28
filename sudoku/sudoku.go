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

	var unitsList [][]string
	// Create `rowUnits`.
	// rowUnits will look similar to > [[A1 A2 A3 A4 A5 A6 A7 A8 A9]...]
	var rowUnits [][]string
	for _, r := range rows {
		rowUnits = append(rowUnits, crossIndex(string(r), cols))
	}
	unitsList = append(unitsList, rowUnits...)

	// Create `colUnits`.
	// colUnits will look similar to > [[A1 B1 C1 D1 E1 F1 G1 H1 I1]...]
	var colUnits [][]string
	for _, c := range cols {
		colUnits = append(colUnits, crossIndex(rows, string(c)))
	}
	unitsList = append(unitsList, colUnits...)

	// Create `blockUnits`.
	// blockUnits will look similar to > [[A1 A2 A3 B1 B2 B3 C1 C2 C3]...]
	var blockUnits [][]string
	rowGroup := [3]string{"ABC", "DEF", "GHI"}
	colGroup := [3]string{"123", "456", "789"}
	for _, ri := range rowGroup {
		for _, ci := range colGroup {
			blockUnits = append(blockUnits, crossIndex(ri, ci))
		}
	}
	unitsList = append(unitsList, blockUnits...)
	fmt.Println(unitsList)

	// Convert to grid

	// solve

	return "n", nil
}
