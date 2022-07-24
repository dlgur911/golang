package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	data.value = 1
	data.data[100] = 123

	ChangeData(&data)
	fmt.Printf("value: %d, data[100]: %d", data.value, data.data[100])
	fmt.Println()

	var p *Data = &data
	p.value = 2
	p.data[199] = 200
	//ChangeData(p)
	fmt.Printf("value: %d, data[100]: %d", data.value, data.data[199])
}
