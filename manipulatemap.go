package main

import "fmt"

func main() {
	w := make(map[string]int)
	w["answer"] = 10

	fmt.Println("The value: ", w["answer"])

	w["anser"] = 20
	fmt.Println("The value: ", w["answer"])
}
