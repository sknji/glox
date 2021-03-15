package lib

import "fmt"

const StackMax = 512

type InterpretResult int

const (
	InterpretOk InterpretResult = iota
	InterpretCompileError
	InterpretRuntimeError
)

type VM struct {
	chunk *Chunk

	// The IP always points to the next instruction, not the one
	// currently being handled.
	ip int

	stack [StackMax]Value

	// Stack pointer pointing at the array element just past the element
	// containing the top value on the stack.
	// Always points just past the last item. Points to where the next value
	// to be pushed will go.
	sp int
}

func NewVM() *VM {
	return &VM{
		sp:    0,
		stack: [StackMax]Value{},
		ip:    0,
		chunk: nil,
	}
}

// Run is the most performance critical part of the entire VM.
func (vm *VM) Run() InterpretResult {
	for ; ; {
		if DebugTraceExecution {
			vm.debug()
		}

		// Given a numeric opcode, we need to get to the right code that implements
		// that instruction's semantics. This process is called "decoding" or
		// "dispatching" the instruction
		instr := vm.readByte()

		// We have a single giant switch statement with a case for each opcode.
		// The body of each case implements that opcode’s behavior.
		switch instr {
		case OpReturn:
			val := vm.Pop()
			val.Print()
			fmt.Println()
			return InterpretOk
		case OpConstant:
			vm.Push(vm.readConstant())
		case OpNegate:
			vm.Push(-vm.Pop())
		case OpAdd:
			vm.binaryOperation(BinaryOpAdd)
		case OpSubtract:
			vm.binaryOperation(BinaryOpSubtract)
		case OpMutliply:
			vm.binaryOperation(BinaryOpMultiply)
		case OpDivide:
			vm.binaryOperation(BinaryOpDivide)
		}
	}
}

func (vm *VM) Interpret(chunk *Chunk) InterpretResult {
	vm.chunk = chunk
	vm.ip = 0
	return vm.Run()
}

func (vm *VM) Push(val Value) {
	vm.stack[vm.sp] = val
	vm.sp += 1
}

func (vm *VM) resetStack() {
	vm.sp = 0
}

func (vm *VM) Pop() Value {
	vm.sp -= 1
	return vm.stack[vm.sp]
}

func (vm *VM) debug() {
	fmt.Printf("          ")
	for _, slot := range vm.stack {
		fmt.Printf("[ ")
		slot.Print()
		fmt.Printf(" ]")
	}
	fmt.Printf("\n")

	vm.chunk.disassembleInstruction(vm.ip)
}

func (vm *VM) Free() {
	vm.sp = 0
	// TODO:
}
