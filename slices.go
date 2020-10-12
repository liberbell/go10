package main

import "fmt"

func main() {
	slices := []int{7, 9, 11, 13}
	fmt.Println("slices ==", slices)

	for i := 0; i < len(slices); i++ {
		fmt.Printf("slices[%d] == %d\n", i, slices[i])
	}
}
