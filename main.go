package main

import (
	"github.com/urijn/glox/backend/vm/stack"
)

func main() {
	var vm = stack.NewVM()

	chunk := stack.NewChunk()

	constant := chunk.AddConstant(1.2)
	chunk.Write(stack.OpConstant, 123)
	chunk.Write(constant, 123)

	constant = chunk.AddConstant(3.4)
	chunk.Write(stack.OpConstant, 123)
	chunk.Write(constant, 123)

	chunk.Write(stack.OpAdd, 123)

	constant = chunk.AddConstant(5.6)
	chunk.Write(stack.OpConstant, 123)
	chunk.Write(constant, 123)

	chunk.Write(stack.OpDivide, 123)
	chunk.Write(stack.OpNegate, 123)

	chunk.Write(stack.OpReturn, 123)

	chunk.Disassemble("test chunk")

	vm.Interpret(chunk)

	vm.Free()
	chunk.Free()
}
