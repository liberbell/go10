package main

import "fmt"

func main() {
	slice1 := []int{7, 8, 9}
	slice2 := append(slice1, 11, 12)

	copy(slice2, slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)
}
