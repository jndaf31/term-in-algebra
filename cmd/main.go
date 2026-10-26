package main

import (
	"github.com/jndaf31/term-in-algebra/internal/matrix"
	"github.com/jndaf31/term-in-algebra/internal/repl"
)

func main() {
	in := repl.Inputs{
		MatrixInputs: make([]matrix.Matrix, 2),
	}
	repl.StartRepl(&in)
}
