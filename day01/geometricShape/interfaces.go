package geometricshape

import "fmt"

// GeometricShape interface
type GeometricShape interface {
	CalcPerimeter() float64
}

// CalcPerimeter function
func CalcPerimeter(shape GeometricShape) float64 {
	return shape.CalcPerimeter()
}

// Circle structure
type Circle struct {
	Radius float64
}

// CalcPerimeter method for Circle
func (circle Circle) CalcPerimeter() float64 {
	return 2 * 3.14 * circle.Radius
}

// String method for Circle
func (circle Circle) String() string {
	return fmt.Sprintf("Circle: { radius = %f }", circle.Radius)
}

// Rectangle structure
type Rectangle struct {
	X, Y float64
}

// CalcPerimeter method for Rectangle
func (rectangle Rectangle) CalcPerimeter() float64 {
	return 2 * (rectangle.X + rectangle.Y)
}

// String method for Rectangle
func (rectangle Rectangle) String() string {
	return fmt.Sprintf("Rectangle: { x = %f, y = %f }", rectangle.X, rectangle.Y)
}

// Triangle structure
type Triangle struct {
	X, Y, Z float64
}

// CalcPerimeter method for Triangle
func (triangle Triangle) CalcPerimeter() float64 {
	return triangle.X + triangle.Y + triangle.Z
}

// String method for Triangle
func (triangle Triangle) String() string {
	return fmt.Sprintf("Triangle: { x = %f, y = %f, z = %f }", triangle.X, triangle.Y, triangle.Z)
}