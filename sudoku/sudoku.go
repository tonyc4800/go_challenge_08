package sudoku

import "fmt"
import "io/ioutil"

// crossIndex 'crosses'|'zips' two strings such that the two individual values
// from each string join together to create a new value.  For example, if string
// one is "ABC" and string two is "123", the resulting return value will be
// ["A1","A2","A3","B1","B2","B3","C1","C2","C3"].
func crossIndex(A string, N string) []string {
	var ks []string
	for _, a := range A {
		for _, n := range N {
			ks = append(ks, (string(a) + string(n)))
		}
	}
	return ks
}

// createUnitsSlice creates all units. units in this context will be groups of
// grid indexes that can only contain one instance of a number 1-9.  In Sudoku,
// a unit will be considered a `rowUnits`(horizontal), a `colUnits`(vertical),
// and a `blockUnits`(3x3 grid).
func createUnitsSlice(rows string, cols string) [][]string {
	var unitsSlice [][]string

	// Create `rowUnits` and append to slice of all units.
	// i.e. > [[A1 A2 A3 A4 A5 A6 A7 A8 A9]...]
	var rowUnits [][]string
	for _, r := range rows {
		rowUnits = append(rowUnits, crossIndex(string(r), cols))
	}
	unitsSlice = append(unitsSlice, rowUnits...)

	// Create `colUnits` and append to slice of all units.
	// i.e. > [[A1 B1 C1 D1 E1 F1 G1 H1 I1]...]
	var colUnits [][]string
	for _, c := range cols {
		colUnits = append(colUnits, crossIndex(rows, string(c)))
	}
	unitsSlice = append(unitsSlice, colUnits...)

	// Create `blockUnits` and append to slice of all units.
	// i.e. > [[A1 A2 A3 B1 B2 B3 C1 C2 C3]...]
	var blockUnits [][]string
	rowGroup := [3]string{"ABC", "DEF", "GHI"}
	colGroup := [3]string{"123", "456", "789"}
	for _, ri := range rowGroup {
		for _, ci := range colGroup {
			blockUnits = append(blockUnits, crossIndex(ri, ci))
		}
	}
	unitsSlice = append(unitsSlice, blockUnits...)

	return unitsSlice
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

	inds := crossIndex(rows, cols)
	fmt.Println(inds)

	// Create slice of all units in the Sudoku board.
	unitsAll := createUnitsSlice(rows, cols)
	//fmt.Println(unitsAll)

	// create map of index : its respective units (rows & cols & blocks)
	// i.e. map['A1'] = [["A2", "A3", "A4", ...],....]
	// TODO: Should I really be using make here? var (zero/nil) value would be
	// better, then w/in `==` statement, I can check to see if it exists first?
	//var indToUnits map[string][][]string
	indToUnits := make(map[string][][]string)
	for _, ind := range inds {
		for _, unit := range unitsAll {
			// determine if the target index is contained in the current unit
			for _, ui := range unit {
				if ind == ui {
					// add unit to map and break loop
					indToUnits[ind] = append(indToUnits[ind], unit)
					break
				}
			}
		}
	}

	fmt.Println(indToUnits)

	// Convert to grid

	// solve

	return "n", nil
}
