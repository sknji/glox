package vm

import (
	"fmt"
)

const DebugTraceExecution = true

func (c *Chunk) Disassemble(name string) {
	fmt.Printf("== %s ==\n", name)

	var offset = 0
	for offset < c.Count {
		offset += c.DisassembleInstruction(offset)
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
	case OpReturn:
		return c.SimpleInstruction("OP_RETURN", offset)
	case OpConstant:
		return c.ConstantInstruction("OP_CONSTANT", offset)
	case OpNegate:
		return c.SimpleInstruction("OP_NEGATE", offset)
	case OpAdd:
		return c.SimpleInstruction("OP_ADD", offset)
	case OpSubtract:
		return c.SimpleInstruction("OP_SUBTRACT", offset)
	case OpMultiply:
		return c.SimpleInstruction("OP_MULTIPLY", offset)
	case OpDivide:
		return c.SimpleInstruction("OP_DIVIDE", offset)
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
