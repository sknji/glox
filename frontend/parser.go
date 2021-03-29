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
	} else if c.match(TokenIf) {
		c.ifStatement()
	} else if c.match(TokenWhile) {
		c.whileStatement()
	} else if c.match(TokenLeftBrace) {
		c.beginScope()
		c.block()
		c.endScope()
	} else {
		c.expressionStatement()
	}
}

func (c *Compiler) expressionStatement() {
	c.expression()
	c.consume(TokenSemicolon, "Expect ';' after expression.")
	c.emitBytes(opcode.OpPop)
}

func (c *Compiler) ifStatement() {
	c.consume(TokenLeftParen, "Expect '(' after 'if'.")
	c.expression()
	c.consume(TokenRightParen, "Expect ')' after condition.")

	thenJump := c.emitJump(opcode.OpJumpIfFalse)
	c.emitBytes(opcode.OpPop)

	c.statement()

	elseJump := c.emitJump(opcode.OpJump)
	c.emitBytes(opcode.OpPop)

	c.patchJump(thenJump)

	if c.match(TokenElse) {
		c.statement()
	}

	c.patchJump(elseJump)
}

func (c *Compiler) expression() {
	c.parsePrecedence(precAssignment)
}

func (c *Compiler) beginScope() {
	c.scope.scopeDepth += 1
}

func (c *Compiler) block() {
	for ; !c.check(TokenRightBrace) && !c.check(TokenEof); {
		c.declaration()
	}

	c.consume(TokenRightBrace, "Expect '}' after block.")
}

func (c *Compiler) endScope() {
	var popCount uint8 = 0
	for ; c.scope.localCount > 0 &&
		c.scope.locals[c.scope.localCount-1].depth > c.scope.scopeDepth; {
		popCount += 1
		c.scope.localCount -= 1
	}

	c.emitBytes(opcode.OpPopN, popCount)

	c.scope.scopeDepth -= 1
}

func (c *Compiler) string(bool) {
	c.emitConstant(value.NewObjectValueString(c.parser.previous.Val))
}

func (c *Compiler) variable(canAssign bool) {
	c.namedVariable(c.parser.previous, canAssign)
}

func (c *Compiler) namedVariable(tok *Token, canAssign bool) {
	var getOp, setOp uint8

	arg, ok := c.resolveLocal(tok)
	if ok {
		getOp = opcode.OpGetLocal
		setOp = opcode.OpSetLocal
	} else {
		arg = c.identifierConstant(tok)
		getOp = opcode.OpGetGlobal
		setOp = opcode.OpSetGlobal
	}

	if canAssign && c.match(TokenEqual) {
		c.expression()
		c.emitBytes(setOp, arg)
	} else {
		c.emitBytes(getOp, arg)
	}
}

func (c *Compiler) number(bool) {
	n, err := strconv.ParseFloat(c.parser.previous.Val, 64)
	if err != nil {
		c.error("Invalid number: " + err.Error())
	}

	c.emitConstant(value.NewValue(value.ValNumber, n))
}

func (c *Compiler) or_(bool) {
	elseJump := c.emitJump(opcode.OpJumpIfFalse)
	endJump := c.emitJump(opcode.OpJump)

	c.patchJump(elseJump)
	c.emitBytes(opcode.OpPop)

	c.parsePrecedence(precOr)
	c.patchJump(endJump)
}

func (c *Compiler) grouping(bool) {
	c.expression()
	c.consume(TokenRightParen, "Expect ')' after expression.")
}

func (c *Compiler) unary(bool) {
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

func (c *Compiler) binary(bool) {
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

func (c *Compiler) literal(bool) {
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

func (c *Compiler) whileStatement() {

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

	c.declareVariable()

	if c.scope.scopeDepth > 0 {
		return 0
	}

	return c.identifierConstant(c.parser.previous)
}

func (c *Compiler) identifiersEqual(a, b *Token) bool {
	return a.Val == b.Val
}

func (c *Compiler) resolveLocal(tok *Token) (uint8, bool) {
	for i := c.scope.localCount - 1; i >= 0; i-- {
		local := c.scope.locals[i]
		if local != nil && c.identifiersEqual(tok, local.token) {
			if local.depth == -1 {
				c.error("Can't read local variable in its own initializer.")
			}
			return uint8(i), true
		}
	}

	return 0, false
}

func (c *Compiler) identifierConstant(tok *Token) uint8 {
	return c.makeConstant(value.NewObjectValueString(tok.Val))
}

func (c *Compiler) addLocal(tok *Token) {
	err := c.scope.addLocal(tok)
	if err != nil {
		c.error(err.Error())
	}
}

func (c *Compiler) declareVariable() {
	if c.scope.scopeDepth == 0 {
		return
	}

	tok := c.parser.previous

	for i := c.scope.localCount - 1; i >= 0; i-- {
		local := c.scope.locals[i]
		if local.depth != -1 && local.depth < c.scope.scopeDepth {
			break
		}

		if c.identifiersEqual(tok, local.token) {
			c.error("Already variable with this name in this scope.")
		}
	}

	c.addLocal(tok)
}

func (c *Compiler) defineVariable(global uint8) {
	if c.scope.scopeDepth > 0 {
		c.markInitialized()
		return
	}

	c.emitBytes(opcode.OpDefineGlobal, global)
}

func (c *Compiler) markInitialized() {
	c.scope.locals[c.scope.localCount-1].depth = c.scope.scopeDepth
}
