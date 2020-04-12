package intcode

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Intcode is a system
type Intcode struct {
	ID    int
	Debug bool

	program []int
	running bool

	instPtr        int
	relativeOffset int

	in  chan int
	out chan int
}

// NewIntcode returns a new intcode system
func NewIntcode(r, w chan int) *Intcode {
	var system Intcode

	system.instPtr = 0
	system.relativeOffset = 0
	system.in = r
	system.out = w

	system.program = make([]int, 10000)

	return &system
}

// Load loads the Intcode system with a new program
func (i *Intcode) Load(input string) {
	for j, instruction := range strings.Split(input, ",") {
		op, err := strconv.Atoi(instruction)
		if err != nil {
			fmt.Printf("Failed to load instruction at %d: %s\n%v\n", j, instruction, err)
			os.Exit(1)
		}

		i.program[j] = op
	}
}

// Run runs the loaded program
func (i *Intcode) Run() {
	i.running = true
	i.instPtr = 0

	for i.IsRunning() && i.instPtr < len(i.program) {
		switch i.program[i.instPtr] % 100 {
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

		case 9:
			i.offset()
			break

		case 99:
			i.halt()
			return

		default:
			fmt.Printf("%d: Invalid opcode: %d\n", i.ID, i.program[i.instPtr])
			os.Exit(1)
		}
	}
}

// IsRunning returns if the Intcode is still running
func (i *Intcode) IsRunning() bool {
	return i.running
}

func (i *Intcode) log(message string, params ...interface{}) {
	if i.Debug {
		fmt.Println(fmt.Sprintf(fmt.Sprintf("%d @ %d = %d: %s", i.ID, i.instPtr, i.program[i.instPtr], message), params...))
	}
}

const (
	position  int = 0
	immediate int = 1
	relative  int = 2
)

type param struct {
	position int
	value    int
	mode     int
}

func (p param) String() string {
	switch p.mode {
	case position:
		return fmt.Sprintf("(@%d = %d)", p.position, p.value)

	case immediate:
		return fmt.Sprintf("(= %d)", p.value)

	case relative:
		return fmt.Sprintf("(+@%d = %d)", p.position, p.value)

	default:
		return fmt.Sprintf("Unknown param mode: %d", p.mode)
	}
}

func (i *Intcode) getParams(n int) []param {
	modes := i.program[i.instPtr] / 100

	var params []param
	for j := 1; j <= n; j++ {
		p := param{i.program[i.instPtr+j], 0, modes % 10}
		i.get(&p)

		params = append(params, p)
		modes /= 10
	}

	return params
}

func (i *Intcode) get(p *param) {
	switch p.mode {
	case position:
		if p.position >= len(i.program) {
			log.Fatalf("Invalid access for get(): %d\n", p.position)
		}
		p.value = i.program[p.position]

	case immediate:
		p.value = p.position

	case relative:
		p.value = i.program[i.relativeOffset+p.position]

	default:
		log.Fatalf("Unknown param mode: %d", p.mode)
	}
}

func (i *Intcode) set(p param) {
	var pos int

	switch p.mode {
	case position:
		pos = p.position

	case immediate:
		log.Fatalf("Invalid write access: %s\n", p)

	case relative:
		pos = i.relativeOffset + p.position

	default:
		log.Fatalf("Unknown param mode: %d", p.mode)
	}

	if pos < 0 || pos >= len(i.program) {
		log.Fatalf("Attempt to write outside program: %s\n", p)
	}

	i.log("SET %s, @%d", p, pos)
	i.program[pos] = p.value
}

func (i *Intcode) unaryOp(op string, f func(a int) bool) {
	params := i.getParams(2)
	a, dest := params[0], params[1]

	result := f(a.value)
	i.log("%s %s -> %s", op, a, dest)

	if result {
		i.instPtr = dest.value
	} else {
		i.instPtr += 3
	}
}

func (i *Intcode) binaryOp(op string, f func(a, b int) int) {
	params := i.getParams(3)
	a, b, dest := params[0], params[1], params[2]

	dest.value = f(a.value, b.value)
	i.log("%s %s %s -> %s", op, a, b, dest)

	i.set(dest)
	i.instPtr += 4
}

func (i *Intcode) add() {
	i.binaryOp("ADD", func(a, b int) int { return a + b })
}

func (i *Intcode) multiply() {
	i.binaryOp("MUL", func(a, b int) int { return a * b })
}

func (i *Intcode) jumpIfTrue() {
	i.unaryOp("JNZ", func(a int) bool { return a != 0 })
}

func (i *Intcode) jumpIfFalse() {
	i.unaryOp("JZ", func(a int) bool { return a == 0 })
}

func (i *Intcode) lessThan() {
	i.binaryOp("LESS", func(a, b int) int {
		if a < b {
			return 1
		}
		return 0
	})
}

func (i *Intcode) equals() {
	i.binaryOp("EQ", func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	})
}

func (i *Intcode) offset() {
	params := i.getParams(1)
	a := params[0]

	i.log("OFFSET %d -> %s", i.relativeOffset, a)

	i.relativeOffset += a.value
	i.instPtr += 2
}

func (i *Intcode) read() {
	params := i.getParams(1)
	dest := params[0]

	a, ok := <-i.in
	if !ok {
		log.Fatal("Unable to read from in")
	}

	dest.value = a

	i.log("READ -> %s", dest)

	i.set(dest)
	i.instPtr += 2
}

func (i *Intcode) write() {
	params := i.getParams(1)
	a := params[0]

	i.log("WRITE %s", a)

	i.out <- a.value
	i.instPtr += 2
}

func (i *Intcode) halt() {
	i.log("HALT")

	i.running = false
}
