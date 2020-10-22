package main

import "fmt"

func main() {
	slice := []int{10, 15, 20, 25}
	fmt.Println("\nHere is our slice:")
	fmt.Println("Slice ==", slice)
	fmt.Println("Slice[1:4]", slice[1:4])
	fmt.Println("Slice[:3]", slice[:3])
	fmt.Println("Slice[2:]", slice[2:])
	fmt.Println("len Slice ==", len(slice))
	fmt.Println("cap Slice ==", cap(slice))

	for i, v := range slice {
		slice[i] = v - 5
	}

	fmt.Println("\nThe new values in slices: ")
	report("slice", slice)

	fmt.Println("\nNow we'll append 2 values to slice(what happen?)")
	slice = append(slice, 10, 20)
	report("slice", slice)

	fmt.Println("\nNow we we'll append 8 more values to slice(now what happen?)")
	slice = resize(slice)
	report("slice", slice)

	fmt.Println("\nNow we we'll append 8 more values to slice(guess what happen?")
	slice = resize(slice)
	report("slice", slice)

	fmt.Println("\nMake a copy, only the copy has a cap of 8:")
	slicecopy := make([]int, 8)
	copy(slice, slicecopy)
	report("sliceCopy", slicecopy)

	fmt.Println("\nLet's slice our slice for a reslice:")
	reslice := slice[1:5]
	report("reslice", reslice)
}

func resize(slice []int) []int {
	for i := 0; i < 8; i++ {
		slice = append(slice, 1)
	}
	return slice
}

func report() {
	fmt.Println(name, "=", slice)
	fmt.Println("len", name, len(slice))
	fmt.Println("cap", name, "==", cap(slice))
}
