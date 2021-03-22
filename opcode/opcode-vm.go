package opcode

type OpCode byte

const (
	// Load the constant for use
	OpReturn OpCode = iota + 1
	OpConstant
	OpNegate
	OpAdd
	OpSubtract
	OpMultiply
	OpDivide
	OpNil
	OpTrue
	OpFalse
	OpNot
	OpEqual
	OpGreater
	OpLess
)
