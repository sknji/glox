package stack

import "fmt"

func (v *VM) runtimeError(format string, val ...interface{}) {
	// TODO:
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

	v.currFrameChunk().DisassembleInstruction(v.currFrameIp())
}
