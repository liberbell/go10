package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	fmt.Println(x)
	return func() int {
		x, y = y, x+y
	}
}
