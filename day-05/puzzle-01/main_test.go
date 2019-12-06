package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_Intcode(t *testing.T) {
	tests := []struct{
		program string
		expected []int
		input, output string
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", []int{3500,9,10,70,2,3,11,0,99,30,40,50}, "", ""},
		{"1,0,0,0,99", []int{2,0,0,0,99}, "", ""},
		{"2,3,0,3,99", []int{2,3,0,6,99}, "", ""},
		{"2,4,4,5,99,0", []int{2,4,4,5,99,9801}, "", ""},
		{"1,1,1,4,99,5,6,0,99", []int{30,1,1,4,2,5,6,0,99}, "", ""},
		{"3,0,4,0,99", []int{50,0,4,0,99}, "50\n", "50\n"},
		{"1101,100,-1,4,0", []int{1101,100,-1,4,99}, "", ""},
	}

	for i, test := range tests {
		var sut intcode
		var b bytes.Buffer

		sut.in = bufio.NewReader(strings.NewReader(test.input))
		sut.out = bufio.NewWriter(&b)

		sut.load(test.program)
		sut.run()

		if !equal(test.expected, sut.program) {
			t.Errorf("Test %d: expected %v, actual %v\n", i + 1, test.expected, sut.program)
		}

		if test.input != "" && b.String() != test.output {
			t.Errorf("Test %d: expected output of %s, actual %s\n", i + 1, test.output, b.String())
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := range a {
		if j != b[i] {
			return false
		}
	}

	return true
}
