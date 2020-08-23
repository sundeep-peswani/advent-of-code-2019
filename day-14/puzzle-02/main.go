package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need input file")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n := newNanofactory()
	r := bufio.NewScanner(f)
	for r.Scan() {
		n.addReaction(r.Text())
	}

	fmt.Printf("Needs %d ORE per fuel\n", n.calculateOre(c("FUEL", 1)))
	fmt.Printf("1 trillion ORE could produce %d FUEL\n", n.calculateFuel(1000000000000, 0, 1000000000000))
}

type nanofactory struct {
	reactions map[string]*reaction
}

func newNanofactory() *nanofactory {
	return &nanofactory{make(map[string]*reaction)}
}

func (n *nanofactory) addReaction(input string) {
	inputs, output := n.parse(input)
	n.reactions[output.name] = newReaction(inputs, output)
}

func (n *nanofactory) calculateOre(in *chemical) int {
	total := 0

	extras := make(map[string]int)
	q := []*chemical{in}

	for len(q) > 0 {
		h := q[0]
		q = q[1:]

		r, ok := n.reactions[h.name]
		if !ok {
			log.Fatalf("Unable to find reaction to produce %s\n", h.name)
		}

		needed := h.quantity
		if x, ok := extras[h.name]; ok {
			if x > needed {
				extras[h.name] -= needed
				continue
			} else {
				delete(extras, h.name)
				needed -= x
			}
		}

		factor := roundUp(needed, r.output.quantity)
		extra := r.output.quantity*factor - needed
		if extra > 0 {
			extras[h.name] = extra
		}

		for _, i := range r.inputs {
			t := i.quantity * factor
			if i.name == "ORE" {
				total += t
				continue
			}

			q = append(q, c(i.name, t))
		}
	}

	return total
}

func (n *nanofactory) calculateFuel(ore, lo, hi int) int {
	mid := (hi + lo) / 2
	needed := n.calculateOre(c("FUEL", mid))

	if mid == lo || needed == ore {
		return mid
	}

	if needed < ore {
		return n.calculateFuel(ore, mid, hi)
	}

	return n.calculateFuel(ore, lo, mid)
}

func (n nanofactory) parse(input string) ([]*chemical, *chemical) {
	parts := strings.Split(input, "=>")
	if len(parts) != 2 {
		log.Fatalf("Invalid chemical input: %s\n", input)
	}

	var inputs []*chemical
	for _, i := range strings.Split(parts[0], ",") {
		inputs = append(inputs, newChemical(i))
	}
	return inputs, newChemical(parts[1])
}

type reaction struct {
	inputs []*chemical
	output *chemical
}

func newReaction(inputs []*chemical, output *chemical) *reaction {
	return &reaction{inputs, output}
}

func (r *reaction) String() string {
	var ins []string

	for _, i := range r.inputs {
		ins = append(ins, i.String())
	}

	return fmt.Sprintf("%s => %s", strings.Join(ins, " + "), r.output)
}

type chemical struct {
	name     string
	quantity int
}

func c(name string, quantity int) *chemical {
	return &chemical{name, quantity}
}

func newChemical(input string) *chemical {
	var q int
	var n string

	fmt.Sscanf(strings.Trim(input, " "), "%d %s", &q, &n)
	return &chemical{n, q}
}

func (c *chemical) String() string {
	return fmt.Sprintf("%d %s", c.quantity, c.name)
}

func (c *chemical) equals(o *chemical) bool {
	return o != nil && c.name == o.name && c.quantity == o.quantity
}

func roundUp(num, div int) int {
	return int(math.Ceil(float64(num) / float64(div)))
}
