package frontend

import (
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
)

func (c *Compiler) emitBytes(op opcode.OpCode, operands ...byte) {
	c.chunk.EmitBytes(c.prevToken().Line, op, operands...)
}

func (c *Compiler) emitConstant(value *value.Value) {
	c.emitBytes(opcode.OpConstant, c.makeConstant(value))
}
