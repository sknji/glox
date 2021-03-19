package vm

import "github.com/urijn/glox/chunk"

type InterpretResult int

const (
	InterpretOk InterpretResult = iota + 1
	InterpretCompileError
	InterpretRuntimeError
)

type VirtualMachine interface {
	Run() InterpretResult
	Interpret(chunk *chunk.Chunk) InterpretResult
	Free()
}
