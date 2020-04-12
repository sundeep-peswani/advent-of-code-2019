package main

import "testing"

func Test_System(t *testing.T) {
	tests := []struct {
		steps    int
		input    string
		expected []string
		energy   int
	}{
		{0, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=-1, y=0, z=2>, vel=<x=0, y=0, z=0>", "pos=<x=2, y=-10, z=-7>, vel=<x=0, y=0, z=0>", "pos=<x=4, y=-8, z=8>, vel=<x=0, y=0, z=0>", "pos=<x=3, y=5, z=-1>, vel=<x=0, y=0, z=0>"}, -1},
		{1, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=2, y=-1, z=1>, vel=<x=3, y=-1, z=-1>", "pos=<x=3, y=-7, z=-4>, vel=<x=1, y=3, z=3>", "pos=<x=1, y=-7, z=5>, vel=<x=-3, y=1, z=-3>", "pos=<x=2, y=2, z=0>, vel=<x=-1, y=-3, z=1>"}, -1},
		{2, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=5, y=-3, z=-1>, vel=<x=3, y=-2, z=-2>", "pos=<x=1, y=-2, z=2>, vel=<x=-2, y=5, z=6>", "pos=<x=1, y=-4, z=-1>, vel=<x=0, y=3, z=-6>", "pos=<x=1, y=-4, z=2>, vel=<x=-1, y=-6, z=2>"}, -1},
		{3, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=5, y=-6, z=-1>, vel=<x=0, y=-3, z=0>", "pos=<x=0, y=0, z=6>, vel=<x=-1, y=2, z=4>", "pos=<x=2, y=1, z=-5>, vel=<x=1, y=5, z=-4>", "pos=<x=1, y=-8, z=2>, vel=<x=0, y=-4, z=0>"}, -1},
		{4, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=2, y=-8, z=0>, vel=<x=-3, y=-2, z=1>", "pos=<x=2, y=1, z=7>, vel=<x=2, y=1, z=1>", "pos=<x=2, y=3, z=-6>, vel=<x=0, y=2, z=-1>", "pos=<x=2, y=-9, z=1>, vel=<x=1, y=-1, z=-1>"}, -1},
		{5, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=-1, y=-9, z=2>, vel=<x=-3, y=-1, z=2>", "pos=<x=4, y=1, z=5>, vel=<x=2, y=0, z=-2>", "pos=<x=2, y=2, z=-4>, vel=<x=0, y=-1, z=2>", "pos=<x=3, y=-7, z=-1>, vel=<x=1, y=2, z=-2>"}, -1},
		{6, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=-1, y=-7, z=3>, vel=<x=0, y=2, z=1>", "pos=<x=3, y=0, z=0>, vel=<x=-1, y=-1, z=-5>", "pos=<x=3, y=-2, z=1>, vel=<x=1, y=-4, z=5>", "pos=<x=3, y=-4, z=-2>, vel=<x=0, y=3, z=-1>"}, -1},
		{7, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=2, y=-2, z=1>, vel=<x=3, y=5, z=-2>", "pos=<x=1, y=-4, z=-4>, vel=<x=-2, y=-4, z=-4>", "pos=<x=3, y=-7, z=5>, vel=<x=0, y=-5, z=4>", "pos=<x=2, y=0, z=0>, vel=<x=-1, y=4, z=2>"}, -1},
		{8, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=5, y=2, z=-2>, vel=<x=3, y=4, z=-3>", "pos=<x=2, y=-7, z=-5>, vel=<x=1, y=-3, z=-1>", "pos=<x=0, y=-9, z=6>, vel=<x=-3, y=-2, z=1>", "pos=<x=1, y=1, z=3>, vel=<x=-1, y=1, z=3>"}, -1},
		{9, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=5, y=3, z=-4>, vel=<x=0, y=1, z=-2>", "pos=<x=2, y=-9, z=-3>, vel=<x=0, y=-2, z=2>", "pos=<x=0, y=-8, z=4>, vel=<x=0, y=1, z=-2>", "pos=<x=1, y=1, z=5>, vel=<x=0, y=0, z=2>"}, -1},
		{10, "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", []string{"pos=<x=2, y=1, z=-3>, vel=<x=-3, y=-2, z=1>", "pos=<x=1, y=-8, z=0>, vel=<x=-1, y=1, z=3>", "pos=<x=3, y=-6, z=1>, vel=<x=3, y=2, z=-3>", "pos=<x=2, y=0, z=4>, vel=<x=1, y=-1, z=-1>"}, 179},

		{0, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=-8, y=-10, z=0>, vel=<x=0, y=0, z=0>", "pos=<x=5, y=5, z=10>, vel=<x=0, y=0, z=0>", "pos=<x=2, y=-7, z=3>, vel=<x=0, y=0, z=0>", "pos=<x=9, y=-8, z=-3>, vel=<x=0, y=0, z=0>"}, -1},
		{10, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=-9, y=-10, z=1>, vel=<x=-2, y=-2, z=-1>", "pos=<x=4, y=10, z=9>, vel=<x=-3, y=7, z=-2>", "pos=<x=8, y=-10, z=-3>, vel=<x=5, y=-1, z=-2>", "pos=<x=5, y=-10, z=3>, vel=<x=0, y=-4, z=5>"}, -1},
		{20, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=-10, y=3, z=-4>, vel=<x=-5, y=2, z=0>", "pos=<x=5, y=-25, z=6>, vel=<x=1, y=1, z=-4>", "pos=<x=13, y=1, z=1>, vel=<x=5, y=-2, z=2>", "pos=<x=0, y=1, z=7>, vel=<x=-1, y=-1, z=2>"}, -1},
		{30, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=15, y=-6, z=-9>, vel=<x=-5, y=4, z=0>", "pos=<x=-4, y=-11, z=3>, vel=<x=-3, y=-10, z=0>", "pos=<x=0, y=-1, z=11>, vel=<x=7, y=4, z=3>", "pos=<x=-3, y=-2, z=5>, vel=<x=1, y=2, z=-3>"}, -1},
		{40, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=14, y=-12, z=-4>, vel=<x=11, y=3, z=0>", "pos=<x=-1, y=18, z=8>, vel=<x=-5, y=2, z=3>", "pos=<x=-5, y=-14, z=8>, vel=<x=1, y=-2, z=0>", "pos=<x=0, y=-12, z=-2>, vel=<x=-7, y=-3, z=-3>"}, -1},
		{50, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=-23, y=4, z=1>, vel=<x=-7, y=-1, z=2>", "pos=<x=20, y=-31, z=13>, vel=<x=5, y=3, z=4>", "pos=<x=-4, y=6, z=1>, vel=<x=-1, y=1, z=-3>", "pos=<x=15, y=1, z=-5>, vel=<x=3, y=-3, z=-3>"}, -1},
		{60, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=36, y=-10, z=6>, vel=<x=5, y=0, z=3>", "pos=<x=-18, y=10, z=9>, vel=<x=-3, y=-7, z=5>", "pos=<x=8, y=-12, z=-3>, vel=<x=-2, y=1, z=-7>", "pos=<x=-18, y=-8, z=-2>, vel=<x=0, y=6, z=-1>"}, -1},
		{70, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=-33, y=-6, z=5>, vel=<x=-5, y=-4, z=7>", "pos=<x=13, y=-9, z=2>, vel=<x=-2, y=11, z=3>", "pos=<x=11, y=-8, z=2>, vel=<x=8, y=-6, z=-7>", "pos=<x=17, y=3, z=1>, vel=<x=-1, y=-1, z=-3>"}, -1},
		{80, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=30, y=-8, z=3>, vel=<x=3, y=3, z=0>", "pos=<x=-2, y=-4, z=0>, vel=<x=4, y=-13, z=2>", "pos=<x=-18, y=-7, z=15>, vel=<x=-8, y=2, z=-2>", "pos=<x=-2, y=-1, z=-8>, vel=<x=1, y=8, z=0>"}, -1},
		{90, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=-25, y=-1, z=4>, vel=<x=1, y=-3, z=4>", "pos=<x=2, y=-9, z=0>, vel=<x=-3, y=13, z=-1>", "pos=<x=32, y=-8, z=14>, vel=<x=5, y=-4, z=6>", "pos=<x=-1, y=-2, z=-8>, vel=<x=-3, y=-6, z=-9>"}, -1},
		{100, "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", []string{"pos=<x=8, y=-12, z=-9>, vel=<x=-7, y=3, z=0>", "pos=<x=13, y=16, z=-3>, vel=<x=3, y=-11, z=-5>", "pos=<x=-29, y=-11, z=-1>, vel=<x=-3, y=7, z=4>", "pos=<x=16, y=-13, z=23>, vel=<x=7, y=1, z=1>"}, 1940},
	}

	for i, test := range tests {
		s := newSystem(test.input)

		if len(s.moons) != len(test.expected) {
			t.Fatalf("Invalid test %d", i+1)
		}

		s.simulate(test.steps)
		for j, m := range s.moons {
			if m.String() != test.expected[j] {
				t.Errorf("Test %d, '%s' + %d steps, moon %d: expected %s, actual %s\n", i+1, test.input, test.steps, j+1, test.expected[j], m.String())
			}
		}

		if test.energy != -1 && s.energy() != test.energy {
			t.Errorf("Test %d, '%s' + %d steps, expected energy of %d, actual %d\n", i+1, test.input, test.steps, test.energy, s.energy())
		}
	}
}

func Test_ThreeDim(t *testing.T) {
	tests := []struct {
		input    string
		expected threeDim
	}{
		{"<x=-1, y=0, z=2>", threeDim{-1, 0, 2}},
		{"<x=2, y=-10, z=-7>", threeDim{2, -10, -7}},
		{"<x=4, y=-8, z=8>", threeDim{4, -8, 8}},
		{"<x=3, y=5, z=-1>	", threeDim{3, 5, -1}},
	}

	for i, test := range tests {
		actual := newThreeDim(test.input)

		if !equal(test.expected, actual) {
			t.Errorf("Test %d: expected %s, actual %s\n", i+1, test.expected, actual)
		}
	}
}

func equal(a, b threeDim) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}
