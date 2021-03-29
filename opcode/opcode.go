package opcode

const (
	// Load the constant for use
	OpReturn byte = iota + 1
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
	OpPop
	OpPopN
	OpPrint
	OpDefineGlobal
	OpGetGlobal
	OpSetGlobal
	OpGetLocal
	OpSetLocal
	OpJumpIfFalse
	OpJump
	OpLoop
)
