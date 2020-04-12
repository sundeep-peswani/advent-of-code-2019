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

	max := maxThrusterSignal(string(input), []int{5, 6, 7, 8, 9})
	fmt.Println(max)

	os.Exit(0)
}

func maxThrusterSignal(program string, options []int) int {
	max := 0
	perms := generatePermutations(options)

	for _, perm := range perms {
		signal := runAmps(program, perm)

		if signal > max {
			max = signal
		}
	}

	return max
}

func runAmps(program string, options []int) int {
	var streams []chan int
	var amps []*intcode.Intcode

	for i := 0; i < len(options); i++ {
		streams = append(streams, make(chan int))
	}

	for i, phase := range options {
		amp := intcode.NewIntcode(streams[i], streams[(i+1)%len(options)])
		amp.ID = phase
		amp.Load(program)
		amps = append(amps, amp)

		go amp.Run()
		streams[i] <- phase
	}
	streams[0] <- 0

	for amps[0].IsRunning() {
	}

	return <-streams[0]
}

func generatePermutations(options []int) [][]int {
	if len(options) == 0 {
		return [][]int{}
	}

	if len(options) == 1 {
		return [][]int{options}
	}

	var result [][]int

	for i, opt := range options {
		optionsExceptThis := append(append([]int{}, options[:i]...), options[i+1:]...)
		perms := generatePermutations(optionsExceptThis)

		for _, perm := range perms {
			result = append(result, append(append([]int{}, opt), perm...))
		}
	}

	return result
}
