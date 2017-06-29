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

// eliminate iterates the accepted Sudoku puzzle and eliminates values from
// peers of the given box.  A Sudoku puzzle is returned after all values have
// been evaluated.
func eliminate(sVals map[string][]string, indToPeers map[string][]string) map[string][]string {

	// Obtain all solved indexes i.e. it contains only one value.
	var solvI []string
	for indx, vals := range sVals {
		if len(vals) == 1 {
			solvI = append(solvI, indx)
		}
	}

	// Iterate solved values and remove this value from its peers.
	// Loop each index that has been solved.
	for _, si := range solvI {
		val := sVals[si][0]
		peers := indToPeers[si]

		// Iterate indexes that is a peer of a solved index.
		for _, peerI := range peers {
			potSol := sVals[peerI]

			// Copy potential values to new array, excluding the value to remove.
			var rSol []string
			for _, pV := range potSol {
				if pV != val {
					rSol = append(rSol, pV)
				}
			}

			// Ensure the reduced solution slice is the same, or one smaller,
			// than the previous solution slice.
			if len(rSol) == len(potSol) || len(rSol) == len(potSol)-1 {

				// Assign the reduced solution to the index.
				sVals[peerI] = rSol
			}
		}
	}

	return sVals
}

// onlyChoice assigns a value for to a box when there are no other locations
// within a unit for the given value to be placed.
func onlyChoice(sVals map[string][]string, unitList [][]string) map[string][]string {
	for _, u := range unitList {
		for _, d := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {

			// create slice of all locations that could be filled by d within
			// the current unit.
			var locS []string
			for _, ind := range u {
				pS := sVals[ind]
				for _, v := range pS {
					if d == v {
						locS = append(locS, ind)
					}
				}
			}

			// if there is only one location where the value could be placed,
			// assign that value to the location.
			if len(locS) == 1 {
				for _, i := range locS {
					sVals[i] = []string{d}
				}
			}
		}
	}

	return sVals
}

// nakedGroup eliminates values from the passed Sudoku puzzle and eliminates
// values using the naked twins strategy.  A more general form of the naked
// twins strategy is used here, where "naked tripplets" could also be solved.
func nakedGroup(sVals map[string][]string) map[string][]string {
	// TODO: complete

	return sVals
}

// reduce applies constraints to the puzzle in attempt to reduce the number of
// potential solutions for each box.  Various methods are applied in loop until
// the methods no longer reduce the size of the puzzle.
func reduce(sVals map[string][]string, unitsAll [][]string) map[string][]string {
	// TODO: complete
	improving := true
	// TODO: look into make v var here
	sValsRed := make(map[string][]string)
	sValsCurrent := make(map[string][]string)
	for improving {
		// check how many boxes have been solved

		// eliminate

		// only choice
		sValsCurrent = onlyChoice(sValsCurrent, unitsAll)
		// naked_group

		// check how many boxes were solved this round
		// if no improvement, improv = false

	}
	return sValsRed
}

// search accepts a map of potential solutions for the Sudoku puzzle, iterates
// all boxes and finds indexes with the fewest possible potential value options.
// a more complete Sudoku puzzle will be returned if possible.
// NOTE: this function is recursive
func search(sVals map[string][]string) map[string][]string {
	// TODO: complete
	sValsNew := make(map[string][]string)
	// reduce

	// check if solved

	// choose a box with the fewest possible solutions

	// use recurrence to attempt to solve each resulting puzzle
	return sValsNew
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

	// Create slice of all units in the Sudoku board.
	unitsAll := createUnitsSlice(rows, cols)

	// indToUnits is a map of index : its respective units (rows & cols & blocks)
	// i.e. `H8:[[H1 H2 H3 H4 H5 H6 H7 H8 H9] [A8 B8 C8 D8 E8 F8 G8 H8 I8]
	// [G7 G8 G9 H7 H8 H9 I7 I8 I9]]``
	// TODO: Should I be using make here? var (zero/nil) value would be
	// better, then w/in `==` statement, I can check to see if it exists first?
	//var indToUnits map[string][][]string
	indToUnits := make(map[string][][]string)
	for _, ind := range inds {
		for _, unit := range unitsAll {
			// Determine if the target index is contained in the current unit.
			for _, ui := range unit {
				if ind == ui {
					// The value is contained within the unit add unit to map
					// and break the current loop.
					// https://stackoverflow.com/questions/12677934/create-a-golang-map-of-lists
					indToUnits[ind] = append(indToUnits[ind], unit)
					break
				}
			}
		}
	}

	//fmt.Println(indToUnits)

	// indToPeers is a map of index : its respective peers. peers are all grid
	// locations (indexes) in the same unit as a given index, no overlap.
	// i.e."H8:[B8 G7 I9 G8 I8 H1 D8 E8 H8 H9 A8 G9 H2 H3 H5 C8 F8 I7 H4 H6 H7]"
	// TODO: Should I be using make here? var (zero/nil) value would be
	// better, then w/in `==` statement, I can check to see if it exists first?
	// var indToUnits map[string][][]string
	indToPeers := make(map[string][]string)
	for _, ind := range inds {
		peerSet := make(map[string]bool)
		var peerSlice []string
		uS := indToUnits[ind]
		for _, u := range uS {
			// build set of all values within a unit for a target index
			for _, v := range u {
				peerSet[v] = true
			}
		}
		// convert set to slice of strings
		for peer := range peerSet {
			peerSlice = append(peerSlice, peer)
		}
		// assign slice of strings to indToPeers map
		indToPeers[ind] = peerSlice
	}
	//fmt.Println(indToPeers)

	// convert the string representing the board into a grid(map) that maps a
	// key (index) to the values (label for the box, or possible label for the
	// box). for instance, if we know A1=7, map['A1'] = '7', but if the given
	// index is empty (B2, as an example), the corresponding value would be
	// '123456789' (map['B2'] = '123456789')
	// NOTE: though ranging though the data, a seperate index value (`i`) is needed
	// since we only increment the value when we find a character that needs to
	// be matched to the grid index.
	// TODO: this loop should occur before we initialize everything
	// incase the input is faulty
	sVals := make(map[string][]string)
	i := 0
	for _, c := range data {
		switch string(c) {
		case "_":
			sVals[inds[i]] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
			i++
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			sVals[inds[i]] = []string{string(c)}
			i++
		case "\n", " ", "\r":
			continue
		default:
			return "", fmt.Errorf("unexpected value (%v) in Sudoku input", c)
		}
	}
	fmt.Println(sVals)

	// solve

	return "n", nil
}
