package main

func main() {
	c := make(chan int, 5)
	c <- 1
}
