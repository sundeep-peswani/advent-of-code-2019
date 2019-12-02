package main

import "testing"

func Test_CalculateFuel(t *testing.T) {
	tests := []struct{
		input, expected int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for i, test := range tests {
		actual := calculateFuel(test.input)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}