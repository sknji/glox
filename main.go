package main

import "github.com/urijn/glox/lib"

func main() {
	var vm = lib.NewVM()

	chunk := lib.NewChunk()

	constant := chunk.AddConstant(1.2)
	chunk.Write(lib.OpConstant, 123)
	chunk.Write(constant, 123)

	constant = chunk.AddConstant(3.4)
	chunk.Write(lib.OpConstant, 123)
	chunk.Write(constant, 123)

	chunk.Write(lib.OpAdd, 123)

	constant = chunk.AddConstant(5.6)
	chunk.Write(lib.OpConstant, 123)
	chunk.Write(constant, 123)

	chunk.Write(lib.OpDivide, 123)
	chunk.Write(lib.OpNegate, 123)

	chunk.Write(lib.OpReturn, 123)

	chunk.Disassemble("test chunk")

	vm.Interpret(chunk)

	vm.Free()
	chunk.Free()
}
