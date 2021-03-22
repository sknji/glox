package stack

import (
	"github.com/urijn/glox/backend/vm"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
)

func (v *VM) incrementIP() {
	v.ip += 1
}

func (v *VM) readByte() byte {
	defer v.incrementIP()
	return v.chunk.Code[v.ip]
}

func (v *VM) readConstant() *value.Value {
	return v.chunk.Constants.Values[v.readByte()]
}

func (v *VM) binaryOperation(op opcode.OpCode) vm.InterpretResult {
	b := v.Peek(0)
	a := v.Peek(1)

	if !b.Is(value.ValNumber) || !a.Is(value.ValNumber) {
		v.runtimeError("Operands must be numbers.")
		return vm.InterpretRuntimeError
	}

	bNum := v.Pop().Val.GetAsNumber()
	aNum := v.Pop().Val.GetAsNumber()

	var result float64
	switch op {
	case opcode.OpAdd:
		result = aNum + bNum
	case opcode.OpDivide:
		result = aNum / bNum
	case opcode.OpMultiply:
		result = aNum * bNum
	case opcode.OpSubtract:
		result = aNum - bNum
	}

	v.Push(value.NewValue(value.ValNumber, result))

	return vm.InterpretOk
}

func (v *VM) isFalsey(val *value.Value) bool {
	return val.Is(value.ValNil) ||
		(val.Is(value.ValBool) && !val.Val.GetAsBool())
}
