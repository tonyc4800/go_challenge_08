package sudoku

import "fmt"
import "io/ioutil"

// global board information
// `:=` only works within function bodies
// TODO: eventually, this should be moved within an init function
var rows = "ABCDEFGHI"
var cols = "123456789"

func createBoardIndex(A string, N string) []string {
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

	ind := createBoardIndex(rows, cols)
	// convert to grid
	fmt.Println(ind)
	//
	return "n", nil
}
