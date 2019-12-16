package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	
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

	max := maxThrusterSignal(string(input), []int{0, 1, 2, 3, 4})
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
	var ins []*io.PipeReader
	var outs []*io.PipeWriter

	for i := 0; i <= len(options); i++ {
		in, out := io.Pipe()
		ins = append(ins, in)
		outs = append(outs, out)			
	}	
	output := bufio.NewScanner(ins[len(options)])

	for i, phase := range options {
		amp := intcode.NewIntcode(ins[i], outs[i + 1])
		amp.Load(program)
		go amp.Run()
		io.WriteString(outs[i], fmt.Sprintf("%d\n", phase))
	}
	io.WriteString(outs[0], fmt.Sprintf("0\n"))
		
	if !output.Scan() {
		fmt.Printf("Failed to scan: %v\n", output.Err())
		os.Exit(1)
	}
	
	if signal, err := strconv.Atoi(output.Text()); err != nil {
		fmt.Printf("Failed to parse output: %v\n", err)
		os.Exit(1)
	} else {
		return signal
	}

	return -1
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
