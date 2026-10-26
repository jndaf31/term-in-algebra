package parse

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func BuildMatrix(cIn []string, inputBuff *[]string) (resMatrix [][]*big.Rat, err error) {

	// Append input to buffer
	*inputBuff = append(*inputBuff, cIn...)

	// Check for matrix termination signal in the input buffer.
	// Since it is not in this loop, return.
	if len(*inputBuff) > 0 && (*inputBuff)[len(*inputBuff)-1] != "///" {
		return nil, nil
	}

	// Slice finished string matrix.
	strMatrix, err := SliceSlicer(*inputBuff)
	if err != nil {
		return nil, err
	}

	// Allocate each element in each row
	resMatrix = make([][]*big.Rat, len(strMatrix))
	for i, row := range strMatrix {
		resMatrix[i] = make([]*big.Rat, len(row))
		for j, val := range row {
			// Direct assignment of the string numeric value to fraction.
			r, ok := new(big.Rat).SetString(val)
			if !ok {
				return nil, fmt.Errorf("invalid number, r: %v, c: %v", i+1, j+1)
			}
			resMatrix[i][j] = r
		}
	}
	*inputBuff = []string{}
	return resMatrix, nil
	/*		DEBUGGING BLOCK FOR PRINTING STRARRAY
	for i, row := range strMatrix {
	fmt.Printf("[ ")
	for j, val := range row {
	fmt.Printf("%s ", val)
	}
	fmt.Printf("]\n")
	}
	*/
}

// SliceSlicer
/*
Slices the input string slices into one slice per row, delimited by "/".
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
			stringMatrix = append(stringMatrix, row)
			row = nil
		default:
			row = append(row, val)
		}
	}

	matrixLen := len(stringMatrix[0])
	for _, row := range stringMatrix[1:] {
		if len(row) != matrixLen {
			return [][]string{}, fmt.Errorf("invalid input, all rows must have an equal number of elements")
		}
	}

	return stringMatrix, nil
}

// Compute if each input is numeric
func toFloat64(txt string) (float64, error) {
	return strconv.ParseFloat(txt, 64)
}

// CleanInput
/*
Initial input processing. Separates values by whitespace, only accepts numeric
values, except for "/" which is used as a row separation indicator.
Creates a slice of numbers and "/" characters equal to n-1 number of rows.
Ex.: [ 123 1234 / 456 4567 ].
*/
func CleanInput(inputText string) ([]string, error) {
	var outputBuff []string
	for txt := range strings.FieldsSeq(inputText) {
		if strings.Contains(txt, "///") {
			outputBuff = append(outputBuff, strings.Replace(txt, "///", "", -1), "///")
			break
		}
		if strings.Contains(txt, "|") {
			outputBuff = append(outputBuff, strings.Replace(txt, "|", "", -1), "|")
			continue
		}
		if _, err := toFloat64(txt); err != nil {
			return []string{}, err
		}
		outputBuff = append(outputBuff, txt)
	}
	return outputBuff, nil
}
