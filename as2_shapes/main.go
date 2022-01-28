package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func main() {
	sq := square{side: 12.2}
	tr := triangle{base: 8.3, height: 12.9}
	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}
