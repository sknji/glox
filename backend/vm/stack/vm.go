package stack

import (
	"fmt"
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

	stack [StackMax]value.Value

	// Stack pointer pointing at the array element just past the element
	// containing the top value on the stack.
	// Always points just past the last item. Points to where the next value
	// to be pushed will go.
	sp int
}

func NewVM() *VM {
	return &VM{
		sp:    0,
		stack: [StackMax]value.Value{},
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
		instr := v.readByte()

		// We have a single giant switch statement with a case for each opcode.
		// The body of each case implements that opcode’s behavior.
		switch instr {
		case opcode.OpReturn:
			val := v.Pop()
			val.PrintLn()
			return vm.InterpretOk
		case opcode.OpConstant:
			v.Push(v.readConstant())
		case opcode.OpNegate:
			v.Push(-v.Pop())
		case opcode.OpAdd:
			v.binaryOperation(BinaryOpAdd)
		case opcode.OpSubtract:
			v.binaryOperation(BinaryOpSubtract)
		case opcode.OpMultiply:
			v.binaryOperation(BinaryOpMultiply)
		case opcode.OpDivide:
			v.binaryOperation(BinaryOpDivide)
		}
	}
}

func (v *VM) Interpret(chunk *chunk.Chunk) vm.InterpretResult {
	v.chunk = chunk
	v.ip = 0
	return v.Run()
}

func (v *VM) Push(val value.Value) {
	v.stack[v.sp] = val
	v.sp += 1
}

func (v *VM) resetStack() {
	v.sp = 0
}

func (v *VM) Pop() value.Value {
	v.sp -= 1
	return v.stack[v.sp]
}

func (v *VM) debug() {
	fmt.Printf("          ")
	for i, slot := range v.stack {
		fmt.Printf("[ ")
		fmt.Printf("%d: ", i)
		slot.Print()
		fmt.Printf(" ]")
	}
	fmt.Printf("\n")

	v.chunk.DisassembleInstruction(v.ip)
}

func (v *VM) Free() {
	v.sp = 0
	// TODO:
}
