package lib

type BinaryOp int

const (
	BinaryOpAdd BinaryOp = iota
	BinaryOpSubtract
	BinaryOpMultiply
	BinaryOpDivide
)

func (vm *VM) incrementIP() {
	vm.ip += 1
}

func (vm *VM) readByte() byte {
	defer vm.incrementIP()
	return vm.chunk.code[vm.ip]
}

func (vm *VM) readConstant() Value {
	return vm.chunk.constants.values[vm.readByte()]
}

func (vm *VM) binaryOperation(op BinaryOp) {
	b := vm.Pop()
	a := vm.Pop()

	var result Value
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
	vm.Push(result)
}
