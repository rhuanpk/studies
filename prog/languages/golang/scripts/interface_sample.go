//go:build interfaces

package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.height
}
func (rect Rect) perim() float64 {
	return (2 * rect.width) + (2 * rect.height)
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (circle Circle) perim() float64 {
	return 2 * math.Pi * circle.radius
}

func measure(geometry Geometry) {
	fmt.Println(geometry)
	fmt.Println(geometry.area())
	fmt.Println(geometry.perim())
}

func main() {
	rect := Rect{width: 3, height: 4}
	circle := Circle{radius: 5}

	// Printando dessa forma você terá o mesmo resultado, então a vantagem é que você conseguiu implementar o princípio do polimorfismo e economizou código.

	// fmt.Println(rect)
	// fmt.Println(rect.area())
	// fmt.Println(rect.perim())

	// fmt.Println(circle)
	// fmt.Println(circle.area())
	// fmt.Println(circle.perim())

	measure(rect)
	measure(circle)
}
