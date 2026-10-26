package matrix

import (
	"math/big"
)

type Matrix struct {
	Data [][]*big.Rat
	Rows int
	Cols int
}
