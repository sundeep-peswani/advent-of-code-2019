package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/sundeep-peswani/advent-of-code-2019/intcode"
)

const (
	up = iota
	left
	down
	right
)

const (
	black = 0
	white = 1
)

const size = 100

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

	r := newRobot()

	r.run(string(input))

	fmt.Println(r)

	fmt.Println(r.countPainted())

	os.Exit(0)
}

type robot struct {
	minX, maxX, minY, maxY int
	x, y                   int
	d                      int
	panels                 [][]int
	painted                [][]bool
}

func newRobot() *robot {
	r := robot{0, size, 0, size, size / 2, size / 2, up, make([][]int, size), make([][]bool, size)}

	for y := 0; y < size; y++ {
		r.panels[y] = make([]int, size)
		r.painted[y] = make([]bool, size)
	}

	return &r
}

func (r robot) String() string {
	result := ""

	for y, row := range r.panels {
		for x, cell := range row {
			if x == r.x && y == r.y {
				switch r.d {
				case up:
					result += "^"
				case down:
					result += "v"
				case left:
					result += "<"
				case right:
					result += ">"
				}
			} else if cell == black {
				result += "."
			} else {
				result += "#"
			}
		}
		result += "\n"
	}

	return result
}

func (r *robot) run(program string) {
	in, out := make(chan int, 1), make(chan int, 2)
	system := intcode.NewIntcode(in, out)
	// system.Debug = true

	system.Load(program)

	go func() { system.Run() }()

	for {
		in <- r.read()

		color := <-out
		direction := <-out

		r.paint(color)
		if direction == 0 {
			r.left()
		} else {
			r.right()
		}

		if !system.IsRunning() {
			return
		}
	}
}

func (r *robot) left() {
	r.d = []int{up, left, down, right}[(r.d+1)%4]
	r.forward()
}

func (r *robot) right() {
	r.d = []int{up, left, down, right}[(r.d+3)%4]
	r.forward()
}

func (r *robot) forward() {
	switch r.d {
	case up:
		r.y--
	case down:
		r.y++
	case left:
		r.x--
	case right:
		r.x++
	}

	r.minX = min(r.x-3, r.minX)
	r.maxX = max(r.x+3, r.maxX)
	r.minY = min(r.y-3, r.minY)
	r.maxY = max(r.y+3, r.maxY)
}

func (r robot) read() int {
	return r.panels[r.y][r.x]
}

func (r *robot) paint(color int) {
	if color != black && color != white {
		log.Fatalf("Unknown color: %d\n", color)
	}

	r.panels[r.y][r.x] = color
	r.painted[r.y][r.x] = true
}

func (r robot) countPainted() int {
	count := 0

	for _, row := range r.painted {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}

	return count
}

func (r robot) pos() string {
	return pos(r.x, r.y)
}

func pos(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
