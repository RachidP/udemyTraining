package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {

	return math.Pi * math.Pow(c.radius, 2)
}
func (r square) area() float64 {

	return math.Pow(r.side, 2)
}

func main() {
	c := circle{6}
	s := square{2}
	info(c)
	info(s)
}
func info(s shape) {
	tipo := fmt.Sprintf("%T", s)
	fmt.Printf(" area of %s = %v \n", tipo[len(tipo)-6:], s.area())

}
