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
		{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", 139629729},
		{"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", 18216},
	}

	for i, test := range tests {
		actual := maxThrusterSignal(test.program, []int{5,6,7,8,9})

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
		{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", []int{9,8,7,6,5}, 139629729},
		{"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", []int{9,7,8,5,6}, 18216},
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
		{[]int{5,6,7,8,9}, 120},
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