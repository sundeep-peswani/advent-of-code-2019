package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const COM = "COM"

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

	u := newUniversalOrbitMap()
	u.load(strings.Split(string(input), "\n"))

	fmt.Println(u.countRelationships())
	fmt.Println(u.countOrbitalTransfers("YOU", "SAN"))
}

type spaceObject struct {
	label string
	orbits *spaceObject
	isOrbitedBy []*spaceObject
}

// exclusive of s itself
func (s spaceObject) getPath() []string {
	var path []string

	curr := s.orbits
	for curr != nil && curr.label != COM {
		path = append(append([]string{}, curr.label), path...)
		curr = curr.orbits
	}

	return path
}

func (s spaceObject) countOrbits() int {
	total := len(s.getPath())
	if s.label != COM {
		return total + 1
	}
	return total
}

type universalOrbitMap struct {
	com *spaceObject
	spaceObjects map[string]*spaceObject
}

func newUniversalOrbitMap() *universalOrbitMap {
	var u universalOrbitMap
	u.spaceObjects = make(map[string]*spaceObject)

	return &u
}

func (u *universalOrbitMap) load(orbitalMap []string) {
	for _, orbitalRelationship := range orbitalMap {
		relationship := strings.Split(orbitalRelationship, ")")
		u.addRelationship(relationship[0], relationship[1])
	}
}

func (u *universalOrbitMap) addRelationship(orbitee, orbiter string) {
	a, b := u.addObject(orbitee), u.addObject(orbiter)

	a.isOrbitedBy = append(a.isOrbitedBy, b)
	b.orbits = a

	if orbitee == COM {
		u.com = a
	}
}

func (u *universalOrbitMap) addObject(label string) *spaceObject {
	if existing := u.get(label); existing != nil {
		return existing
	}

	var newSpaceObject spaceObject
	newSpaceObject.label = label
	u.spaceObjects[label] = &newSpaceObject
	return &newSpaceObject
}

func (u universalOrbitMap) countRelationships() int {
	if u.com == nil {
		return 0
	}
	
	total := 0

	for _, obj := range u.spaceObjects {
		total += obj.countOrbits()
	}

	return total
}

func (u universalOrbitMap) countOrbitalTransfers(a, b string) int {
	src, dst := u.get(a), u.get(b)
	if src == nil {
		return -1
	}
	if dst == nil {
		return -1
	}

	srcPath, dstPath := src.getPath(), dst.getPath()

	total := len(srcPath) + len(dstPath)
	for i := 0; i < len(srcPath) && i < len(dstPath) && srcPath[i] == dstPath[i]; i++ {
		total -= 2
	}
	return total
}

func (u universalOrbitMap) get(label string) *spaceObject {
	if existing, ok := u.spaceObjects[label]; ok {
		return existing
	}

	return nil
}