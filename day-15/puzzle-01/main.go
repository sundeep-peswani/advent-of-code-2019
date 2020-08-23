package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sundeep-peswani/advent-of-code-2019/intcode"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need input file")
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	d := newDroid(string(input))
	fmt.Println(d.findOxygenSystem())
}

const (
	moveNorth = 1
	moveSouth = 2
	moveWest  = 3
	moveEast  = 4
)

const (
	statusWall   = 0
	statusMoved  = 1
	statusOxygen = 2
	statusDroid  = 3
)

type droid struct {
	system        *intcode.Intcode
	in, out, quit chan int

	x, y, last int
	g          *grid
}

func newDroid(program string) *droid {
	var d droid

	d.in, d.out, d.quit = make(chan int, 1), make(chan int, 1), make(chan int, 1)
	d.system = intcode.NewIntcode(d.in, d.out, d.quit)
	d.system.Load(program)
	// d.system.Debug = true

	d.g = newGrid([]rune{'#', '.', 'O', 'D'})
	return &d
}

func (d *droid) String() string {
	return d.g.String()
}

func (d *droid) findOxygenSystem() int {
	go func() { d.system.Run() }()

	count := 0
	d.in <- moveNorth

	for status := range d.out {
		switch d.last {
		case moveNorth:
			d.y--
		case moveSouth:
			d.y++
		case moveWest:
			d.x--
		case moveEast:
			d.x++
		}

		d.g.add(d.x, d.y, status)

		if status == statusWall {
			switch
		}
		d.last = d.getAvailableMove()
		d.in <- d.last


		switch status {
		case statusWall:
			switch d.last {
			case moveNorth:
				d.g.add(d.x, d.y-1, status)
			case moveSouth:
				d.g.add(d.x, d.y+1, status)
			case moveWest:
				d.g.add(d.x-1, d.y, status)
			case moveEast:
				d.g.add(d.x+1, d.y, status)
			}

			d.last = d.getAvailableMove()
			d.in <- d.last

		case statusMoved:


		case statusOxygen:
			d.system.Stop()
			break
		}

		time.Sleep(250 * time.Millisecond)
		// clear()
		fmt.Println(d)
		fmt.Printf("\nMoves: %d\n", count)
	}

	return count
}

func (d *droid) getAvailableMove() int {
	available := d.getAvailableMoves()
	if len(available) == 0 {
		return -1
	}
	return available[rand.Intn(len(available))]
}

func (d *droid) getAvailableMoves() []int {
	var moves []int

	if d.g.peek(d.x, d.y-1) != statusWall {
		moves = append(moves, moveNorth)
	}
	if d.g.peek(d.x, d.y+1) != statusWall {
		moves = append(moves, moveSouth)
	}
	if d.g.peek(d.x-1, d.y) != statusWall {
		moves = append(moves, moveEast)
	}
	if d.g.peek(d.x+1, d.y) != statusWall {
		moves = append(moves, moveWest)
	}

	return moves
}

type grid struct {
	minX, maxX, minY, maxY int
	cells                  map[string]int
	symbols                []rune
}

func newGrid(symbols []rune) *grid {
	return &grid{-5, 5, -5, 5, make(map[string]int), symbols}
}

func (g *grid) String() string {
	var b strings.Builder

	for y := g.minY; y < g.maxY; y++ {
		for x := g.minX; x < g.maxX; x++ {
			if c, ok := g.cells[g.coord(x, y)]; !ok {
				b.WriteRune(' ')
			} else {
				b.WriteRune(g.symbols[c])
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

func (g *grid) add(x, y, s int) {
	if x < g.minX {
		g.minX = x - 10
	}
	if x > g.maxX {
		g.maxX = x + 10
	}
	if y < g.minX {
		g.minY = y - 10
	}
	if y > g.maxY {
		g.maxY = y + 10
	}

	if s < 0 || s > len(g.symbols) {
		log.Fatalf("Unknown symbol: %d\n", s)
	}

	g.cells[g.coord(x, y)] = s
}

func (g *grid) peek(x, y int) int {
	if s, ok := g.cells[g.coord(x, y)]; ok {
		return s
	}
	return -1
}

func (g *grid) coord(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
