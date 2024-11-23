package oop

import (
	"fmt"
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func TestPolymorphism(t *testing.T) {
	circle := Circle{
		Radius: 5.0,
	}
	rectangle := Rectangle{
		Width:  4.0,
		Height: 6.0,
	}
	PrintArea(circle)
	PrintArea(rectangle)
}
