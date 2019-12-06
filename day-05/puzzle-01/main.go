package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
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

	var system intcode
	system.in = bufio.NewReader(os.Stdin)
	system.out = bufio.NewWriter(os.Stdout)

	system.load(string(input))
	system.run()
	
	os.Exit(0)
}

type intcode struct {
	program []int
	instPtr int
	in *bufio.Reader
	out *bufio.Writer
}

func (i *intcode) load(input string) {
	for _, instruction := range strings.Split(input, ",") {
		op, err := strconv.Atoi(instruction)
		if err != nil {
			fmt.Printf("Failed to load instruction: %s\n%v\n", instruction, err)
			os.Exit(1)
		}

		i.program = append(i.program, op)
	}
}

func (i *intcode) run() {
	i.instPtr = 0

	for i.instPtr < len(i.program) {
		switch (i.program[i.instPtr] % 100) {
		case 1:
			i.add()
			break

		case 2:
			i.multiply()
			break

		case 3:
			i.read()
			break

		case 4:
			i.write()
			break

		case 99:
			return

		default:
			fmt.Printf("Invalid opcode: %d\n", i.program[i.instPtr])
			os.Exit(1)
		}
	}
}

func (i *intcode) isImmediateMode(parameter int) bool {
	return (i.program[i.instPtr] / int(math.Pow10(parameter + 1))) % 10 == 1
}

func (i *intcode) get(position int, isImmediateMode bool) int {
	if isImmediateMode {
		return i.program[position]
	}

	if position >= len(i.program) {
		fmt.Printf("Attempting to access invalid position: %d\n", position)
		os.Exit(1)
	}

	return i.program[i.program[position]]
}

func (i *intcode) getParam(parameter int) int {
	return i.get(i.instPtr + parameter, i.isImmediateMode(parameter))
}

func (i *intcode) getDest(parameter int) int {
	return i.get(i.instPtr + parameter, true)
}

func (i *intcode) set(position, value int) {
	if position >= len(i.program) {
		fmt.Printf("Attempting to access invalid position: %d\n", position)
		os.Exit(1)
	}

	i.program[position] = value
}

func (i *intcode) add() {
	a, b, dest := i.getParam(1), i.getParam(2), i.getDest(3)
	i.set(dest, a + b)
	i.instPtr += 4
}

func (i *intcode) multiply() {
	a, b, dest := i.getParam(1), i.getParam(2), i.getDest(3)
	i.set(dest, a * b)
	i.instPtr += 4
}

func (i *intcode) read() {
	str, _ := i.in.ReadString('\n')
	a, err := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		fmt.Printf("Invalid input '%s': %v\n", str, err)
		os.Exit(1)
	}
	
	dest := i.getDest(1)
	i.set(dest, a)
	i.instPtr += 2
}

func (i *intcode) write() {
	src := i.getDest(1)
	a := i.get(src, true)

	defer i.out.Flush()
	_, err := i.out.WriteString(fmt.Sprintf("%d\n", a))
	if err != nil {
		fmt.Printf("Unable to write to output: %v\n", err)
		os.Exit(1)
	}
	i.instPtr += 2
}
