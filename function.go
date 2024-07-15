package main

import (
	"fmt"
	"strings"
)

func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	elements := []string{"CDAC", "ACTS", "DHPCAP"}
	fmt.Println(joinstr(elements...))
	fmt.Println(joinstr("CDAC", "ACTS", "DHPCAP"))

	sum(4, 5, 6)
	sum(4, 5, 6, 7, 8, 9)

	func() {
		fmt.Println("Welcome to CDAC ACTS")
	}()

	value := func() {
		fmt.Println("Welcome to CDAC ACTS")
	}

	value()

	var sum = func(n1, n2 int) {
		sum := n1 + n2
		fmt.Println("Sum is", sum)
	}

	sum(5, 3)

	var sum1 = func(n1, n2 int) int {

		sum1 := n1 + n2

		return sum1

	}

	result := sum1(10, 10)

	println(result)

}