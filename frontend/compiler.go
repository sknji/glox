package frontend

import (
	"fmt"
	"github.com/urijn/glox/chunk"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/shared"
	"github.com/urijn/glox/value"
	"math"
)

var gCompiler = NewGlobalCompiler()

type GlobalCompiler struct {
	current *Compiler
}

func setCurrentCompiler(curr *Compiler) {
	gCompiler.current = curr
}

func currCompiler() *Compiler {
	return gCompiler.current
}

func NewGlobalCompiler() *GlobalCompiler {
	return &GlobalCompiler{}
}

// A compiler has roughly two jobs. It parses the user’s source code to understand
// what it means. Then it takes that knowledge and outputs low-level instructions
// that produce the same semantics. Many languages split those two roles into two
// separate passes in the implementation. A parser produces an AST and then a code
// generator traverses the AST and outputs target code.
type Compiler struct {
	parser     *Parser
	parseRules map[TokenType]*ParseRule
	scanner    *Scanner
	scope      *CompilerScope

	function *ObjectFunction
	typ      FunctionType
}

func NewCompiler(source []byte) *Compiler {
	compiler := &Compiler{
		scanner:  NewScanner(source),
		parser:   NewParser(),
		scope:    NewCompilerScope(),
		function: NewObjectFunction(TypeScript),
	}

	compiler.defineRules()

	setCurrentCompiler(compiler)

	local := currCompiler().scope.locals[currCompiler().scope.localCount]
	currCompiler().scope.localCount += 1
	local.depth = 0
	local.token = nil

	return compiler
}

func (c *Compiler) currChunk() *chunk.Chunk {
	return c.function.chunk
}

func (c *Compiler) Compile() (*ObjectFunction, bool) {
	c.advance()

	for !c.match(TokenEof) {
		c.declaration()
	}

	c.currChunk().EmitBytes(c.prevToken().Line, opcode.OpReturn)

	if shared.DebugPrintCode && !c.parser.hadError {


		c.currChunk().Disassemble(currCompiler().function.ParsedName())
	}

	return currCompiler().function, !c.parser.hadError
}

func (c *Compiler) match(tokenType TokenType) bool {
	if !c.check(tokenType) {
		return false
	}

	c.advance()

	return true
}

func (c *Compiler) check(tokenType TokenType) bool {
	return c.parser.current.Type.Is(tokenType)
}

func (c *Compiler) advance() {
	c.setPrevToken(c.currToken())

	for {
		c.setCurrToken(c.scanner.scanToken())
		//fmt.Printf("advance - curr: %+v, prev: %+v\n",
		//	c.currToken(), c.prevToken())
		if !c.parser.current.Type.Is(TokenError) {
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
	if c.parser.current.Type.Is(tokType) {
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

	if token.Type.Is(TokenEof) {
		fmt.Printf(" at end")
	} else if token.Type.Is(TokenError) {
		// Do nothing
	} else {
		fmt.Printf(" at '%s'", token.Val)
	}

	fmt.Printf(": %s\n", msg)
	c.parser.hadError = true
}

func (c *Compiler) makeConstant(value *value.Value) uint8 {
	constant := c.currChunk().AddConstant(value)
	if constant > math.MaxUint8 {
		c.error("Too many constants in one chunk.")
		return 0
	}

	return constant
}

func (c *Compiler) synchronize() {
	c.parser.panicMode = false

	for !c.parser.current.Type.Is(TokenEof) {
		if c.parser.previous.Type.Is(TokenSemicolon) {
			return
		}

		switch c.parser.current.Type {
		case TokenClass, TokenFun, TokenVar,
			TokenFor, TokenIf, TokenWhile, TokenPrint,
			TokenReturn:
			return
		}

		c.advance()
	}
}
