package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/urijn/glox/backend/vm"
	"github.com/urijn/glox/backend/vm/stack"
	"github.com/urijn/glox/frontend"
	"io/ioutil"
	"os"
)

var fileName = flag.String("filename", "", "-f")

func runFile(v vm.VirtualMachine, filename string) error {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	switch run(v, source) {
	case vm.InterpretCompileError:
		os.Exit(65)
	case vm.InterpretRuntimeError:
		os.Exit(70)
	}

	return nil
}

func repl(v vm.VirtualMachine) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")

		source, err := reader.ReadBytes('\n')
		if err != nil {
			continue
		}

		run(v, source)
	}
}

func run(v vm.VirtualMachine, source []byte) vm.InterpretResult {
	//fmt.Println(string(source))

	compiler := frontend.NewCompiler(source)

	chunk, ok := compiler.Compile()
	if !ok {
		return vm.InterpretCompileError
	}

	return v.Interpret(chunk)
}

func main() {
	flag.Parse()

	var v = stack.NewVM()

	if *fileName == "" {
		_ = repl(v)
		goto cleanup
	}

	_ = runFile(v, *fileName)

cleanup:
	v.Free()
}
