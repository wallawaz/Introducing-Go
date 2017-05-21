package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x, y, x2, y2 float64
}

//do not define fields; define `method set`
//-->a list of methods that a type must have
type Shape interface {
	area() float64
}

//Circle `method`
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) distance(side string) float64 {
	w := 0.0
	l := 0.0
	switch side {
	case "width":
		w += r.x2 - r.x
	case "length":
		l += r.y2 - r.y
	default:
		w += 0.0
		l += 0.0
	}
	return math.Sqrt(w*w + l*l)
}

func (r *Rectangle) area() float64 {
	l := r.distance("length")
	w := r.distance("width")
	return w * l
}

func totalArea(circles []Circle, rectangles []Rectangle) float64 {
	var total float64

	for _, c := range circles {
		total += c.area()
	}
	for _, r := range rectangles {
		total += r.area()
	}
	return total
}

//using the `Shape` interface
func totalAreaShape(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{x: 1, y: 2, r: 1.5}
	fmt.Println(c)

	fmt.Printf("The area of the %T is: %v\n", c, c.area())

	r := Rectangle{x: 1.5, y: 2.5, x2: 3, y2: 5}
	fmt.Printf("The area of the %T is: %v\n", r, r.area())

	fmt.Println("Total Area combined")
	fmt.Println(totalAreaShape(&c, &r))
}
