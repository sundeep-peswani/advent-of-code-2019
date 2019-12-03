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

	wires := strings.Split(string(input), "\n")
	fmt.Println(findNearestIntersection(wires))
}

type panel struct {
	grid map[coord][]int
	wires []string
}

type coord struct { x, y int }

func findNearestIntersection(wires []string) int {
	if len(wires) != 2 {
		return -1
	}

	var p panel
	p.grid = make(map[coord][]int)
	p.wires = wires

	for i, wire := range wires {
		// fmt.Printf("Adding wire %d\n", i+1)
		p.addWire(i, wire)
	}

	result := -1
	for _, paths := range(p.grid) {
		if !allTrue(paths) {
			continue
		}

		// fmt.Printf("Found intersection at %s\n", c)

		d := sum(paths)
		if result == -1 || d < result {
			// fmt.Printf("Smallest distance at %s: %d\n", c, d)
			result = d
		}
	}
	
	return result
}

func allTrue(paths []int) bool {
	for _, p := range paths {
		if p == 0 {
			return false
		}
	}

	return true
}

func sum(paths []int) int {
	total := 0
	for _, p := range paths {
		total += p
	}
	return total
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return n * -1
}

func (c coord) distance() int {
	return abs(c.x) + abs(c.y)
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (p *panel) addWire(index int, path string) {
	current := coord{0, 0}
	stepCount := 0

	for _, step := range strings.Split(path, ",") {
		direction:= step[0]
		length, err := strconv.Atoi(step[1:])
		if err != nil {
			fmt.Printf("Invalid step: %v\n", step)
			os.Exit(1)
		}

		// fmt.Printf("Starting at %s, going %v for %d steps\n", current, direction, length)

		switch (direction) {
		case 'U':
			current = p.drawLine(current, index, stepCount, length, 0, -1)
			break
		case 'D':
			current = p.drawLine(current, index, stepCount, length, 0, 1)
			break
		case 'L':
			current = p.drawLine(current, index, stepCount, length, -1, 0)
			break
		case 'R':
			current = p.drawLine(current, index, stepCount, length, 1, 0)
			break
		default:
			fmt.Printf("Invalid step: %v\n", step)
			os.Exit(1)
		}

		stepCount += length
	}
}

func (p *panel) drawLine(start coord, wireIndex, stepCount, length, xStep, yStep int) coord {
	for i := 0; i < length; i++ {
		start.x += xStep;
		start.y += yStep;
		stepCount++

		if _, ok := p.grid[start]; !ok {
			p.grid[start] = make([]int, len(p.wires))
		}

		p.grid[start][wireIndex] = stepCount
	}

	return start
}
