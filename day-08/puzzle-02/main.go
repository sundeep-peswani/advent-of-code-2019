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

	image := renderImage(string(input), width, height)
	if image == nil {
		fmt.Printf("Unable to convert image: %s\n", string(input))
		os.Exit(1)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if image[i][j] == 1 {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	os.Exit(0)
}

func renderImage(input string, width, height int) [][]int {
	layers := makeLayers(input, width, height)
	if layers == nil {
		return nil
	}

	image := make([][]int, height)
	for i := 0; i < height; i++ {
		image[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			index := i * width + j

			for _, layer := range(layers) {
				if layer[index] != 2 {
					image[i][j] = layer[index]
					break
				}
			}
		}
	}

	return image
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
