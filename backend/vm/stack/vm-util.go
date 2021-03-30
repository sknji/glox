package stack

import (
	"bytes"
	"github.com/urijn/glox/backend/vm"
	"github.com/urijn/glox/chunk"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
)

func (v *VM) currFrame() *CallFrame {
	return v.frames[v.frameCount-1]
}

func (v *VM) currFrameChunk() *chunk.Chunk {
	return v.currFrame().function.Chunk()
}

func (v *VM) currFrameCode() []byte {
	return v.currFrameChunk().Code
}

func (v *VM) currFrameIp() int {
	return v.currFrame().ip
}

func (v *VM) incrementIP(count int) {
	v.currFrame().ip += count
}

func (v *VM) readByte() byte {
	v.incrementIP(1)
	return v.currFrameCode()[v.currFrameIp()-1]
}

func (v *VM) readShort() uint16 {
	v.incrementIP(2)
	return uint16(v.currFrameCode()[v.currFrameIp()-2])<<8 |
		uint16(v.currFrameCode()[v.currFrameIp()-1])
}

func (v *VM) readConstant() *value.Value {
	return v.currFrameChunk().Constants.Values[v.readByte()]
}

func (v *VM) binaryOperation(valType value.ValueType, op byte) vm.InterpretResult {
	b := v.Peek(0)
	a := v.Peek(1)

	if !b.Is(value.ValNumber) || !a.Is(value.ValNumber) {
		v.runtimeError("Operands must be numbers.")
		return vm.InterpretRuntimeError
	}

	bVal := v.Pop().Val.GetNumber()
	aVal := v.Pop().Val.GetNumber()

	var result interface{}
	switch op {
	case opcode.OpAdd:
		result = aVal + bVal
	case opcode.OpDivide:
		result = aVal / bVal
	case opcode.OpMultiply:
		result = aVal * bVal
	case opcode.OpSubtract:
		result = aVal - bVal
	case opcode.OpGreater:
		result = aVal > bVal
	case opcode.OpLess:
		result = aVal < bVal
	}

	v.Push(value.NewValue(valType, result))

	return vm.InterpretOk
}

func (v *VM) isFalsey(val *value.Value) bool {
	return val.Is(value.ValNil) ||
		(val.Is(value.ValBool) && !val.Val.GetBool())
}

func (v *VM) valuesEqual(a, b *value.Value) bool {
	if a.ValType != b.ValType {
		return false
	}

	switch a.ValType {
	case value.ValBool:
		return a.Val.GetBool() == b.Val.GetBool()
	case value.ValNil:
		return true
	case value.ValNumber:
		return a.Val.GetNumber() == a.Val.GetNumber()
	case value.ValObj:
		return a.Val.Get() == b.Val.Get()
	default:
		return false
	}
}

func (v *VM) concatenate() {
	bStr := (v.Pop().Val.GetObject()).(*value.ObjectString)
	aStr := v.Pop().Val.GetObject().(*value.ObjectString)

	var buffer bytes.Buffer
	buffer.WriteString(aStr.String())
	buffer.WriteString(bStr.String())
	v.Push(value.NewObjectValueString(buffer.String()))
}
