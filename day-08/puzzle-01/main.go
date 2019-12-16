package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		os.Exit(1)
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		os.Exit(1)
	}

	filename := os.Args[3]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	layers := makeLayers(string(input), width, height)
	if layers == nil {
		fmt.Printf("Unable to convert image: %s\n", string(input))
		os.Exit(1)
	}

	indexOfFewestZeroes := -1
	zeroes, ones, twos := width * height, 0, 0
	for i, layer := range layers {
		count := make([]int, 9)
		for j := 0; j < len(layer); j++ {
			count[layer[j]]++
		}

		if count[0] < zeroes {
			indexOfFewestZeroes = i
			zeroes, ones, twos = count[0], count[1], count[2]
		}
	}

	fmt.Printf("Layer %d: 1s * 2s = %d\n", indexOfFewestZeroes, ones * twos)

	os.Exit(0)
}

func makeLayers(image string, width, height int) [][]int {
	layerSize := width * height
	
	if len(image) % (width * height) != 0 {
		fmt.Printf("Invalid image size: %d (%d x %d)\n", len(image), width, height)
		return nil
	}

	numLayers := len(image) / layerSize
	layers := make([][]int, numLayers)

	for i := 0; i < numLayers; i++ {
		layers[i] = make([]int, layerSize)
	}
	
	for i := 0; i < len(image); i++ {
		layerIndex := i / layerSize
		pixelIndex := i % layerSize

		if p, err := strconv.Atoi(string(image[i])); err != nil {
			fmt.Printf("Invalid image at %d: %v: %v\n", i, image[i], err)
			os.Exit(1)
		} else {

			layers[layerIndex][pixelIndex] = p
		}
	}

	return layers
}
