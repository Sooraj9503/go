package main

import "fmt"

func main() {
	var student2 = "Acts" // type is inferred
	var x = 20 // type is inferred
	var Student1 = "CDAC"

	fmt.Println(Student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a, b, c, d int = 1, 2, 3, 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var (
		a1 int = 7
		b1 int = 1
		c1 string = "Hello"
	)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
}