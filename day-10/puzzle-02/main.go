package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

const twoPi = 2 * math.Pi

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
	targets := sm.vaporize(200, coord{20, 21, 0, 0})

	fmt.Printf("%s => %d\n", targets[199], targets[199].x*100+targets[199].y)

	os.Exit(0)
}

type spaceMap struct {
	asteroids []coord
}

func newSpaceMap(input string) *spaceMap {
	var sm spaceMap

	for y, line := range strings.Split(input, "\n") {
		for x, pos := range line {
			if pos == '#' {
				sm.asteroids = append(sm.asteroids, coord{x, y, 0, 0})
			}
		}
	}

	return &sm
}

func (sm *spaceMap) toPolar(pole coord) {
	for i := range sm.asteroids {
		if !sm.asteroids[i].Equals(pole) {
			sm.asteroids[i].toPolar(pole)
		}
	}
}

func (sm spaceMap) vaporize(n int, laser coord) []coord {
	if n > len(sm.asteroids) {
		n = len(sm.asteroids)
	}

	sm.toPolar(laser)

	arrangement := make(map[float64][]coord)
	var angles []float64

	for _, c := range sm.asteroids {
		if _, ok := arrangement[c.phi]; !ok {
			arrangement[c.phi] = []coord{c}
			angles = append(angles, c.phi)
		} else {
			arrangement[c.phi] = append(arrangement[c.phi], c)
		}
	}

	sort.Slice(angles, func(i, j int) bool { return clockwisePolar(angles[i]) < clockwisePolar(angles[j]) })
	for _, angle := range angles {
		sort.Slice(arrangement[angle], func(i, j int) bool { return arrangement[angle][i].r < arrangement[angle][j].r })
	}

	var coords []coord
	for len(coords) < n {
		for _, angle := range angles {
			if len(arrangement[angle]) == 0 {
				continue
			}

			coords = append(coords, arrangement[angle][0])
			arrangement[angle] = arrangement[angle][1:]
		}
	}

	return coords
}

type coord struct {
	x, y   int
	r, phi float64
}

func (c coord) Equals(other coord) bool {
	return c.x == other.x && c.y == other.y
}

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d, %.5f, %.5f)", c.x, c.y, c.r, c.phi)
}

func (c *coord) toPolar(pole coord) {
	x, y := float64(c.x-pole.x), float64(pole.y-c.y)
	c.phi = math.Atan2(y, x)
	if c.phi < 0 {
		c.phi += twoPi
	}
	c.phi = c.phi / twoPi * 360.0
	c.r = math.Sqrt(x*x + y*y)
}

// polar coordinates start at 0 in the east and rotate counter-clockwise: east = 0, north = pi/2, west = pi, south = 3pi/2
// to sort with north, we need to adjust all angles by pi/2 rad (90deg) and then scale to 2*pi
func clockwisePolar(angle float64) float64 {
	//return math.Mod(twoPi+angle-twoPi/4, twoPi) / twoPi
	return math.Mod(math.Mod(360.0-angle, 360.0)+90.0, 360.0)
}
