package main

import (
	"fmt"
)

func main() {
	//There are 3 differnet ways to declare map
	//1
	colors := map[string]string{
		"red":   "f443#",
		"green": "cfrtgv",
		"yelow": "456tgy",
	}
	//2
	//var colors1 map[string]string

	//3
	//colors2 := make(map[string]string)

	colors["white"] = "sddddd"

	//How to delete value from map
	delete(colors, "green")
	fmt.Println(colors)
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, key := range c {
		fmt.Println(color, key)
	}
}
