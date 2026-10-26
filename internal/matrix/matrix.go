package matrix

import (
	"fmt"
	"math/big"
)

type Matrix struct {
	Data [][]*big.Rat
	Rows int
	Cols int
}

func New(data [][]*big.Rat) (Matrix, error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return Matrix{},
			fmt.Errorf("can't create a matrix with empty data")
	}

	// Check if all the rows have the same number of elements in the matrix.
	for _, row := range data {
		if len(row) != len(data[0]) {
			return Matrix{},
				fmt.Errorf("invalid input, all rows must have an equal number of elements")
		}
	}
	return Matrix{
		data,
		len(data),
		len(data[0]),
	}, nil
}
