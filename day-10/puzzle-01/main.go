package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
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

	sm := newSpaceMap(string(input))
	optimal, count := sm.getOptimalLocation()

	fmt.Printf("Monitoring station should be at %s: %d\n", optimal, count)

	os.Exit(0)
}

type spaceMap struct {
	asteroids []coord
}

func removeBlocked(polars []polar) []polar {
	result := []polar{polars[0]}

	for i := 1; i < len(polars); i++ {
		if !polars[i].isBehind(polars[i-1]) {
			result = append(result, polars[i])
		}
	}

	return result
}

func newSpaceMap(input string) *spaceMap {
	var sm spaceMap

	for y, line := range strings.Split(input, "\n") {
		for x, pos := range line {
			if pos == '#' {
				sm.asteroids = append(sm.asteroids, coord{x, y})
			}
		}
	}

	return &sm
}

func (sm spaceMap) getOptimalLocation() (coord, int) {
	opt, maxCount := sm.asteroids[0], 0

	for _, c := range sm.asteroids {
		polars := sm.getPolars(c)
		sort.Sort(byDistance(polars))

		remaining := removeBlocked(polars)
		count := len(remaining)

		if count > maxCount {
			opt = c
			maxCount = count
		}
	}

	return opt, maxCount
}

func (sm spaceMap) getPolars(src coord) []polar {
	var polars []polar

	for _, c := range sm.asteroids {
		if !c.Equals(src) {
			polars = append(polars, c.toPolar(src))
		}
	}

	return polars
}

type coord struct {
	x, y int
}

func (c coord) Equals(other coord) bool {
	return c.x == other.x && c.y == other.y
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c coord) toPolar(pole coord) polar {
	x, y := float64(c.x-pole.x), float64(c.y-pole.y)
	p := polar{math.Sqrt(x*x + y*y), math.Atan2(y, x)}
	return p
}

type polar struct {
	r, phi float64
}

func (p polar) String() string {
	return fmt.Sprintf("polar{%.5f, %.5f}", p.r, p.phi)
}

func (p polar) isBehind(c polar) bool {
	if p.phi == c.phi {
		return p.r > c.r
	}
	return false
}

type byDistance []polar

func (b byDistance) Len() int { return len(b) }

func (b byDistance) Less(i, j int) bool {
	if b[i].phi < b[j].phi {
		return true
	}
	if b[i].phi > b[j].phi {
		return false
	}
	return b[i].r < b[j].r
}

func (b byDistance) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
