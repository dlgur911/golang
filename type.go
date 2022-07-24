package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := a * c

	var e int64 = 7
	f := int64(a) * e

	g := b * float64(a)

	fmt.Println(a, d, f, g)

}
