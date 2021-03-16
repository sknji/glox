package stack

import (
	"fmt"
	"github.com/urijn/glox/backend/vm"
)

const StackMax = 512

type VM struct {
	chunk *vm.Chunk

	// The IP always points to the next instruction, not the one
	// currently being handled.
	ip int

	stack [StackMax]vm.Value

	// Stack pointer pointing at the array element just past the element
	// containing the top value on the stack.
	// Always points just past the last item. Points to where the next value
	// to be pushed will go.
	sp int
}

func NewVM() *VM {
	return &VM{
		sp:    0,
		stack: [StackMax]vm.Value{},
		ip:    0,
		chunk: nil,
	}
}

// Run is the most performance critical part of the entire VM.
func (v *VM) Run() vm.InterpretResult {
	for {
		if vm.DebugTraceExecution {
			v.debug()
		}

		// Given a numeric opcode, we need to get to the right code that implements
		// that instruction's semantics. This process is called "decoding" or
		// "dispatching" the instruction
		instr := v.readByte()

		// We have a single giant switch statement with a case for each opcode.
		// The body of each case implements that opcode’s behavior.
		switch instr {
		case vm.OpReturn:
			val := v.Pop()
			val.Print()
			fmt.Println()
			return vm.InterpretOk
		case vm.OpConstant:
			v.Push(v.readConstant())
		case vm.OpNegate:
			v.Push(-v.Pop())
		case vm.OpAdd:
			v.binaryOperation(BinaryOpAdd)
		case vm.OpSubtract:
			v.binaryOperation(BinaryOpSubtract)
		case vm.OpMultiply:
			v.binaryOperation(BinaryOpMultiply)
		case vm.OpDivide:
			v.binaryOperation(BinaryOpDivide)
		}
	}
}

func (v *VM) Interpret(chunk *vm.Chunk) vm.InterpretResult {
	v.chunk = chunk
	v.ip = 0
	return v.Run()
}

func (v *VM) Push(val vm.Value) {
	v.stack[v.sp] = val
	v.sp += 1
}

func (v *VM) resetStack() {
	v.sp = 0
}

func (v *VM) Pop() vm.Value {
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
