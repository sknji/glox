package frontend

import (
	"fmt"
	"github.com/urijn/glox/chunk"
	"github.com/urijn/glox/value"
)

type FunctionType int

const (
	TypeFunction FunctionType = iota + 1
	TypeScript
)

type ObjectFunction struct {
	name string

	// Stores the number of parameters the function expects.
	arity int

	ftype FunctionType

	chunk *chunk.Chunk
}

func NewObjectFunction(ftype FunctionType) *ObjectFunction {
	return &ObjectFunction{
		name:  "",
		arity: 0,
		ftype: ftype,
		chunk: chunk.NewChunk(),
	}
}

func (of *ObjectFunction) Chunk() *chunk.Chunk {
	return of.chunk
}

func (of *ObjectFunction) Type() value.ObjType {
	return value.ObjFunction
}

func (of *ObjectFunction) String() string {
	return fmt.Sprintf("<fn %s>", of.ParsedName())
}

func (of *ObjectFunction) ParsedName() string {
	name := "<script>"
	if of.name == "" {
		name = of.name
	}
	return name
}
