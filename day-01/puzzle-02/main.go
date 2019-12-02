package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	total := 0
	for _, mass := range(strings.Split(string(input), "\n")) {
		if len(mass) == 0 {
			continue
		}

		m, err := strconv.Atoi(mass)
		if err != nil {
			fmt.Println(mass)
			fmt.Println(err)
			os.Exit(1)
		}
		total += calculateFuel(m)
	}

	fmt.Println(total)
}

func calculateFuel(mass int) int {
	fuel := int(math.Floor(float64(mass) / 3.0)) - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + calculateFuel(fuel)
}