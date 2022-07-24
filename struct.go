package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
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

	user := User{"이혁", "hyklee", 52}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이 %d VIP레벨: %d VIP가격 %d만원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
