package stack

import (
	"github.com/urijn/glox/backend/vm"
	"github.com/urijn/glox/chunk"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/shared"
	"github.com/urijn/glox/value"
)

const StackMax = 512

type VM struct {
	chunk *chunk.Chunk

	// The IP always points to the next instruction, not the one
	// currently being handled.
	ip int

	stack [StackMax]*value.Value

	globals map[string]*value.Value

	// Stack pointer pointing at the array element just past the element
	// containing the top value on the stack.
	// Always points just past the last item. Points to where the next value
	// to be pushed will go.
	sp int
}

func NewVM() *VM {
	return &VM{
		sp:      0,
		stack:   [StackMax]*value.Value{},
		globals: make(map[string]*value.Value),
		ip:      0,
		chunk:   nil,
	}
}

// Run is the most performance critical part of the entire VM.
func (v *VM) Run() vm.InterpretResult {
	for {
		if shared.DebugTraceExecution {
			v.debug()
		}

		// Given a numeric opcode, we need to get to the right code that implements
		// that instruction's semantics. This process is called "decoding" or
		// "dispatching" the instruction
		instr := v.readByte()

		// We have a single giant switch statement with a case for each opcode.
		// The body of each case implements that opcode’s behavior.
		switch instr {
		case opcode.OpReturn:
			return vm.InterpretOk
		case opcode.OpConstant:
			v.Push(v.readConstant())
		case opcode.OpNegate:
			val := v.Peek(0)
			if !val.Is(value.ValNumber) {
				v.runtimeError("Operand must be a number.")
				return vm.InterpretRuntimeError
			}
			val = v.Pop()
			v.Push(value.NewValue(value.ValNumber, -val.Val.GetNumber()))
		case
			opcode.OpSubtract,
			opcode.OpMultiply,
			opcode.OpDivide:
			r := v.binaryOperation(value.ValNumber, instr)
			if !r.IsSuccess() {
				return r
			}
		case opcode.OpAdd:
			b := v.Peek(0)
			a := v.Peek(1)

			if a.IsObjType(value.ObjString) && b.IsObjType(value.ObjString) {
				v.concatenate()
			} else if a.Is(value.ValNumber) && b.Is(value.ValNumber) {
				bVal := v.Pop().Val.GetNumber()
				aVal := v.Pop().Val.GetNumber()
				v.Push(value.NewValue(value.ValNumber, bVal+aVal))
			} else {
				v.runtimeError("Operands must be two numbers or two strings.")
				return vm.InterpretRuntimeError
			}
		case opcode.OpGreater,
			opcode.OpLess:
			r := v.binaryOperation(value.ValBool, instr)
			if !r.IsSuccess() {
				return r
			}
		case opcode.OpNil:
			v.Push(value.NewValue(value.ValNil, nil))
		case opcode.OpTrue:
			v.Push(value.NewValue(value.ValBool, true))
		case opcode.OpFalse:
			v.Push(value.NewValue(value.ValBool, false))
		case opcode.OpNot:
			val := v.Pop()
			v.Push(value.NewValue(value.ValBool, v.isFalsey(val)))
		case opcode.OpEqual:
			b, a := v.Pop(), v.Pop()
			v.Push(value.NewValue(value.ValBool, v.valuesEqual(a, b)))
		case opcode.OpPrint:
			v.Pop().Println()
		case opcode.OpPop:
			v.Pop()
		case opcode.OpPopN:
			v.PopN(int(v.readByte()))
		case opcode.OpDefineGlobal:
			name := v.readConstant().Val.GetObject().(*value.ObjectString)
			v.globals[name.String()] = v.Peek(0)
			v.Pop()
		case opcode.OpGetGlobal:
			name := v.readConstant().Val.GetObject().(*value.ObjectString)
			val, ok := v.globals[name.String()]
			if !ok {
				v.runtimeError("Undefined variable '%s'.", name)
				return vm.InterpretRuntimeError
			}
			v.Push(val)
		case opcode.OpSetGlobal:
			name := v.readConstant().Val.GetObject().(*value.ObjectString)
			if _, ok := v.globals[name.String()]; !ok {
				delete(v.globals, name.String())
				v.runtimeError("Undefined variable '%s'.", name)
				return vm.InterpretRuntimeError
			}

			v.globals[name.String()] = v.Peek(0)
		case opcode.OpSetLocal:
			v.stack[v.readByte()] = v.Peek(0)
		case opcode.OpGetLocal:
			v.Push(v.stack[v.readByte()])

		}
	}
}

func (v *VM) Interpret(chunk *chunk.Chunk) vm.InterpretResult {
	v.chunk = chunk
	v.ip = 0
	return v.Run()
}

func (v *VM) Push(val *value.Value) {
	v.stack[v.sp] = val
	v.sp += 1
}

func (v *VM) PopN(n int) *value.Value {
	v.sp -= n
	return v.stack[v.sp]
}

func (v *VM) Pop() *value.Value {
	return v.PopN(1)
}

func (v *VM) Peek(distance int) *value.Value {
	return v.stack[v.sp-1-distance]
}

func (v *VM) resetStack() {
	v.sp = 0
}

func (v *VM) Free() {
	v.sp = 0
	v.globals = nil
	// TODO:
}
