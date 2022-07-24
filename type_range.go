package main

import "fmt"

var g int64 = 10

func main() {
	a := 3456
	c := a
	m := 20

	fmt.Println(a, c)

	{
		s := 50
		fmt.Println(m, s, g)
	}

	//m = s + 20

}
