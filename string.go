package main

import "fmt"

func main() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	}

	fmt.Println()

	for _, val := range arr {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", val, val, val)
	}
}
