package vm

type InterpretResult int

const (
	InterpretOk InterpretResult = iota
	InterpretCompileError
	InterpretRuntimeError
)

type VirtualMachine interface {
	Run() InterpretResult
	Interpret(chunk *Chunk) InterpretResult
	Free()
}
