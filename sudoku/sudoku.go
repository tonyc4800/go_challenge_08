package sudoku

import "fmt"
import "io/ioutil"

// crossIndex 'crosses' two strings such that the two individual values from
// each string join together to create a new value.  For example, if string one
// is "ABC" and string two is "123", the resulting return value will be;
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

// display produces a 2-D grid representation of the current state of the board.
func display(sVals map[string][]string, inds []string) {
	//maxW := 1
	i := 1
	for _, u := range inds {
		vS := sVals[u]
		if len(vS) == 9 {
			fmt.Printf(" . ")
		} else if len(vS) == 1 {
			fmt.Printf(" %v ", vS[0])
		} else {
			fmt.Printf("%v", vS)
		}

		if i%3 == 0 && i%9 != 0 {
			fmt.Printf("|")
		}
		if i%9 == 0 {
			fmt.Printf("\n")
		}
		if i%27 == 0 {
			fmt.Printf("----------------------------\n")
		}
		i++
	}
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

			// TODO: happy path should be along main sightline
			if len(potSol) != 1 {
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
					//fmt.Printf("val:%v p:%v, r:%v\n", val, potSol, rSol)
					sVals[peerI] = rSol
				}
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
func nakedGroup(sVals map[string][]string, indToPeers map[string][]string) map[string][]string {

	// TODO: complete

	// Create reverse map, mapping values (of length n) to indexes.
	ns := [2]int{2, 3}

	// Loop each target group length. e.g. if n=2, we are searching for
	// 'naked_twins'
	for _, n := range ns {
		rsV := make(map[string][]string)

		// loop each index in the Sudoku puzzle
		for k, vs := range sVals {

			// if the number of possible values is equal to our target length,
			// add its details to the rsV[value] = indexes
			if len(vs) == n {
				for _, v := range vs {

					// https://stackoverflow.com/questions/12677934/create-a-golang-map-of-lists
					rsV[v] = append(rsV[v], k)
				}
			}
		}

		// iterate rsV to create a slice of slices that contain indexes
		// where the values only have n solutions
		var ngS [][]string
		for _, vs := range rsV {
			if len(vs) == n {
				ngS = append(ngS, vs)
			}
		}

		// loop all target naked groups and create a set of the intersect peers
		// from all groups.

		for _, ng := range ngS {
			fmt.Println("===================")
			fmt.Println(ng)
			pCount := make(map[string]int)
			var inxP []string
			var pSs [][]string

			// Create a slice of slices that contain all peers for the target
			// index.
			for _, ind := range ng {
				ps := indToPeers[ind]
				pSs = append(pSs, ps)
			}

			// `AND` all slices together to create a slice of all peers at the
			// intersection of the included indexes. - create a mapping of val:count
			for _, ps := range pSs {
				// check if in the set
				for _, v := range ps {
					count, ok := pCount[v]
					if !ok {
						pCount[v] = 1
					} else {
						count++
						pCount[v] = count
					}
				}
			}

			// LOOK INTO: why does toggling this line change the output of
			// fmt.Println(inxP)? Something to due with order/hashing, I'd assume
			//fmt.Println(pCount)

			// all values in the set that are in all ps, are now at the
			// intersection of the indexes
			for k, v := range pCount {
				if v == n {
					inxP = append(inxP, k)
				}
			}

			fmt.Println("-------------------")
			fmt.Println(inxP)

			// for _, d := range ng[0] {
			// 	fmt.Printf("%v-", d)
			// }
			// fmt.Printf("\n")

			fmt.Println("===================")

		}
	}

	// Create slice of all values of length n

	// Iterate the slice of all values of length n and eliminate the values from
	// the peers at the intersection of the naked_group.

	return sVals
}

// reduce applies constraints to the puzzle in attempt to reduce the number of
// potential solutions for each box.  Various methods are applied in loop until
// the methods no longer reduce the size of the puzzle.
func reduce(sVals map[string][]string, unitsAll [][]string, indToPeers map[string][]string) (map[string][]string, bool) {

	improving := true
	for improving {

		// Count how many boxes have been solved before reducing.
		nSolI := 0
		for _, vals := range sVals {
			if len(vals) == 1 {
				nSolI++
			}
		}

		// Attempt to solve puzzle using various strategies.
		sVals = eliminate(sVals, indToPeers)
		sVals = onlyChoice(sVals, unitsAll)
		// TODO: check to see if puzzle is solved before calling nakedGroup
		sVals = nakedGroup(sVals, indToPeers)

		// Count how many boxes are solved after reducing and compare to initial
		// number.
		nSolE := 0
		for _, vals := range sVals {
			if len(vals) == 1 {
				nSolE++
			}
		}

		if nSolE == nSolI {
			// There were no solutions obtained during reduction.
			improving = false
		}

		// Ensure all boxes have at least one possible solution value.
		for _, valS := range sVals {
			if len(valS) == 0 {
				return sVals, false
			}
		}

	}
	return sVals, true
}

// search accepts a map of potential solutions for the Sudoku puzzle, iterates
// all boxes and finds indexes with the fewest possible potential value options.
// a more complete Sudoku puzzle will be returned if possible.
// NOTE: this function is recursive
func search(sVals map[string][]string, unitsAll [][]string, indToPeers map[string][]string) (map[string][]string, bool) {

	// First, reduce the board to eliminate unnecessary work.
	sVals, ok := reduce(sVals, unitsAll, indToPeers)
	if !ok {
		return sVals, false
	}

	// Check if solved and obtain unfilled within min possible solutions.
	// 9 is equal to number of possible values in any given box.
	minV := 9
	var mK string
	for cK, valS := range sVals {

		// Check if any values are unsolved.
		if len(valS) > 1 {

			// Choose a box with the fewest possible solutions
			if len(valS) < minV {
				minV = len(valS)
				mK = cK
				//tempVals := valS
			}
		}
	}
	if minV < 9 {

		// Create a new copy of the Sudoku puzzle.
		sValsCopy := make(map[string][]string)
		for k, v := range sVals {
			sValsCopy[k] = v
		}

		// Attempt solution on new board for each potential value
		for _, pS := range sVals[mK] {

			// Assign one of the values to the position and use recurrence to
			// attempt to solve each resulting puzzle
			sValsCopy[mK] = []string{pS}
			sValsCopy, ok = search(sValsCopy, unitsAll, indToPeers)
			if !ok {
				return sValsCopy, false
			}
			return sValsCopy, true
		}
	}

	return sVals, true
}

func solveSudoku(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

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
			if peer != ind {
				peerSlice = append(peerSlice, peer)
			}
		}

		// assign slice of strings to indToPeers map
		indToPeers[ind] = peerSlice
	}

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

	display(sVals, inds)

	// solve
	sVals, ok := search(sVals, unitsAll, indToPeers)
	if !ok {
		return "n", fmt.Errorf("unsolved puzzle")
	}

	display(sVals, inds)

	return "n", nil
}
