package main

import (
	"calculator/arithmetic"
	"calculator/geometric"
	"fmt"
)

func main() {
	fmt.Println("Addition: ",arithmetic.Add(10, 20))
	fmt.Println("Subtraction: ",arithmetic.Sub(10, 20))
	fmt.Println("Multiplicatin: ",arithmetic.Mul(10, 20))
	fmt.Println("Division: ",arithmetic.Div(10, 20))
	fmt.Println("Area of circle: ",geometric.AreaOfCircle(12.0))
}