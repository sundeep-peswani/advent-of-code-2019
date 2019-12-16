package intcode

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Intcode is a system
type Intcode struct {
	Id int

	program []int
	running bool

	instPtr int
	in *bufio.Reader
	out *bufio.Writer
	scanner *bufio.Scanner
}

// NewIntcode returns a new intcode system
func NewIntcode(r io.Reader, w io.Writer) *Intcode {
	var system Intcode
	system.out = bufio.NewWriter(w)
	system.in = bufio.NewReader(r)
	system.scanner = bufio.NewScanner(r)

	return &system
}

// Load loads the Intcode system with a new program
func (i *Intcode) Load(input string) {
	for _, instruction := range strings.Split(input, ",") {
		op, err := strconv.Atoi(instruction)
		if err != nil {
			fmt.Printf("Failed to load instruction: %s\n%v\n", instruction, err)
			os.Exit(1)
		}

		i.program = append(i.program, op)
	}
}

// Run runs the loaded program
func (i *Intcode) Run() {
	i.running = true
	i.instPtr = 0

	for i.IsRunning() && i.instPtr < len(i.program) {
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

		case 5:
			i.jumpIfTrue()
			break

		case 6:
			i.jumpIfFalse()
			break
		
		case 7:
			i.lessThan()
			break
		
		case 8:
			i.equals()
			break

		case 99:
			i.running = false
			return

		default:
			fmt.Printf("Invalid opcode: %d\n", i.program[i.instPtr])
			os.Exit(1)
		}
	}
}

// IsRunning returns if the Intcode is still running
func (i *Intcode) IsRunning() bool {
	return i.running
}

func (i *Intcode) isImmediateMode(parameter int) bool {
	return (i.program[i.instPtr] / int(math.Pow10(parameter + 1))) % 10 == 1
}

func (i *Intcode) get(position int) int {
	if position >= len(i.program) {
		fmt.Printf("Attempting to access invalid position: %d\n", position)
		os.Exit(1)
	}

	return i.program[position]
}

func (i *Intcode) getParam(parameter int) int {
	if i.isImmediateMode(parameter) {
		return i.get(i.instPtr + parameter)
	}

	return i.get(i.get(i.instPtr + parameter))
}

func (i *Intcode) getDest(parameter int) int {
	return i.get(i.instPtr + parameter)
}

func (i *Intcode) set(position, value int) {
	if position >= len(i.program) {
		fmt.Printf("Attempting to access invalid position: %d\n", position)
		os.Exit(1)
	}

	i.program[position] = value
}

func (i *Intcode) add() {
	a, b, dest := i.getParam(1), i.getParam(2), i.getDest(3)
	i.set(dest, a + b)
	i.instPtr += 4
}

func (i *Intcode) multiply() {
	a, b, dest := i.getParam(1), i.getParam(2), i.getDest(3)
	i.set(dest, a * b)
	i.instPtr += 4
}

func (i *Intcode) read() {
	i.scanner.Scan()
	str := i.scanner.Text()
	
	if len(str) == 0 {
		fmt.Println("Empty input received")
		os.Exit(1)
	}

	a, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Intcode %d: Invalid input '%s': %v\n", i.Id, str, err)
		os.Exit(1)
	}
	
	dest := i.getDest(1)
	i.set(dest, a)
	i.instPtr += 2
}

func (i *Intcode) write() {
	src := i.getParam(1)

	defer i.out.Flush()
	_, err := i.out.WriteString(fmt.Sprintf("%d\n", src))
	if err != nil {
		fmt.Printf("Unable to write to output: %v\n", err)
		os.Exit(1)
	}
	i.instPtr += 2
}


func (i *Intcode) jumpIfTrue() {
	a, dest := i.getParam(1), i.getParam(2)
	if a != 0 {
		i.instPtr = dest
	} else {
		i.instPtr += 3
	}
}

func (i *Intcode) jumpIfFalse() {
	a, dest := i.getParam(1), i.getParam(2)
	if a == 0 {
		i.instPtr = dest
	} else {
		i.instPtr += 3
	}
}

func (i *Intcode) lessThan() {
	a, b, dest := i.getParam(1), i.getParam(2), i.getDest(3)
	if a < b {
		i.set(dest, 1)
	} else {
		i.set(dest, 0)
	}
	i.instPtr += 4
}

func (i *Intcode) equals() {
	a, b, dest := i.getParam(1), i.getParam(2), i.getDest(3)
	if a == b {
		i.set(dest, 1)
	} else {
		i.set(dest, 0)
	}
	i.instPtr += 4
}
