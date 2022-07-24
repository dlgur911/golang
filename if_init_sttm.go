package main

import "fmt"

func getMyAge(a int) (int, bool) {
	return a, true
}

func main() {
	var a int

	n, err := fmt.Scan(&a)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a)
	}

	if age, ok := getMyAge(a); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("You are is", age)
}
