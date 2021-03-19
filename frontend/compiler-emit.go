package frontend

import (
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
)

func (c *Compiler) emitBytes(bytes ...byte) {
	c.chunk.EmitBytes(c.prevToken().Line, bytes...)
}

func (c *Compiler) emitConstant(value value.Value) {
	c.emitBytes(opcode.OpConstant, c.makeConstant(value))
}
