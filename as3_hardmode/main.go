package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	fmt.Println(string(file))
}
