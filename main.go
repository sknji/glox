package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/urijn/glox/backend/vm"
	"github.com/urijn/glox/backend/vm/stack"
	"github.com/urijn/glox/frontend"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ErrWrongFileExt = errors.New("filename has to have .glox extension")
)

func runFile(v vm.VirtualMachine, filename string) error {
	if !strings.HasSuffix(filename, ".glox") {
		return ErrWrongFileExt
	}

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
	source = bytes.TrimSpace(source)
	if len(source) <= 0 {
		return vm.InterpretOk
	}

	compiler := frontend.NewCompiler(source)

	function, ok := compiler.Compile()
	if !ok {
		return vm.InterpretCompileError
	}

	return v.Interpret(function)
}

func main() {
	flag.Parse()

	var v = stack.NewVM()

	var filename string

	if len(flag.Args()) <= 0 {
		_ = repl(v)
		goto cleanup
	}

	filename = flag.Arg(0)
	_ = runFile(v, filename)

cleanup:
	v.Free()
}
