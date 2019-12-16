package main

import (
	"fmt"
	"testing"
)

func Test_FindMaxThrusterSignal(t *testing.T) {
	tests := []struct{
		program string
		expected int
	}{
		{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", 43210},
		{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", 54321},
		{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", 65210},
	}

	for i, test := range tests {
		actual := maxThrusterSignal(test.program, []int{0,1,2,3,4})

		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}

func Test_RunAmps(t *testing.T) {
	tests := []struct{
		program string
		options []int
		expected int
	}{
		{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", []int{4,3,2,1,0}, 43210},
		{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", []int{0,1,2,3,4}, 54321},
		{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", []int{1,0,4,3,2}, 65210},
	}

	for i, test := range tests {
		actual := runAmps(test.program, test.options)

		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}

func Test_GeneratePermutations(t *testing.T) {
	tests := []struct{
		options []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{0,1,2,3,4}, 120},
	}

	for i, test := range tests {
		actual := generatePermutations(test.options)

		if len(actual) != test.expected {
			t.Errorf("Test %d: expected %d items, found %d\n", i + 1, len(actual), test.expected)
		}

		if !isValidPermutations(test.options, actual) {
			t.Errorf("Test %d: invalid permutation set: %v\n", i + 1, actual)
		}
	}
}

func isValidPermutations(options []int, perms [][]int) bool {
	seen := make(map[string]bool)

	for _, perm := range perms {
		key := fmt.Sprintf("%v", perm)

		if len(perm) != len(options) {
			fmt.Printf("Invalid permutation size: %v\n", perm)
			return false
		}

		for _, p := range perm {
			if !has(options, p) {
				fmt.Printf("Found invalid perm: %v\n", perm)
				return false	
			}
		}

		if _, ok := seen[key]; ok {
			fmt.Printf("Found duplicate %v\n", perm)
			return false
		}

		seen[key] = true
	}

	return true
}

func has(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}