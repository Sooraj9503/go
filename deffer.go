package main

import "fmt"

func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello")
}

func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0

}

func main() {
	mul(10, 20)

	defer func() {
		fmt.Println("Defer")
	}()

	defer mul(23, 56)

	show()

	fmt.Println("Start")

	defer func() {
		fmt.Println("Defer 2")
	}()

	defer fmt.Println("End")
	defer add(34, 56)
	defer add(20, 20)

}