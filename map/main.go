package main

import (
	"fmt"
)

func main() {
	// var colors map[int]string => empty map[]
	// colors := make(map[int]string) => empty map[]
	// colors[10] = "#ffffff" // add item on map
	// delete(colors, 10) // delete item on map

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
