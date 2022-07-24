package main

import "fmt"

func getMyAge(a int) (int, bool) {
	return a, true
}

func main() {
	var arr []float32 = []float32{0: 10.0, 3: 20, 2: 30, 1: 40, 4: 50, 7: 60}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, val := range arr {
		fmt.Println(val)
	}
}
