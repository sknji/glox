package frontend

import (
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
	"strconv"
)

type Precedence uint

const (
	precNone Precedence = iota + 1
	precAssignment
	precOr
	precAnd
	precEquality
	precComparison
	precTerm
	precFactor
	precUnary
	precCall
	precPrimary
)

type Parser struct {
	current   *Token
	previous  *Token
	hadError  bool
	panicMode bool
}

func NewParser() *Parser {
	return &Parser{}
}

func (c *Compiler) expression() {
	c.parsePrecedence(precAssignment)
}

func (c *Compiler) number() {
	n, err := strconv.ParseFloat(c.parser.previous.Val, 64)
	if err != nil {
		c.error("Invalid number: " + err.Error())
	}

	c.emitConstant(value.NewValue(value.ValNumber, n))
}

func (c *Compiler) grouping() {
	c.expression()
	c.consume(TokenRightParen, "Expect ')' after expression.")
}

func (c *Compiler) unary() {
	tokType := c.parser.previous.Type

	c.parsePrecedence(precUnary)

	switch tokType {
	case TokenMinus:
		c.emitBytes(opcode.OpNegate)
	case TokenBang:
		c.emitBytes(opcode.OpNot)
	default:
		return
	}
}

func (c *Compiler) binary() {
	tokType := c.parser.previous.Type

	rule := c.getRule(tokType)
	c.parsePrecedence(rule.precedence + 1)

	switch tokType {
	case TokenPlus:
		c.emitBytes(opcode.OpAdd)
	case TokenMinus:
		c.emitBytes(opcode.OpSubtract)
	case TokenStar:
		c.emitBytes(opcode.OpMultiply)
	case TokenSlash:
		c.emitBytes(opcode.OpDivide)
	default:
		return
	}
}

func (c *Compiler) literal() {
	tok := c.prevToken()
	switch tok.Type {
	case TokenFalse:
		c.emitBytes(opcode.OpFalse)
	case TokenNil:
		c.emitBytes(opcode.OpNil)
	case TokenTrue:
		c.emitBytes(opcode.OpTrue)
	}
}

func (c *Compiler) parsePrecedence(precedence Precedence) {
	c.advance()
	//fmt.Printf("parsePrecedence - %+v\n",c.prevToken())
	prefix := c.getRule(c.prevToken().Type).prefix
	if prefix == nil {
		c.error("Expect expression.")
		return
	}

	prefix()

	//fmt.Printf("parsePrecedence infix - %+v\n",c.currToken())
	for precedence <= c.getRule(c.currToken().Type).precedence {
		c.advance()
		//fmt.Printf("parsePrecedence infix ok - %+v\n",c.prevToken())
		infix := c.getRule(c.prevToken().Type).infix
		infix()
	}
}
