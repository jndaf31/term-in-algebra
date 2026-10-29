package parse

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/jndaf31/term-in-algebra/internal/matrix"
)

// BuildMatrix
/*
Appends input slice to current input buffer and returns immediately, until
termination symbol is detected. Then, slices the buffer into a slice of string
slices (string matrix). Creates and allocates a matrix of the same size but with
pointers to big.Rat (rational fractional numbers) structs inside. Converts and
assigns the string matrix data directly to big.Rat types using the .SetString
method. Finally, clears the input buffer and returns the big.Rat input matrix.
*/
func BuildMatrix(cIn []string, inputBuff *[]string) (matrix.Matrix, error) {
	// Append input to buffer
	*inputBuff = append(*inputBuff, cIn...)

	// Check for matrix termination signal in the input buffer.
	// Since it is not in this loop, return.
	if len(*inputBuff) > 0 &&
		(*inputBuff)[len(*inputBuff)-1] != "///" {
		return matrix.Matrix{}, nil
	}
	toParse := *inputBuff
	*inputBuff = nil

	// Slice finished string matrix.
	strMatrix, err := SliceSlicer(toParse)
	if err != nil {
		return matrix.Matrix{}, err
	}

	// Allocate each element or result matrix in each row
	resMatrix := make([][]*big.Rat, len(strMatrix))
	for i, row := range strMatrix {
		resMatrix[i] = make([]*big.Rat, len(row))
		for j, val := range row {
			// Direct assignment of the string numeric value to fraction.
			r, ok := new(big.Rat).SetString(val)
			if !ok {
				return matrix.Matrix{},
					fmt.Errorf("invalid number, r: %v, c: %v", i+1, j+1)
			}
			resMatrix[i][j] = r
		}
	}

	Matrix, err := matrix.NewMatrix(resMatrix)
	if err != nil {
		return matrix.Matrix{}, err
	}
	return Matrix, nil
}

// SliceSlicer
/*
Slices the input string slices into one slice per row, delimited by "|".
Creates an 'm * n' matrix of strings with 'n' rows of 'm' elements each.
*/
func SliceSlicer(inputSlice []string) ([][]string, error) {
	var stringMatrix [][]string
	var row []string
	if len(inputSlice) == 0 {
		return [][]string{}, fmt.Errorf("invalid input, blank")
	}
	for _, val := range inputSlice {
		switch val {
		case "":
			continue
		case "|", "///":
			// Append the special characters and start the row over
			stringMatrix = append(stringMatrix, row)
			row = nil
		default:
			row = append(row, val)
		}
	}
	return stringMatrix, nil
}

// CleanInput
/*
Initial input processing. Separates values by whitespace, only accepts numeric
values, except for "|" which is used as a row separation indicator.
Creates a slice of numbers and "|" characters equal to n-1 number of rows.
Ex.: [ 123 1234 | 456 4567 ].
*/
func CleanInput(inputText string) ([]string, error) {
	var outputBuff []string
	if strings.Contains(inputText, "|") {
		inputText = strings.Replace(inputText, "|", " | ", -1)
	}
	for txt := range strings.FieldsSeq(inputText) {
		if strings.Contains(txt, "///") {
			outputBuff = append(outputBuff, strings.Replace(txt, "///", "", -1), "///")
			break
		}
		if strings.Contains(txt, "|") {
			outputBuff = append(outputBuff, strings.Replace(txt, "|", "", -1), "|")
			continue
		}
		if _, ok := new(big.Rat).SetString(txt); !ok {
			return []string{}, fmt.Errorf("invalid input: non numeric")
		}
		outputBuff = append(outputBuff, txt)
	}
	return outputBuff, nil
}
