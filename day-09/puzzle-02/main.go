package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sundeep-peswani/advent-of-code-2019/intcode"
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

	in, out := make(chan int, 1), make(chan int, 1)
	system := intcode.NewIntcode(in, out)
	system.Load(string(input))

	go system.Run()
	in <- 2

	fmt.Println(<-out)
}
