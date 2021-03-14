package main

import "github.com/urijn/glox/lib"

func main() {
	chunk := lib.NewChunk()

	var constant = chunk.AddConstant(1.2)
	chunk.Write(lib.OpConstant)
	chunk.Write(constant)

	chunk.Write(lib.OpReturn)

	chunk.Disassemble("test chunk")

	chunk.Free()
}
