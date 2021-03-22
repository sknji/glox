package chunk

import "github.com/urijn/glox/opcode"

func (c *Chunk) EmitBytes(line uint, op opcode.OpCode, operands ...byte) {
	c.Write(byte(op), line)
	for _, b := range operands {
		c.Write(b, line)
	}
}
