package stack

import "github.com/urijn/glox/backend/vm"

type BinaryOp int

const (
	BinaryOpAdd BinaryOp = iota
	BinaryOpSubtract
	BinaryOpMultiply
	BinaryOpDivide
)

func (v *VM) incrementIP() {
	v.ip += 1
}

func (v *VM) readByte() byte {
	defer v.incrementIP()
	return v.chunk.Code[v.ip]
}

func (v *VM) readConstant() vm.Value {
	return v.chunk.Constants.Values[v.readByte()]
}

func (v *VM) binaryOperation(op BinaryOp) {
	b := v.Pop()
	a := v.Pop()

	var result vm.Value
	switch op {
	case BinaryOpAdd:
		result = a + b
	case BinaryOpDivide:
		result = a / b
	case BinaryOpMultiply:
		result = a * b
	case BinaryOpSubtract:
		result = a - b
	}
	v.Push(result)
}
