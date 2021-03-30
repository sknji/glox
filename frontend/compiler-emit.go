package frontend

import (
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
	"math"
)

func (c *Compiler) emitBytes(op ...byte) {
	c.currChunk().EmitBytes(c.prevToken().Line, op...)
}

func (c *Compiler) emitJump(op byte) int {
	c.emitBytes(op, 0xff, 0xff)
	return c.currChunk().Count - 2
}

func (c *Compiler) emitConstant(value *value.Value) {
	c.emitBytes(opcode.OpConstant, c.makeConstant(value))
}

func (c *Compiler) patchJump(offset int) {
	// -2 to adjust for the bytecode for the jump offset itself.
	jump := c.currChunk().Count - offset - 2

	if jump > math.MaxUint16 {
		c.error("Too much code to jump over.")
	}

	c.currChunk().Code[offset] = byte((jump >> 8) & 0xff)
	c.currChunk().Code[offset+1] = byte(jump & 0xff)
}

func (c *Compiler) emitLoop(loopStart int) {
	c.emitBytes(opcode.OpLoop)

	offset := c.currChunk().Count - loopStart + 2
	if offset > math.MaxUint16 {
		c.error("Loop body too large.")
	}

	c.emitBytes(byte(offset>>8&0xff), byte(offset&0xff))
}
