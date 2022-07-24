package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "인천 동구 화도진로 110-4 B동 401호"
	house.Price = 3500
	house.Size = 15
	house.Type = "빌라"

	fmt.Println("주소: ", house.Address)
	fmt.Printf("가격: %.0f만원\n", house.Price)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Println("형태: ", house.Type)
}
