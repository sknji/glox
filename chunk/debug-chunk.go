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
	case opcode.OpPop:
		return c.SimpleInstruction("OP_POP", offset)
	case opcode.OpPopN:
		return c.ByteInstruction("OP_POP_N", offset)
	case opcode.OpPrint:
		return c.SimpleInstruction("OP_PRINT", offset)
	case opcode.OpDefineGlobal:
		return c.ConstantInstruction("OP_DEFINE_GLOBAL", offset)
	case opcode.OpGetGlobal:
		return c.ConstantInstruction("OP_GET_GLOBAL", offset)
	case opcode.OpSetGlobal:
		return c.ConstantInstruction("OP_SET_GLOBAL", offset)
	case opcode.OpGetLocal:
		return c.ByteInstruction("OP_GET_LOCAL", offset)
	case opcode.OpSetLocal:
		return c.ByteInstruction("OP_SET_LOCAL", offset)
	case opcode.OpJump:
		return c.jumpInstruction("OP_JUMP", 1, offset)
	case opcode.OpJumpIfFalse:
		return c.jumpInstruction("OP_JUMP_IF_FALSE", 1, offset)
	case opcode.OpLoop:
		return c.jumpInstruction("OP_LOOP", -1, offset)
	default:
		fmt.Printf("Unknown opcode %s\n", opcode.OpCodeToString(instr))
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
	c.Constants.Values[constant].Println()
	return offset + 2
}

func (c *Chunk) ByteInstruction(name string, offset int) int {
	slot := c.Code[offset+1]
	fmt.Printf("%-16s %4d\n", name, slot)
	return offset + 2
}

func (c *Chunk) jumpInstruction(name string, sign, offset int) int {
	jump := uint16(c.Code[offset+1]) << 8
	jump |= uint16(c.Code[offset+2])
	fmt.Printf("%-16s %4d -> %d\n", name, offset,
		offset+3+sign*int(jump))

	return offset + 3
}
