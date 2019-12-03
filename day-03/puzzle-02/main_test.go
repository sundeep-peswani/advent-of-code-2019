package main

import "testing"

func Test_FindNearestIntersection(t *testing.T) {
	tests := []struct{
		wires []string
		expected int
	}{
		{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 30},
		{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
		{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410},
	}

	for i, test := range tests {
		actual := findNearestIntersection(test.wires)

		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}