package vm

import (
	"github.com/urijn/glox/frontend"
)

type InterpretResult int

const (
	InterpretOk InterpretResult = iota + 1
	InterpretCompileError
	InterpretRuntimeError
)

func (ir *InterpretResult) IsSuccess() bool {
	return *ir == InterpretOk
}

type VirtualMachine interface {
	Run() InterpretResult
	Interpret(function *frontend.ObjectFunction) InterpretResult
	Free()
}
