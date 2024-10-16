package interface_package

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func PrintAbser () {
	var a Abser 
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a= f
	// a=&v

	fmt.Print(a.Abs(), f.Abs(), v.Abs())
}