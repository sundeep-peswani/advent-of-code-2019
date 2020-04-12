package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	filename := os.Args[1]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	steps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	s := newSystem(string(input))
	s.simulate(steps)

	fmt.Println(s.energy())
}

type system struct {
	moons []*moon
}

func newSystem(input string) *system {
	var s system

	for _, line := range strings.Split(string(input), "\n") {
		s.moons = append(s.moons, newMoon(newThreeDim(line)))
	}

	return &s
}

func (s system) String() string {
	var result string

	for _, moon := range s.moons {
		result += moon.String() + "\n"
	}

	return result
}

func (s system) energy() int {
	total := 0

	for _, moon := range s.moons {
		total += moon.energy()
	}

	return total
}

func (s *system) simulate(steps int) {
	for i := 0; i < steps; i++ {
		for _, m := range s.moons {
			for _, o := range s.moons {
				m.applyGravity(*o)
			}
		}

		for _, m := range s.moons {
			m.applyVelocity()
		}
	}
}

type moon struct {
	position threeDim
	velocity threeDim
}

func newMoon(position threeDim) *moon {
	return &moon{position, threeDim{0.0, 0.0, 0.0}}
}

func (m moon) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", m.position, m.velocity)
}

func (m moon) energy() int {
	return m.position.energy() * m.velocity.energy()
}

func (m *moon) applyGravity(o moon) {
	var mGravity threeDim

	mGravity.x = calculateGravity(m.position.x, o.position.x)
	mGravity.y = calculateGravity(m.position.y, o.position.y)
	mGravity.z = calculateGravity(m.position.z, o.position.z)

	m.velocity.applyForce(mGravity)
}

func (m *moon) applyVelocity() {
	m.position.applyForce(m.velocity)
}

type threeDim struct {
	x, y, z int
}

func newThreeDim(input string) threeDim {
	re := regexp.MustCompile(`.=([-\d]+)`)
	matches := re.FindAllStringSubmatch(input, 3)

	var values []int
	for _, match := range matches {
		val, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, val)
	}

	x, y, z := values[0], values[1], values[2]
	return threeDim{x, y, z}
}

func (t threeDim) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", t.x, t.y, t.z)
}

func (t *threeDim) applyForce(f threeDim) {
	t.x += f.x
	t.y += f.y
	t.z += f.z
}

func (t threeDim) energy() int {
	return abs(t.x) + abs(t.y) + abs(t.z)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func calculateGravity(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}
