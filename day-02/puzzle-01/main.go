package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	filename := os.Args[1]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var gap intcode
	gap.read(string(input))
	gap.run()
	fmt.Println(gap.opcodes[0])
}

type intcode struct {
	opcodes []int
}

func (i *intcode) read(input string) {
	for _, instruction := range strings.Split(input, ",") {
		op, err := strconv.Atoi(instruction)
		if err != nil {
			fmt.Printf("Failed to read instruction: %s\n%v\n", instruction, err)
			os.Exit(1)
		}

		i.opcodes = append(i.opcodes, op)
	}
}

func (i *intcode) run() {
	curr := 0

	for curr < len(i.opcodes) {
		switch (i.opcodes[curr]) {
		case 1:
			i.add(i.opcodes[curr + 1], i.opcodes[curr + 2], i.opcodes[curr + 3])
			curr += 4
			break

		case 2:
			i.multiply(i.opcodes[curr + 1], i.opcodes[curr + 2], i.opcodes[curr + 3])
			curr += 4
			break

		case 99:
			return

		default:
			fmt.Printf("Invalid opcode: %d\n", i.opcodes[curr])
			os.Exit(1)
		}
	}
}

func (i *intcode) get(position int) int {
	if position >= len(i.opcodes) {
		fmt.Printf("Attempting to access invalid position: %d\n", position)
		os.Exit(1)
	}

	return i.opcodes[position]
}

func (i *intcode) set(position, value int) {
	if position >= len(i.opcodes) {
		fmt.Printf("Attempting to access invalid position: %d\n", position)
		os.Exit(1)
	}

	i.opcodes[position] = value
}

func (i *intcode) add(registerA, registerB, dest int) {
	a, b := i.get(registerA), i.get(registerB)
	i.set(dest, a + b)
}

func (i *intcode) multiply(registerA, registerB, dest int) {
	a, b := i.get(registerA), i.get(registerB)
	i.set(dest, a * b)
}
