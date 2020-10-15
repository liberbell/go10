package main

import "fmt"

type Rect struct {
	Width, Height int
}

var (
	r1 = Rect{7, 8}
	r2 = Rect{Width: 4}
	r3 = Rect{}
	p  = &Rect{7, 8}
)

func main() {
	fmt.Println(r1, r2, r3, p)
}
