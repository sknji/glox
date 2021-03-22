package chunk

import (
	"fmt"
	"github.com/urijn/glox/opcode"
)

func (c *Chunk) Disassemble(name string) {
	fmt.Printf("== %s ==\n", name)

	offset := 0
	for offset < c.Count {
		offset = c.DisassembleInstruction(offset)
	}
}

func (c *Chunk) DisassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.Lines[offset] == c.Lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.Lines[offset])
	}

	var instr = c.Code[offset]
	switch instr {
	case opcode.OpReturn:
		return c.SimpleInstruction("OP_RETURN", offset)
	case opcode.OpConstant:
		return c.ConstantInstruction("OP_CONSTANT", offset)
	case opcode.OpNegate:
		return c.SimpleInstruction("OP_NEGATE", offset)
	case opcode.OpAdd:
		return c.SimpleInstruction("OP_ADD", offset)
	case opcode.OpSubtract:
		return c.SimpleInstruction("OP_SUBTRACT", offset)
	case opcode.OpMultiply:
		return c.SimpleInstruction("OP_MULTIPLY", offset)
	case opcode.OpDivide:
		return c.SimpleInstruction("OP_DIVIDE", offset)
	case opcode.OpNil:
		return c.SimpleInstruction("OP_NIL", offset)
	case opcode.OpTrue:
		return c.SimpleInstruction("OP_TRUE", offset)
	case opcode.OpFalse:
		return c.SimpleInstruction("OP_FALSE", offset)
	case opcode.OpNot:
		return c.SimpleInstruction("OP_NOT", offset)
	case opcode.OpEqual:
		return c.SimpleInstruction("OP_EQUAL", offset)
	case opcode.OpGreater:
		return c.SimpleInstruction("OP_GREATER", offset)
	case opcode.OpLess:
		return c.SimpleInstruction("OP_LESS", offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instr)
		return offset + 1
	}
}

func (c *Chunk) SimpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func (c *Chunk) ConstantInstruction(name string, offset int) int {
	var constant = c.Code[offset+1]

	fmt.Printf("%-16s %4d '", name, constant)
	c.Constants.Values[constant].Print()
	fmt.Printf("'\n")
	return offset + 2
}
