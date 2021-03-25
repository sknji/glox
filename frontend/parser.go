package frontend

import (
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
	"strconv"
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

func (c *Compiler) declaration() {
	if c.match(TokenVar) {
		c.varDeclaration()
	} else {
		c.statement()
	}

	if c.parser.panicMode {
		c.synchronize()
	}
}

func (c *Compiler) varDeclaration() {
	global := c.parseVariable("Expect variable name")

	if c.match(TokenEqual) {
		c.expression()
	} else {
		c.emitBytes(opcode.OpNil)
	}

	c.consume(TokenSemicolon, "Expect ';' after variable declaration.")

	c.defineVariable(global)
}

func (c *Compiler) statement() {
	if c.match(TokenPrint) {
		c.printStatement()
	} else {
		c.expressionStatement()
	}
}

func (c *Compiler) expressionStatement() {
	c.expression()
	c.consume(TokenSemicolon, "Expect ';' after expression.")
	c.emitBytes(opcode.OpPop)
}

func (c *Compiler) expression() {
	c.parsePrecedence(precAssignment)
}

func (c *Compiler) string(canAssign bool) {
	c.emitConstant(value.NewObjectValueString(c.parser.previous.Val))
}

func (c *Compiler) variable(canAssign bool) {
	c.namedVariable(c.parser.previous, canAssign)
}

func (c *Compiler) namedVariable(tok *Token, canAssign bool) {
	arg := c.identifierConstant(tok)
	if canAssign && c.match(TokenEqual) {
		c.expression()
		c.emitBytes(opcode.OpSetGlobal, arg)
	} else {
		c.emitBytes(opcode.OpGetGlobal, arg)
	}
}

func (c *Compiler) number(canAssign bool) {
	n, err := strconv.ParseFloat(c.parser.previous.Val, 64)
	if err != nil {
		c.error("Invalid number: " + err.Error())
	}

	c.emitConstant(value.NewValue(value.ValNumber, n))
}

func (c *Compiler) grouping(canAssign bool) {
	c.expression()
	c.consume(TokenRightParen, "Expect ')' after expression.")
}

func (c *Compiler) unary(canAssign bool) {
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

func (c *Compiler) binary(canAssign bool) {
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
	case TokenBangEqual:
		c.emitBytes(opcode.OpEqual, opcode.OpNot)
	case TokenEqualEqual:
		c.emitBytes(opcode.OpEqual)
	case TokenGreater:
		c.emitBytes(opcode.OpGreater)
	case TokenGreaterEqual:
		c.emitBytes(opcode.OpLess, opcode.OpNot)
	case TokenLess:
		c.emitBytes(opcode.OpLess)
	case TokenLessEqual:
		c.emitBytes(opcode.OpLess, opcode.OpNot)
	default:
		return
	}
}

func (c *Compiler) literal(canAssign bool) {
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

func (c *Compiler) printStatement() {
	c.expression()
	c.consume(TokenSemicolon, "Expect ';' after value.")
	c.emitBytes(opcode.OpPrint)
}

func (c *Compiler) parsePrecedence(precedence Precedence) {
	c.advance()
	//fmt.Printf("parsePrecedence - %+v\n",c.prevToken())
	prefix := c.getRule(c.prevToken().Type).prefix
	if prefix == nil {
		c.error("Expect expression.")
		return
	}

	canAssign := precedence <= precAssignment
	prefix(canAssign)

	//fmt.Printf("parsePrecedence infix - %+v\n",c.currToken())
	for precedence <= c.getRule(c.currToken().Type).precedence {
		c.advance()
		//fmt.Printf("parsePrecedence infix ok - %+v\n",c.prevToken())
		infix := c.getRule(c.prevToken().Type).infix
		infix(canAssign)
	}

	if canAssign && c.match(TokenEqual) {
		c.error("Invalid assignment target.")
	}
}

func (c *Compiler) parseVariable(errMsg string) uint8 {
	c.consume(TokenIdentifier, errMsg)
	return c.identifierConstant(c.parser.previous)
}

func (c *Compiler) identifierConstant(tok *Token) uint8 {
	return c.makeConstant(value.NewObjectValueString(tok.Val))
}

func (c *Compiler) defineVariable(global uint8)  {
	c.emitBytes(opcode.OpDefineGlobal, global)
}
