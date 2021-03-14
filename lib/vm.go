package lib

const StackMax = 512

type InterpretResult int

type VM struct {
	chunk *Chunk
	ip    *uint8

	stack [StackMax]Value
}

func NewVM() *VM {
	return &VM{}
}

func (vm *VM) run() InterpretResult {
	return 0
}

func (vm *VM) free() {

}
