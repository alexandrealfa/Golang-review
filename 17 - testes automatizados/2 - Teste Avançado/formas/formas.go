package formas

import (
	"math"
)

//Form interface
type Form interface {
	area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Raio float64
}

func (r Rectangle) area() float64 {
	return r.Height * r.Width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
