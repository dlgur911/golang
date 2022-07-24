package main

import "fmt"

func main() {
	var a int = 1234
	var b int = 3456
	var f float32 = float32(a * b)

	fmt.Print("a :", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	fmt.Printf("a: %d, b: %d, f: %f\n", a, b, f)
}
