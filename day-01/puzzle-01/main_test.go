package main

import "testing"

func Test_CalulateFuel(t *testing.T) {
	tests := []struct {
		input, expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for i, test := range tests {
		actual := calculateFuel(test.input)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}