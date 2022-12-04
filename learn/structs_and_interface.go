package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

// In between the keyword func and the name of the function we've added a “receiver”.
// The receiver is like a parameter
// (c * Circle) === receiver
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	var l float64 = distance(x1, y1, x1, y2)
	var w float64 = distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

// func circleArea2(c Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// Embedded Types
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

// Rectangle and Circle has samething = area method
// TODO:: read again
// go has a way of making these accidental similarities explicit throught
// a type known as  an interface. Here is a example of a Shape interface
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}

	return area
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 0, 0
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	// var c Cicle
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	// Fields
	// fmt.Println(c.x, c.y, c.r)
	// c.x = 10
	// c.y = 5
	// Method
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	// Embedded types
	// a := new(Android)
	// a.Person.Talk()
	//
	// a := new(Android)
	// a.Talk()

	fmt.Println(totalArea(&c, &r))

}
