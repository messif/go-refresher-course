package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
	}
	//delete(colors, "red")
	var colors2 map[string]string
	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	fmt.Println(colors, colors2, colors3)
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color: %v, Hex: %v\n", color, hex)
	}
}
