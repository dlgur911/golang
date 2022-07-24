package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Level int
	Price int
}

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수 %.2f\n", s.Age, s.No, s.Score)
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
	fmt.Println()

	user := User{"이혁", "hyklee", 52, 5}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 1},
		0,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이 %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이 %d VIP레벨: %d VIP가격 %d만원 레벨 %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.Price,
		vip.User.Level,
	)
	fmt.Println()

	var student = Student{15, 23, 88.2}
	/*
		var student Student
		student.Age = 15
		student.No = 23
		student.Score = 88.3
	*/
	student2 := student //구조체 값 복사
	PrintStudent(student2)
}
