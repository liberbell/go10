package main

type Root float64

func (r Root) Abs() float64 {
	if < 0 {
		return float64(-r)
	}
	return float64(r)
}