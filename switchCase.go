package main

import "fmt"

func main() {
	mychannels := make(chan int)
	select {
	case <-mychannels:

	default:
		fmt.Println("Not Found")
	}
}