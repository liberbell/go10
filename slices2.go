package main

import "fmt"

func main() {
	slice := []int{10, 15, 20, 25}
	fmt.Println("\nHere is our slice:")
	fmt.Println("Slice ==", slice)
	fmt.Println("Slice[1:4]", slice[1:4])
	fmt.Println("Slice[:3]", slice[:3])
	fmt.Println("Slice[2:]", slice[2:])
}