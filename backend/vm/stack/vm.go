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

	// Stack pointer pointing at the array element just past the element
	// containing the top value on the stack.
	// Always points just past the last item. Points to where the next value
	// to be pushed will go.
	sp int
}

func NewVM() *VM {
	return &VM{
		sp:    0,
		stack: [StackMax]*value.Value{},
		ip:    0,
		chunk: nil,
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
		instr := opcode.OpCode(v.readByte())

		// We have a single giant switch statement with a case for each opcode.
		// The body of each case implements that opcode’s behavior.
		switch instr {
		case opcode.OpReturn:
			v.Pop().Println()
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
			v.Push(value.NewValue(value.ValNumber, -val.Val.GetAsNumber()))
		case opcode.OpAdd,
			opcode.OpSubtract,
			opcode.OpMultiply,
			opcode.OpDivide:
			r := v.binaryOperation(instr)
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

func (v *VM) Pop() *value.Value {
	v.sp -= 1
	return v.stack[v.sp]
}

func (v *VM) Peek(distance int) *value.Value {
	return v.stack[v.sp-1-distance]
}

func (v *VM) resetStack() {
	v.sp = 0
}

func (v *VM) Free() {
	v.sp = 0
	// TODO:
}
