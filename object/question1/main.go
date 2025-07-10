package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rec *Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (rec *Rectangle) Perimeter() float64 {
	return 2 * (rec.Width + rec.Height)
}

func (cir *Circle) Area() float64 {
	return 3.14 * cir.Radius * cir.Radius
}

func (cir *Circle) Perimeter() float64 {
	return 2 * 3.14 * cir.Radius
}

func main() {
	shapes := []Shape{
		&Rectangle{Width: 5, Height: 10},
		&Circle{Radius: 7},
	}

	for _, shape := range shapes {
		switch s := shape.(type) {
		case *Rectangle:
			fmt.Printf("Rectangle's Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
		case *Circle:
			fmt.Printf("Circle's Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
		default:
			fmt.Println("Unknown shape")
		}
	}
}
