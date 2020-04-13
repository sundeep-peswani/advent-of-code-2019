package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	in, out := make(chan int, 1), make(chan int, 3)
	i := intcode.NewIntcode(in, out)
	i.Load(string(input))

	go func() { i.Run() }()

	grid := make(map[int]int)
	for {
		x, y, tile := <-out, <-out, <-out

		grid[coord(x, y)] = tile

		if !i.IsRunning() {
			break
		}
	}

	tiles := 0
	for _, tile := range grid {
		if tile == 2 {
			tiles++
		}
	}

	fmt.Println(tiles)
}

func coord(x, y int) int {
	return y*100 + x
}
