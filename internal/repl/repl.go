package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jndaf31/term-in-algebra/internal/matrix"
	"github.com/jndaf31/term-in-algebra/internal/parse"
)

type Inputs struct {
	MatrixInputs []matrix.Matrix
	Operation    string
}

func StartRepl(in *Inputs) {
	var inputBuff []string
	idx := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for idx < len(in.MatrixInputs) && scanner.Scan() {
		fmt.Print("> ")

		cIn, err := parse.CleanInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		mat, err := parse.BuildMatrix(cIn, &inputBuff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if mat == nil {
			continue
		}
		in.MatrixInputs[idx].Data = mat
		fmt.Printf("Current input1: %v \nCurrent input2: %v \n",
			in.MatrixInputs[0].Data,
			in.MatrixInputs[1].Data)
		idx++
	}
}
