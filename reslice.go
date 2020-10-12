package main

import "fmt"

func main() {
	slice := []int{7, 9, 10, 12}
	fmt.Println("Slice == ", slice)

	fmt.Println("Slice[1:4] == ", slice[1:4])
}
