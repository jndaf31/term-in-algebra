package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	fmt.Println("Welcome to term-in-algebra!")
	fmt.Println("Accepts any whole, fractional or dot separated decimal as input.")
	fmt.Println("Row separator:                 |   Matrix END separator:   ///")
	fmt.Println("Print current input buffer:    ?    Clear input buffer:    ?clear")
	fmt.Print("> ")
	for idx < len(in.MatrixInputs) && scanner.Scan() {

		// Raw processing, for commands
		cmd := strings.TrimSpace(string(scanner.Bytes()))

		switch cmd {
		case "?":
			fmt.Println("input1: ", in.MatrixInputs[0].GetData())
			fmt.Println("input2: ", in.MatrixInputs[1].GetData())
			fmt.Println("input buffer: ", inputBuff)
			fmt.Print("> ")
			continue

		case "?clear":
			inputBuff = []string{}
			fmt.Print("> ")
			continue

		default:
		}

		// Text processing, for input
		cIn, err := parse.CleanInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			fmt.Print("> ")
			continue
		}
		if len(cIn) == 0 {
			fmt.Print("> ")
			continue
		}

		mat, err := parse.BuildMatrix(cIn, &inputBuff)
		if err != nil {
			fmt.Println(err)
			fmt.Print("> ")
			continue
		}

		if mat.Rows() == 0 {
			fmt.Print("> ")
			continue
		}
		in.MatrixInputs[idx] = mat

		fmt.Printf("Current input1: %v \nCurrent input2: %v \n",
			in.MatrixInputs[0].GetData(),
			in.MatrixInputs[1].GetData())
		fmt.Print("> ")
		idx++
	}
}
