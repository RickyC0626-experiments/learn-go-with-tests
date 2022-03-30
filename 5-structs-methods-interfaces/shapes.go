package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Returns area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Returns area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Returns area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
