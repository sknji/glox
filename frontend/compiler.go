package frontend

import (
	"fmt"
	chunk2 "github.com/urijn/glox/chunk"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/shared"
	"github.com/urijn/glox/value"
	"math"
)

// A compiler has roughly two jobs. It parses the user’s source code to understand
// what it means. Then it takes that knowledge and outputs low-level instructions
// that produce the same semantics. Many languages split those two roles into two
// separate passes in the implementation. A parser produces an AST and then a code
// generator traverses the AST and outputs target code.
type Compiler struct {
	parser     *Parser
	parseRules map[TokenType]*ParseRule
	scanner    *Scanner
	chunk      *chunk2.Chunk
}

func NewCompiler(source []byte) *Compiler {
	compiler := &Compiler{
		scanner: NewScanner(source),
		parser:  NewParser(),
		chunk:   chunk2.NewChunk(),
	}

	compiler.defineRules()

	return compiler
}

func (c *Compiler) Compile() (*chunk2.Chunk, bool) {
	c.advance()
	c.expression()
	c.consume(TokenEof, "Expect end of expression.")
	c.chunk.EmitBytes(c.prevToken().Line, opcode.OpReturn)

	if shared.DebugPrintCode && !c.parser.hadError {
		c.chunk.Disassemble("code")
	}

	return c.chunk, !c.parser.hadError
}

func (c *Compiler) advance() {
	c.setPrevToken(c.currToken())

	for {
		c.setCurrToken(c.scanner.scanToken())
		if c.parser.current.Type != TokenError {
			break
		}

		c.errorAtCurrent(c.parser.current.Val)
	}
}

func (c *Compiler) setPrevToken(tok *Token) {
	c.parser.previous = tok
}

func (c *Compiler) setCurrToken(tok *Token) {
	c.parser.current = tok
}

func (c *Compiler) prevToken() *Token {
	return c.parser.previous
}

func (c *Compiler) currToken() *Token {
	return c.parser.current
}

func (c *Compiler) consume(tokType TokenType, msg string) {
	if c.parser.current.Type == tokType {
		c.advance()
		return
	}

	c.errorAtCurrent(msg)
}

func (c *Compiler) errorAtCurrent(msg string) {
	c.errorAt(c.parser.current, msg)
}

func (c *Compiler) error(msg string) {
	c.errorAt(c.parser.previous, msg)
}

func (c *Compiler) errorAt(token *Token, msg string) {
	if c.parser.panicMode {
		return
	}
	c.parser.panicMode = true

	fmt.Printf("[line %d] Error", token.Line)

	if token.Type == TokenEof {
		fmt.Printf(" at end")
	} else if token.Type == TokenError {
		// Do nothing
	} else {
		fmt.Printf(" at '%s'", token.Val)
	}

	fmt.Printf(": %s\n", msg)
	c.parser.hadError = true
}

func (c *Compiler) makeConstant(value value.Value) uint8 {
	constant := c.chunk.AddConstant(value)
	if constant > math.MaxUint8 {
		c.error("Too many constants in one chunk.")
		return 0
	}

	return constant
}
