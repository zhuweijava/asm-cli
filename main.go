package main

import (
	"fmt"
	"io"
	"os"

	"github.com/c-bata/go-prompt"
)

const (
	keyX86  = "x86"
	keyX64  = "x64"
	key8086 = "8086"
)

var machineMap = map[string]machine{}

type machine interface {
	displayRegisters()
	displayStack()
	setOutput(io.Writer)
	execute(string) error
}

func init() {
	initX86()
	initX64()
	init8086()
}

func main() {
	machineName := keyX64
	//machineName := keyX86
	ma, ok := machineMap[machineName]
	if !ok {
		fmt.Println("wrong key")
		os.Exit(1)
	}

	ma.displayRegisters()
	ma.displayStack()

	for {
		fmt.Println("Input q to quit.")
		t := prompt.Input("> ", completer)
		if t == "q" || t == "quit" {
			break
		}
		ma.execute(t)
		ma.displayRegisters()
		ma.displayStack()
	}

}
