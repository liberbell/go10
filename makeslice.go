package main

import "fmt"

func main() {
	a := make([]int, 4)
}

func printS1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
