package main

import "fmt"

func main() {
	slice1 := []int{7, 8, 9}
	slice2 := make([]int, 2)

	copy(slice2, slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)
}
