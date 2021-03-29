package frontend

import (
	"fmt"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
	"math"
)

func (c *Compiler) emitBytes(op ...byte) {
	c.chunk.EmitBytes(c.prevToken().Line, op...)
}

func (c *Compiler) emitJump(op byte) int {
	c.emitBytes(op, 0xff, 0xff)
	return c.chunk.Count - 2
}

func (c *Compiler) emitConstant(value *value.Value) {
	c.emitBytes(opcode.OpConstant, c.makeConstant(value))
}

func (c *Compiler) patchJump(offset int) {
	// -2 to adjust for the bytecode for the jump offset itself.
	jump := c.chunk.Count - offset - 2

	if jump > math.MaxUint16 {
		c.error("Too much code to jump over.")
	}

	fmt.Printf("Jump: %d, offset: %d, current:%d\n", jump, offset, c.chunk.Count)
	c.chunk.Code[offset] = byte((jump >> 8) & 0xff)
	c.chunk.Code[offset + 1] = byte(jump & 0xff)
}
