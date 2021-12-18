package Fundamental

import "math"

type Shape interface {
	GetArea() float32
	GetPerimeter() float32
}

type Rectangles struct {
	Width  float32
	Length float32
}

//menghitung luas
func (r Rectangles) GetArea() float32 {
	return r.Width * r.Length
}

//menghitung keliling
func (r Rectangles) GetPerimeter() float32 {
	return r.Width * r.Length
}

type Circle struct {
	radius float32
}

// fungsi tersendiri
func (c *Circle) GetRadius() float32 {
	return c.radius
}

//fungsi tersendiri
func (c *Circle) SetRadius(radius float32) {
	c.radius = radius
}

func (c Circle) GetArea() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) GetPerimeter() float32 {
	return 2 * math.Pi * c.radius
}
