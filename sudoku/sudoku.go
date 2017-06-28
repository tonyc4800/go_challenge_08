package sudoku

import "fmt"
import "io/ioutil"

func solveSudoku(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v\n", string(data))
	return "n", nil
}
