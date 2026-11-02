package main

import (
	"github.com/jndaf31/term-in-algebra/internal/matrix"
)

func main() {
	in := Inputs{
		MatrixInputs: make([]matrix.Matrix, 2),
	}
	StartRepl(&in)
}
