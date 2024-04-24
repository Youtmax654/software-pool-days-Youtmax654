package main

// Artists Book application
import (
	"SoftwareGoDay1/artistsbook"
)

func main() {
	artistsbook.Router();
}

// ---------------------------------------------------------------

// Geometric Shape application
// import (
// 	"SoftwareGoDay1/geometricshape"
// 	"fmt"
// )

// func main() {
// 	var circle = geometricshape.Circle{Radius: 12.0}
// 	var rectangle = geometricshape.Rectangle{X: 2.0, Y: 3.0}
// 	var triangle = geometricshape.Triangle{X: 5.0, Y: 2.0, Z: 1.0}

// 	var circlePerimeter = geometricshape.CalcPerimeter(circle);
// 	var rectanglePerimeter = geometricshape.CalcPerimeter(rectangle);
// 	var trianglePerimeter = geometricshape.CalcPerimeter(triangle);

// 	fmt.Println(&circle, "has a perimeter of", circlePerimeter)
// 	fmt.Println(&rectangle, "has a perimeter of", rectanglePerimeter)
// 	fmt.Println(&triangle, "has a perimeter of", trianglePerimeter)
// }