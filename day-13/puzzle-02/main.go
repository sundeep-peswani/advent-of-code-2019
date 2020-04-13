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
		log.Fatal("Needs input file")
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	a := newArcade(string(input), 43, 22)
	fmt.Println(a.play())
}

const (
	empty  = 0
	wall   = 1
	block  = 2
	paddle = 3
	ball   = 4
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

type arcade struct {
	system        *intcode.Intcode
	in, out, quit chan int

	grid       [][]int
	blocksLeft int
	score      int

	ball, paddle coord
}

func newArcade(input string, w, h int) *arcade {
	var a arcade

	rand.Seed(time.Now().UnixNano())

	a.in, a.out, a.quit = make(chan int, 1), make(chan int), make(chan int, 1)
	a.system = intcode.NewIntcode(a.in, a.out, a.quit)
	a.system.Load(string(input))
	// a.system.Debug = true

	a.grid = make([][]int, h)
	for i := 0; i < h; i++ {
		a.grid[i] = make([]int, w)
	}

	a.blocksLeft = 0
	a.ball = coord{0, 0}
	a.paddle = coord{0, 0}

	return &a
}

func (a arcade) String() string {
	var b strings.Builder

	for _, row := range a.grid {
		for _, cell := range row {
			switch cell {
			case empty:
				b.WriteRune(' ')
			case wall:
				b.WriteRune('█')
			case block:
				b.WriteRune('▒')
			case paddle:
				b.WriteRune('_')
			case ball:
				b.WriteRune('●')
			}
		}
		b.WriteString("\n")
	}

	b.WriteString(fmt.Sprintf("\nScore: %d\n", a.score))
	// b.WriteString(fmt.Sprintf("Ball: %s, Paddle: %s\n", a.ball, a.paddle))

	return b.String()
}

func (a *arcade) insertCoins(coins int) {
	a.system.Program[0] = coins
}

func (a *arcade) play() int {
	go func() {
		a.system.Run()
	}()

	go func() {
		buffer := make([]int, 3)
		i := 0

		for {
			select {
			case <-a.quit:
				return

			case o := <-a.out:
				buffer[i] = o

				if i == 2 {
					a.update(buffer[0], buffer[1], buffer[2])
					i = -1
				}

				i++
			}
		}
	}()

	a.insertCoins(2)

	for {
		select {
		case <-a.quit:
			return a.score

		default:
			// time.Sleep(40 * time.Millisecond)
			a.render()
			if a.system.IsRunning() {
				a.in <- a.joystick()
			}
			if a.blocksLeft == 0 {
				a.system.Stop()
			}
		}
	}
}

func (a *arcade) update(x, y, tileType int) {
	if x == -1 && y == 0 {
		a.score = tileType
		return
	}

	if tileType < empty || tileType > ball {
		log.Fatalf("Invalid input: %d, %d, %d\n", x, y, tileType)
	}

	if y > len(a.grid) {
		log.Fatalf("Row id too large: %d\n", y)
	}

	if x > len(a.grid[y]) {
		log.Fatalf("Column id too large: %d\n", x)
	}

	if a.grid[y][x] == block {
		a.blocksLeft--
	}

	switch tileType {
	case block:
		a.blocksLeft++
	case ball:
		a.ball = coord{x, y}
	case paddle:
		a.paddle = coord{x, y}
	}

	a.grid[y][x] = tileType
}

func (a arcade) joystick() int {
	if a.ball.y > a.paddle.y {
		return 0
	}

	if a.ball.x < a.paddle.x {
		return -1
	} else if a.ball.x > a.paddle.x {
		return 1
	}
	return 0
}

func (a arcade) render() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(a)
}
