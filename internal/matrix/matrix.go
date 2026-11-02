package matrix

import (
	"fmt"
	"math/big"
)

type Matrix struct {
	data [][]*big.Rat
	rows int
	cols int
}

func (m Matrix) Rows() int {
	return m.rows
}

func (m Matrix) Cols() int {
	return m.cols
}

// PrintMatrix - prints the formatted matrix to the console.
func (m Matrix) PrintMatrix() {
	var numMatrix = make([][]float64, m.rows)

	// FIXME weird bug here
	for i, row := range m.data {
		numMatrix[i] = make([]float64, m.cols)
		for j, val := range row {
			v, _ := val.Float64()

			numMatrix[i][j] = v
		}
	}
	for _, row := range numMatrix {
		fmt.Printf("%e\n", row)
	}
}

// GetData - return a copy of the data block (equivalent to applying .GetValue()
// to every element)
func (m Matrix) GetData() [][]*big.Rat {
	var dataBlock = make([][]*big.Rat, m.rows)
	for i, row := range m.data {
		dataBlock[i] = make([]*big.Rat, m.cols)
		for j := range row {
			dataBlock[i][j] = m.GetValue(i, j)
		}
	}
	return dataBlock
}

// GetValue - Returns the numeric value of the element ar (row, col).
func (m Matrix) GetValue(row, col int) *big.Rat {
	return new(big.Rat).Set(m.data[row][col])
}

// NewMatrix - Returns a new Matrix object with the given values assigned to it.
func NewMatrix(data [][]*big.Rat) (Matrix, error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return Matrix{},
			fmt.Errorf("can't create a matrix with empty data")
	}

	// Check if all the rows have the same number of elements in the matrix and make a copy
	copyBlock := make([][]*big.Rat, len(data))
	for i, row := range data {
		if len(row) != len(data[0]) {
			return Matrix{}, fmt.Errorf("all matrix rows must be the same size")
		}
		copyBlock[i] = make([]*big.Rat, len(row))
		for j, val := range row {
			if val == nil {
				return Matrix{}, fmt.Errorf("nil value in new matrix")
			}
			copyBlock[i][j] = new(big.Rat).Set(val)
		}
	}
	return Matrix{
		copyBlock,
		len(data),
		len(data[0]),
	}, nil
}
