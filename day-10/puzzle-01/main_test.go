package main

import (
	"math"
	"sort"
	"testing"
)

func Test_GetOptimalLocation(t *testing.T) {
	tests := []struct {
		input    string
		optimal  coord
		expected int
	}{
		{".#..#\n.....\n#####\n....#\n...##", coord{3, 4}, 8},
		{"......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####", coord{5, 8}, 33},
		{"#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###.", coord{1, 2}, 35},
		{".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#..", coord{6, 3}, 41},
		{".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##", coord{11, 13}, 210},
	}

	for i, test := range tests {
		sm := newSpaceMap(test.input)
		optimal, count := sm.getOptimalLocation()
		if optimal != test.optimal {
			t.Errorf("Test %d: expected optimal location to be at %s, got %s\n", i+1, test.optimal, optimal)
		}
		if count != test.expected {
			t.Errorf("Test %d: expected view count to be %d, got %d\n", i+1, test.expected, count)
		}
	}
}

func Test_ByDistance(t *testing.T) {
	coords := []polar{polar{1.00000, 1.57080}, polar{1.41421, 0.78540}, polar{2.00000, 0.00000}, polar{3.00000, 0.00000}, polar{3.16228, 0.32175}, polar{2.23607, 0.46365}, polar{3.60555, 0.98279}, polar{3.16228, 1.24905}, polar{4.00000, 1.57080}, polar{4.12311, 0.24498}, polar{4.24264, 0.78540}, polar{4.12311, 1.32582}, polar{4.47214, 1.10715}, polar{3.00000, 1.57080}}
	sort.Sort(byDistance(coords))

	for i, j := 0, 1; j < len(coords); i++ {
		if coords[i].phi > coords[j].phi {
			t.Errorf("phi: %s > %s\n", coords[i], coords[j])
		} else if coords[i].phi == coords[j].phi && coords[i].r > coords[j].r {
			t.Errorf("r: %s > %s\n", coords[i], coords[j])
		}
		j++
	}
}

func Test_IsBehind(t *testing.T) {
	tests := []struct {
		a, b     polar
		expected bool
	}{
		{polar{3.00000, 0.00000}, polar{2.00000, 0.00000}, true},
		{polar{2.00000, 0.00000}, polar{3.00000, 0.00000}, false},
		{polar{4.12311, 0.24498}, polar{3.00000, 0.00000}, false},
		{polar{3.16228, 0.32175}, polar{4.12311, 0.24498}, false},
		{polar{2.23607, 0.32175}, polar{3.16228, 0.32175}, false},
		{polar{1.41421, 0.78540}, polar{2.23607, 0.46365}, false},
		{polar{4.24264, 0.78540}, polar{1.41421, 0.78540}, true},
		{polar{3.60555, 0.98279}, polar{4.24264, 0.78540}, false},
		{polar{1.41421, 0.78540}, polar{4.24264, 0.78540}, false},
	}

	for i, test := range tests {
		actual := test.a.isBehind(test.b)
		if actual != test.expected {
			t.Errorf("Test %d: %s is behind %s = %t, got %t\n", i+1, test.a, test.b, test.expected, actual)
		}
	}
}

func Test_ToPolar(t *testing.T) {
	tests := []struct {
		pole     coord
		input    coord
		expected polar
	}{
		{coord{0, 0}, coord{2, 0}, polar{2, 0}},
		{coord{2, 0}, coord{0, 0}, polar{2, 3.141592}},
	}

	for i, test := range tests {
		actual := test.input.toPolar(test.pole)
		if !test.expected.Equals(actual) {
			t.Errorf("Test %d: expected %#v, got %#v\n", i+1, test.expected, actual)
		}
	}
}

func (p polar) Equals(other polar) bool {
	return isAcceptable(p.r-other.r) && isAcceptable(p.phi-other.phi)
}

func isAcceptable(delta float64) bool { return math.Abs(delta) < 0.005 }
