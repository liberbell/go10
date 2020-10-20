package main

import (
	"fmt"
	"time"
)

func main() {
	go msg()

	fmt.Println("\nMessage from func main, I'm finished.")
	time.Sleep(time.Millisecond * 7500)
}

func msg() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		if i > 3 {
			fmt.Println(i, "Seconds...yawn")
		} else {
			fmt.Println(i, "seconds.")
		}
	}
	fmt.Println("\nMessage from func msg: I'm finished.")
}
