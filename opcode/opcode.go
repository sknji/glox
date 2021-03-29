package opcode

type OpCode byte

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

var opCodeToString = map[byte]string{
	OpReturn:       "OpReturn",
	OpConstant:     "OpConstant",
	OpNegate:       "OpNegate",
	OpAdd:          "OpAdd",
	OpSubtract:     "OpSubtract",
	OpMultiply:     "OpMultiply",
	OpDivide:       "OpDivide",
	OpNil:          "OpNil",
	OpTrue:         "OpTrue",
	OpFalse:        "OpFalse",
	OpNot:          "OpNot",
	OpEqual:        "OpEqual",
	OpGreater:      "OpGreater",
	OpLess:         "OpLess",
	OpPop:          "OpPop",
	OpPopN:         "OpPopN",
	OpPrint:        "OpPrint",
	OpDefineGlobal: "OpDefineGlobal",
	OpGetGlobal:    "OpGetGlobal",
	OpSetGlobal:    "OpSetGlobal",
	OpGetLocal:     "OpGetLocal",
	OpSetLocal:     "OpSetLocal",
	OpJumpIfFalse:  "OpJumpIfFalse",
	OpJump:         "OpJump",
	OpLoop:         "OpLoop",
}

func OpCodeToString(op byte) string {
	str, ok := opCodeToString[op]
	if !ok {
		return ""
	}

	return str
}
