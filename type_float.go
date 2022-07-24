package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3
	var a2 float64 = 1234.523
	var b2 float64 = 3456.123
	var c2 float64 = a2 * b2
	var d2 float64 = c2 * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2)

	a3 := a2 * c2
	fmt.Println(a3)
}
