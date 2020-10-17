package main

import "fmt"

func main() {
	w := make(map[string]int)
	w["answer"] = 10

	fmt.Println("The value: ", w["answer"])

	w["answer"] = 20
	fmt.Println("The value: ", w["answer"])

	delete(w, "answer")
	fmt.Println("The value: ", w["answer"])
}
