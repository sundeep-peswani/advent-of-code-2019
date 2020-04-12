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
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	filename := os.Args[1]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getRepetition(string(input)))
}

func getRepetition(input string) int64 {
	s := newSystem(string(input))

	l := int64(1)
	for _, axis := range []string{"x", "y", "z"} {
		positions, velocities := make([]int64, len(s.moons)), make([]int64, len(s.moons))

		for i := range s.moons {
			positions[i] = s.moons[i].position.get(axis)
			velocities[i] = s.moons[i].velocity.get(axis)
		}

		l = lcm(l, findCycle(positions, velocities))
	}

	return l
}

func findCycle(positions, velocities []int64) int64 {
	p := append([]int64(nil), positions...)
	v := append([]int64(nil), velocities...)

	for i := int64(1); ; i++ {
		for j := range v {
			for k := range p {
				v[j] += calculateGravity(p[j], p[k])
			}
		}

		for j := range v {
			p[j] += v[j]
		}

		if equals(p, positions) && equals(v, velocities) {
			return i
		}
	}
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

func equals(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func (s system) energy() int64 {
	total := int64(0)

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
	return &moon{position, threeDim{0, 0, 0}}
}

func (m moon) equals(o moon) bool {
	return m.position.equals(o.position) && m.velocity.equals(o.velocity)
}

func (m moon) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", m.position, m.velocity)
}

func (m moon) energy() int64 {
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
	x, y, z int64
}

func newThreeDim(input string) threeDim {
	re := regexp.MustCompile(`.=([-\d]+)`)
	matches := re.FindAllStringSubmatch(input, 3)

	var values []int64
	for _, match := range matches {
		val, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, int64(val))
	}

	x, y, z := values[0], values[1], values[2]
	return threeDim{x, y, z}
}

func (t threeDim) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", t.x, t.y, t.z)
}

func (t threeDim) get(axis string) int64 {
	switch axis {
	case "x":
		return t.x
	case "y":
		return t.y
	case "z":
		return t.z
	default:
		return 0
	}
}

func (t threeDim) equals(o threeDim) bool {
	return t.x == o.x && t.y == o.y && t.z == o.z
}

func (t *threeDim) applyForce(f threeDim) {
	t.x += f.x
	t.y += f.y
	t.z += f.z
}

func (t threeDim) energy() int64 {
	return abs(t.x) + abs(t.y) + abs(t.z)
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func calculateGravity(a, b int64) int64 {
	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}
