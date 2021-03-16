package frontend

import "fmt"

type Compiler struct {
	scanner *Scanner
}

func NewCompiler() *Compiler {
	return &Compiler{
		scanner: NewScanner(),
	}
}

func (c *Compiler) Compile(source []byte) error {
	line := -1
	for {
		token := c.scanner.scanToken()
		if token.Line != line {
			fmt.Printf("%4d ", token.Line)
			line = token.Line
		} else {
			fmt.Printf("   | ")
		}

		fmt.Printf("%2d '%s'\n", token.Type, token.Val)

		if token.Type == TokenEof {
			break
		}
	}
	//return nil
}
