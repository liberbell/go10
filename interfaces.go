package main

import "math"

type Calcer interface{ Calc() float64 }
type Square struct{ X, Y float64 }
type myFloat float64

func main() {
	var c Calcer
	f := myFloat(-math.Sqrt2)
	s := Square{8, 6}
}
