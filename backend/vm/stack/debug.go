package stack

import (
	"fmt"
)

const DebugTraceExecution = true

func (c *Chunk) Disassemble(name string) {
	fmt.Printf("== %s ==\n", name)

	var offset = 0
	for offset < c.count {
		offset += c.disassembleInstruction(offset)
	}
}

func (c *Chunk) disassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.lines[offset] == c.lines[offset - 1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.lines[offset])
	}

	var instr = c.code[offset]
	switch instr {
	case OpReturn:
		return c.simpleInstruction("OP_RETURN", offset)
	case OpConstant:
		return c.constantInstruction("OP_CONSTANT", offset)
	case OpNegate:
		return c.simpleInstruction("OP_NEGATE", offset)
	case OpAdd:
		return c.simpleInstruction("OP_ADD", offset)
	case OpSubtract:
		return c.simpleInstruction("OP_SUBTRACT", offset)
	case OpMultiply:
		return c.simpleInstruction("OP_MULTIPLY", offset)
	case OpDivide:
		return c.simpleInstruction("OP_DIVIDE", offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instr)
		return offset + 1
	}
}

func (c *Chunk) simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func (c *Chunk) constantInstruction(name string, offset int) int {
	var constant = c.code[offset + 1]

	fmt.Printf("%-16s %4d '", name, constant)
	c.constants.values[constant].Print()
	fmt.Printf("'\n")
	return offset + 2
}
