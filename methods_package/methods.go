package methods_package

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func (v *Vertex) scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

func scale (v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func CreateVertex() {
	v := Vertex{3, 4}
	// v.scale(10)
	scale(&v, 10)

	fmt.Println(v.abs())
	fmt.Println(v)

}